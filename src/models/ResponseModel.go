package models

type ErrorModel struct {
	ErrorCode        string `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}
type ResponseModel struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
	Error  ErrorModel  `json:"error"`
}

func (ResponseModel *ResponseModel) SetError(errorCode string, errorDescription string) {
	ResponseModel.Error = ErrorModel{
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
	}
}

