package main

import (
	"gofmanaa/idn_checking/client"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	log.Infoln("Start.")

	conf, err := client.GetConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	dateRange := createDataRange()
	for _, date := range dateRange {
		log.Infof("make request on %s ...", date)
		client.PostRequest(log, date, conf)
		log.Infoln("sleep...")
		time.Sleep(time.Second * 30)
	}
	log.Infoln("Stop.")
}

func createDataRange() []string {
	var dateRange []string
	day := time.Hour * 24
	nextDay := time.Now().Add(day)

	for durationDay := 0; durationDay <= 14; durationDay++ {
		newDate := nextDay.Add(day * time.Duration(durationDay))
		dateRange = append(dateRange, newDate.Format("2006-01-02"))
	}

	return dateRange
}

func initLogger() *logrus.Logger {
	var log = logrus.New()
	file, err := os.OpenFile("logfile.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "02-01-2006 15:04:05",
	})
	if err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		log.SetOutput(mw)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	return log
}
