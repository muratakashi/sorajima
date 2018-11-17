package workhourmonthly

import (
	"log"
	"net/http"
)

// HTTPMethodHandler HttpMethodに対する処理分岐を行う
func HTTPMethodHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Request Method  : ", req.Method)
	log.Println("Request URL     : ", req.URL)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")

	switch req.Method {
	case http.MethodGet:
		getHandler(w, req)
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// 指定した年月の勤務時間情報をJSON形式で返す
// Request
//   Method: Get
// Query Parameter
//   year: yyyy
//   month: mm
// Response
//   Content-Type : application/json
func getHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
