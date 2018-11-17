package main

import (
	"log"
	"net/http"
	"sandbox/sorajima/workhourmonthly"
)

func main() {
	rootURI := "/sorajima/v1"
	http.HandleFunc(rootURI+"/workhour/monthly", workhourmonthly.HTTPMethodHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
