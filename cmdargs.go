package main

import (
	"flag"
	"strings"

	"github.com/radiochild/repmeta"
	"github.com/radiochild/utils"
)

const logLevelEnvName = "REP1_LOG_LEVEL"
const defaultLogLevel = "info"

const serverPortEnvName = "REP1_PORT"
const defaultServerPort = 9119

type CmdArgs struct {
	Spec        *repmeta.ReportSpec
	Limit       int
	HideDetails bool
	OutputType  repmeta.OutputType
	OutputFile  string
	LogLevel    string
	CmdExec     bool
	Version     bool
	ServerPort  int
  BucketName  string
}

func NewCmdArguments() *CmdArgs {

	pLimit := flag.Int("limit", -1, "max number of data rows to process")
	pSpecsFile := flag.String("spec", "", "report spec")
	pHideDetails := flag.Bool("hideDetails", false, "hide details")
	pOutputFile := flag.String("outputFile", "", "output file")
	pOutputType := flag.String("outputType", "TEXT", "output type [TEXT | JSON | PACK]")
	pLogLevel := flag.String("logLevel", "", "log level [debug | info | warn | error | panic | fatal]")
	pCmd := flag.Bool("cmd", false, "command execution - no server started")
	pVersion := flag.Bool("version", false, "command version")
	pAltVersion := flag.Bool("v", false, "shorthand command version")
	pServerPort := flag.Int("port", 0, "http server port")
	pBucketName := flag.String("bucket", "", "bucket name [lowercase a-z 0-9 or -]")
	flag.Parse()

	pSpec, _ := repmeta.ReadReportSpec(*pSpecsFile)
	hideDetails := *pHideDetails
	cmdLogLevel := strings.TrimSpace(strings.ToLower(*pLogLevel))

	requestedType := *pOutputType
	outputType := repmeta.OTText
	if strings.EqualFold(requestedType, "TEXT") {
		outputType = repmeta.OTText
	}
	if strings.EqualFold(requestedType, "JSON") {
		outputType = repmeta.OTJSON
	}
	if strings.EqualFold(requestedType, "PACK") {
		outputType = repmeta.OTMessagePack
	}

	wantVersion := *pVersion || *pAltVersion
	if *pVersion {
		*pCmd = wantVersion
	}

  bucketName := strings.ToLower(strings.TrimSpace(*pBucketName))

	args := CmdArgs{
		Spec:        pSpec,
		Limit:       *pLimit,
		HideDetails: hideDetails,
		OutputType:  outputType,
		OutputFile:  *pOutputFile,
		LogLevel:    cmdLogLevel,
		CmdExec:     *pCmd,
		ServerPort:  *pServerPort,
		Version:     wantVersion,
    BucketName:  bucketName,
	}

	// If CmdArgs were not specified for LogLevel:
	// Use the ENV setting, which falls back to a hard-coded default
	if args.LogLevel == "" {
		args.LogLevel = utils.StringFromEnv(logLevelEnvName, defaultLogLevel)
	}

	// If CmdArgs were not specified for ServerPort:
	// Use the ENV setting, which falls back to a hard-coded default
	if args.ServerPort == 0 {
		args.ServerPort = utils.IntFromEnv(serverPortEnvName, defaultServerPort)
	}

	return &args
}
