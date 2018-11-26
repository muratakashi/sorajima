package monthly

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type errorStruct struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type getQueryParamStruct struct {
	userID int
	year   int
	month  int
}

func marshal(v interface{}) (b []byte, err error) {
	return json.MarshalIndent(v, "", " ")
}

func unMarshalGetQueryParam(values url.Values) (param getQueryParamStruct, errors []errorStruct) {
	userID := values.Get("userid")
	year := values.Get("year")
	month := values.Get("month")

	var err error
	param.userID, err = strconv.Atoi(userID)
	if err != nil {
		paramerr := errorStruct{
			Key:   "userid",
			Value: "ユーザーIDは不正です",
		}
		errors = append(errors, paramerr)
	}

	param.year, err = strconv.Atoi(year)
	if err != nil {
		paramerr := errorStruct{
			Key:   "year",
			Value: "年が数字ではありません",
		}
		errors = append(errors, paramerr)
	}

	param.month, err = strconv.Atoi(month)
	if err != nil {
		paramerr := errorStruct{
			Key:   "month",
			Value: "月が数字ではありません",
		}
		errors = append(errors, paramerr)
	} else {
		if param.month < 1 && param.month > 13 {
			paramerr := errorStruct{
				Key:   "month",
				Value: "月が不正です",
			}
			errors = append(errors, paramerr)
		}
	}

	return param, errors
}
