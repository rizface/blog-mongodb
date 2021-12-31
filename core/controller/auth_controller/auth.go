package auth_controller

import "net/http"

type Auth interface {
	Register(w http.ResponseWriter,r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
