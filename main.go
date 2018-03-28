package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("starting")

	cli := http.Client{}

	req, _ := http.NewRequest("GET", "https://www.google.com", nil)

	cli.Do(req)

	logrus.Info("started")
}
