package main

import (
	"log"
	"net/http"

	"github.com/sorajima/workhour/monthly"
)

func main() {
	rootURI := "/sorajima/api"
	http.HandleFunc(rootURI+"/v1/workhour/monthly", monthly.HTTPMethodHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
