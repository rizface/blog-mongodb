package exception

import "net/http"

type BadRequest struct {
	Err error
}

func (b BadRequest) CheckError() bool {
	return b.Err != nil
}

func (b BadRequest) Code() int {
	return http.StatusBadRequest
}

func (b BadRequest) Error() string {
	return b.Err.Error()
}
