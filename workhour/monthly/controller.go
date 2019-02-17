package monthly

import (
	"log"
	"net/http"
)

// HTTPMethodHandler HttpMethodに対する処理分岐を行う
func HTTPMethodHandler(w http.ResponseWriter, req *http.Request) {
	// CORF対応
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
//   user_id: ユーザーID
//   year: yyyy
//   month: mm
// Response
//   Content-Type : application/json
func getHandler(w http.ResponseWriter, req *http.Request) {
	// クエリパラメータを取得
	param, paramerrs := unMarshalGetQueryParam(req.URL.Query())
	if paramerrs != nil {
		b, err := marshal(&paramerrs)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content, err := getMonthly(param.userID, param.year, param.month)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if content == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
	w.WriteHeader(http.StatusOK)
}
