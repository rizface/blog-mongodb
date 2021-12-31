package exception

import "net/http"

type NotFound struct {
	Err error
}

func (nf NotFound) CheckError() bool {
	return nf.Err != nil
}

func (nf NotFound) Code() int {
	return http.StatusNotFound
}

func (nf NotFound) Error() string {
	return nf.Err.Error()
}
