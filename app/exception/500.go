package exception

import (
	"net/http"
)

type InternalError struct {
	Err error
}

func (ie InternalError) CheckError() bool {
	return ie.Err != nil
}

func (ie InternalError) Code() int {
	return http.StatusInternalServerError
}

func (ie InternalError) Error() string {
	return ie.Err.Error()
}
