package main

import "fmt"

var (
	AppName   = "rep1"
	AppDesc   = "produces Leap Reporting data"
	Version   = "v0.0.1"
	Build     = "0011100"
	BuildDate = "<YYYY-MM-DD>"
	BuildTime = "<HH:MM:SS>"
)

func appTitle() string {
	return fmt.Sprintf("%s %s", AppName, AppDesc)
}

func fullVersion() string {
	return fmt.Sprintf("%s (Build %s) %s %s", Version, Build, BuildDate, BuildTime)
}
