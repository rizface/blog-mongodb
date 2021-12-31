package auth_controller

import (
	request2 "blog-mongo/app/model/request"
	"blog-mongo/core/service/auth_service"
	"blog-mongo/helper"
	"encoding/json"
	"net/http"
)

type auth struct {
	service auth_service.Auth
}

func NewAuth(service auth_service.Auth) Auth{
	return &auth {
		service: service,
	}
}

func (a *auth) Register(w http.ResponseWriter, r *http.Request) {
	var request request2.Register
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Panic(err)
	result := a.service.Register(request)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (a *auth) Login(w http.ResponseWriter, r *http.Request) {
	var request request2.Login
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Panic(err)
	result := a.service.Login(request)
	helper.JsonWriter(w,http.StatusOK,"login sukses",result)
}

