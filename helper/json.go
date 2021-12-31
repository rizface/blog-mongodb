package helper

import (
	"blog-mongo/app/model/response"
	"encoding/json"
	"net/http"
)

func JsonWriter(w http.ResponseWriter, code int, msg interface{},data interface{}) {
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response.Standard{
		Code:   code,
		Status: msg,
		Data:   data,
	})
}
