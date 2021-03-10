package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Warn("A walrus appears")

	log.Info("End of msg")
	log.Error("some error has occurred.")
}
