package middleware

import (
	exception2 "blog-mongo/app/exception"
	"blog-mongo/helper"
	"fmt"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				exception,ok := err.(exception2.Exception)
				fmt.Println(exception.Error())
				if ok {
					helper.JsonWriter(writer,exception.Code(),exception.Error(),nil)
				} else {
					error := err.(error)
					helper.JsonWriter(writer,http.StatusInternalServerError,error.Error(),nil)
				}
			}
		}()
		next.ServeHTTP(writer,request)
	})
}
