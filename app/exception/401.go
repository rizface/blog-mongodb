package exception

import "net/http"

type Unauthorized struct {
	Err error
}

func (ua Unauthorized) CheckError() bool {
	return ua.Err != nil
}

func (ua Unauthorized) Code() int {
	return http.StatusUnauthorized
}

func (ua Unauthorized) Error() string {
	return ua.Err.Error()
}
