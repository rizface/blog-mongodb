package helper

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetParamsValue(r *http.Request, key string) string {
	params := mux.Vars(r)
	return params[key]
}

func GetQueryValue(r *http.Request,key string) string {
	return r.URL.Query().Get(key)
}