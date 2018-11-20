package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)

	log.Trace("Trace")
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	log.Fatal("Fatal")
}
