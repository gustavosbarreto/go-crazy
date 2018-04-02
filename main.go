package main

import (
	"net/http"

	log "github.com/gustavosbarreto/gocrazy-log"
)

func main() {
	log.Info("starting")

	cli := http.Client{}

	req, _ := http.NewRequest("GET", "https://www.google.com", nil)

	cli.Do(req)

	log.Info("started")
}
