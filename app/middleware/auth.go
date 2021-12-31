package middleware

import (
	"blog-mongo/app/exception"
	"blog-mongo/helper"
	"context"
	"errors"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		bearer := request.Header.Get("Authorization")
		helper.PanicCustomException(exception.Unauthorized{Err:errors.New("token kamu tidak ada")},strings.Contains(bearer,"Bearer") == false)
		items := strings.Split(bearer, " ")
		helper.PanicCustomException(exception.Unauthorized{Err:errors.New("yang kamu kirim bukan token")},len(items) != 2)
		claims := helper.TokenValidation(items[1])
		request = request.WithContext(context.WithValue(request.Context(),"userData",claims))
		next.ServeHTTP(writer,request)
	})
}
