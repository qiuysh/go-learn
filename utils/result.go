package utils

type Result struct {
	// Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

const (
	// 成功
	RESULT_SUCCESS = true
	// 失败
	RESULT_FAILURE = false
)

func (res *Result) SetSuccess(success bool) *Result {
	res.Success = success
	return res
}

// func (res *Result) SetCode(code int) *Result {
// 	res.Code = code
// 	return res
// }

func (res *Result) SetMessage(msg string) *Result {
	res.Message = msg
	return res
}

func (res *Result) SetData(data interface{}) *Result {
	res.Data = data
	return res
}
