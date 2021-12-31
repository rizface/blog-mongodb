package response

type Standard struct {
	Code   int         `json:"code"`
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
}
