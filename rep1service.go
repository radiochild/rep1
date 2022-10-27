package main

import (
  "context"
  "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapio"
	"github.com/radiochild/repmeta"
	"github.com/radiochild/utils"
  _ "github.com/mattn/go-sqlite3"

  "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
  "github.com/aws/smithy-go"
)


// --------------------------------------------------------------------------------
// Database record retrieval
// Invokes a DetailHandler for each record in SQL ResultSet
// --------------------------------------------------------------------------------
type DetailHandler func(ctx interface{}, dR *repmeta.DataRow)

// pDB was *pgx.Conn
func retrieveRecs(pDB *sql.DB,  qry string, spec *repmeta.ReportSpec, ctx interface{}, dH DetailHandler) error {

	rows, err := pDB.Query(qry)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return err
	}

	for rows.Next() {
		pDR, erx := repmeta.NewDataRow(spec)
		if erx != nil {
			return erx
		}
		ptrs := pDR.GetPointers()

		err = rows.Scan(ptrs...)
		if err != nil {
      fmt.Printf("rows.Scan(ptrs) failed\n")
			return err
		}
		// DetailHandler takes care of the record
		dH(ctx, pDR)
	}

	return rows.Err()
}

// --------------------------------------------------------------------------------
// Rep1Service - Context for dispatching work
// --------------------------------------------------------------------------------
type Rep1Service struct {
	logger    *zap.SugaredLogger
	logWriter io.Writer
  pDB       *sql.DB
	args      *CmdArgs
  bucketName string
  awsConfig aws.Config
  s3Client  *s3.Client
}

func IsSqliteDB(dbName string) bool {
  return strings.HasPrefix(dbName, "sqlite.")
}

func SqliteDBName(dbname string) string {
  parts := strings.SplitN(dbname, "sqlite.", 2)
  lx := len(parts)
  return parts[lx-1]
}

func NeedsS3(pArgs *CmdArgs) bool {
  return len(pArgs.BucketName) > 0
}

func NewRep1Service(pArgs *CmdArgs, dbp *DBParams) *Rep1Service {
	logger := utils.NewLogger(pArgs.LogLevel)
	logWriter := &zapio.Writer{Log: logger.Desugar()}

	logger.Info(appTitle())
	logger.Info(fullVersion())
  var err error
  var db *sql.DB
  if IsSqliteDB(dbp.Database) {
    dbName := SqliteDBName(dbp.Database)
    db, err = sql.Open("sqlite3", dbName)
    if err != nil {
      logger.Fatal(err.Error())
    }
  } else {
    pConnCfg, err := utils.NewDBConfig(dbp.User, dbp.Password, dbp.Host, dbp.Database)
    if err != nil {
      logger.Fatal(err.Error())
    }
    // Connect to database
    logger.Fatalf("Postgres not supported: %v", pConnCfg)
  }

  var awsConfig aws.Config
  var s3Client *s3.Client

  bucketName := pArgs.BucketName
  useS3 := len(bucketName) > 0
  if useS3 {
    awsConfig, err = config.LoadDefaultConfig(context.TODO())
    if err != nil {
      logger.Fatal(err.Error())
    }
	  s3Client = s3.NewFromConfig(awsConfig)

    // Make sure bucket exists (requires s3:GetBucketTagging rights)
    ctx := context.TODO()
    taggingInput := s3.GetBucketTaggingInput {
      Bucket: aws.String(bucketName),
    }
    _, err := s3Client.GetBucketTagging(ctx, &taggingInput)
    if err != nil {
      var ae smithy.APIError
      if errors.As(err, &ae) {
        code := ae.ErrorCode()
        if code == "NoSuchBucket" {
          logger.Infof("Bucket %q: %s", bucketName, ae.ErrorMessage())
          bucketName = ""
        }
      }
    }
  }

	svc := Rep1Service{
		logger:    logger,
		logWriter: logWriter,
		pDB:       db,
		args:      pArgs,
    awsConfig: awsConfig,
    s3Client:  s3Client,
    bucketName: bucketName,
	}
	return &svc
}

func (svc *Rep1Service) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Not found: %s", r.RequestURI)
	svc.logger.Info(msg)
	http.Error(w, msg, http.StatusNotFound)
}

func (svc *Rep1Service) RootHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := os.ReadFile("rep1.go")
	if err != nil {
		svc.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

type Rep1RequestData struct {
	Limit       int
	OutputType  string
	HideDetails bool
	Spec        repmeta.ReportSpec
}

func (svc *Rep1Service) Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data Rep1RequestData
	err := decoder.Decode(&data)
	if errors.Is(err, io.EOF) {
		msg := "Request body must not be empty"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	if err != nil {
		svc.logger.Errorf("JSON Decode Error: %s", err.Error())
		msg := "Decode Error"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	requestedType := data.OutputType
	outputType := repmeta.OTText
	contentType := "text/html"
	if strings.EqualFold(requestedType, "TEXT") {
		outputType = repmeta.OTText
	}
	if strings.EqualFold(requestedType, "JSON") {
		outputType = repmeta.OTJSON
		contentType = "application/x-json"
	}
	if strings.EqualFold(requestedType, "PACK") {
		outputType = repmeta.OTMessagePack
		contentType = "application/x-pack"
	}

	svc.args.Limit = data.Limit
	svc.args.HideDetails = data.HideDetails
	svc.args.OutputType = outputType
	svc.args.Spec = &data.Spec
	svc.args.OutputFile = ""

	w.Header().Set("Content-Type", contentType)
	// w.WriteHeader(http.StatusOK)

	rW := repmeta.NewReportWriter(svc.logger, w, svc.args.OutputType, svc.args.HideDetails, svc.args.Spec, svc.s3Client, svc.bucketName)

	// Fetch and Display records according to spec
	qry := repmeta.FormatQuery(svc.args.Spec, svc.args.Limit, svc.logger)
	err = retrieveRecs(svc.pDB, qry, svc.args.Spec, rW, repmeta.DetailWriter)
	if err != nil {
		svc.logger.Fatalf("Received an error:\n%v\n", err)
	}
	repmeta.DetailWriter(rW, nil)

	rW.ProcessGrandTotals()
}

func (svc *Rep1Service) CmdHandler() {
	svc.logger.Info("Cmdline request")

	fileName := strings.TrimSpace(svc.args.OutputFile)
	outfile := os.Stdout
	if len(fileName) > 0 {
		f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err == nil {
			outfile = f
			defer outfile.Close()
		}
	}
	rW := repmeta.NewReportWriter(svc.logger, outfile, svc.args.OutputType, svc.args.HideDetails, svc.args.Spec, svc.s3Client, svc.bucketName)

	// Fetch and Display records according to spec
	qry := repmeta.FormatQuery(svc.args.Spec, svc.args.Limit, svc.logger)
	err := retrieveRecs(svc.pDB, qry, svc.args.Spec, rW, repmeta.DetailWriter)
	if err != nil {
		svc.logger.Fatalf("Received an error:\n%v\n", err)
	}
	repmeta.DetailWriter(rW, nil)

	rW.ProcessGrandTotals()
}

func (svc *Rep1Service) Close() error {
	svc.logger.Info("Request to close Rep1Service")
	if svc.pDB != nil {
		svc.pDB.Close()
		svc.pDB = nil
		svc.logger.Info("Closing Rep1Service")
	}
	// Sync the logger
	dLogger := svc.logger.Desugar()
	return dLogger.Sync()
}

func (svc *Rep1Service) LogSpecAndQuery() {
	// Read and Display Report Spec
	repmeta.ShowReportSpec(svc.args.Spec, svc.logger)
	svc.logger.Info("")

	svc.logger.Info("Query:")
	qry := repmeta.FormatQuery(svc.args.Spec, svc.args.Limit, svc.logger)
	svc.logger.Info(qry)
	svc.logger.Info("")
}
