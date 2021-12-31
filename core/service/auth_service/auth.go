package auth_service

import request2 "blog-mongo/app/model/request"

type Auth interface {
	Register(request request2.Register) string
	Login(request request2.Login) map[string]string
}
