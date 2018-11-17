package main

import (
	"log"
	"net/http"

	"github.com/sorajima/workhour/monthly"
)

func main() {
	rootURI := "/sorajima/api/v1"
	http.HandleFunc(rootURI+"/workhour/monthly", monthly.HTTPMethodHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
