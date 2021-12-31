package helper

import (
	"blog-mongo/app/exception"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}


func PanicCustomException(exception exception.Exception, condition bool) {
	if condition == true {
		panic(exception)
	}
}