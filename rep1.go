package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "github.com/jackc/pgx/v4"
)

func main() {

	// ----------------------------------------------------------------------
	// Rep1Service Setup
	// ----------------------------------------------------------------------
	args := NewCmdArguments()
	if args.Version {
		fmt.Printf("%s %s\n", AppName, fullVersion())
		return
	}

	dbp := NewDBParamsFromEnv()
	svc := NewRep1Service(args, dbp)

	if args.CmdExec {
		svc.CmdHandler()
		svc.Close()
		return
	}

	// ----------------------------------------------------------------------
	// Setup Router with endpoint handlers
	// Router is used by both AWS Lambda adapter and local HTTP server
	// ----------------------------------------------------------------------
	rtr := mux.NewRouter()
	rtr.NotFoundHandler = http.HandlerFunc(svc.NotFoundHandler)

	subRtr := rtr.PathPrefix("/default").Subrouter()
	subRtr.HandleFunc("/rep1", svc.Handler)
	subRtr.HandleFunc("/", svc.RootHandler)

	// Set up logging and panic recovery middleware.
	rtr.Use(func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(svc.logWriter, h)
	})

	// ----------------------------------------------------------------------
	// Dispatch to Lambda or Setup HTTP service
	// ----------------------------------------------------------------------
	if runtimeAPI, _ := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); runtimeAPI != "" {
		svc.logger.Info("Lambda adapted to http handler")
		adapter := gorillamux.NewV2(rtr)
		lambda.Start(adapter.ProxyWithContext)
	} else {
		svrPort := args.ServerPort
		srv := &http.Server{
			ReadTimeout:       2 * time.Second,
			WriteTimeout:      900 * time.Second,
			IdleTimeout:       2 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
			Addr:              fmt.Sprintf(":%d", svrPort),
			Handler:           rtr,
		}
		svc.logger.Infof("Starting http service on port %d", svrPort)
		_ = srv.ListenAndServe()
	}
}
