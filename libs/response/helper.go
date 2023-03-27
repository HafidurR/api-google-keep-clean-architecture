package response

type NoDataResponse struct {
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

func (o *NoDataResponse) SetData(v interface{}) {
	o.Data = v
}

func (o *NoDataResponse) SetMessage(v string) {
	o.Message = v
}

type ErrorResponse struct {
	Data interface{} `json:"data"`
	Message string `json:"message"`
}