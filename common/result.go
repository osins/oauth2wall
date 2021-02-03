package common

type Result struct {
	Success bool
	Message string
	Code    int
	Error   error
	Data    interface{}
}

func (r *Result) String() string {
	return r.Message
}

func (r *Result) SetError(e error) *Result {
	r.Error = e
	return r
}

func (r *Result) SetSuccess(s bool) *Result {
	r.Success = s
	return r
}

func (r *Result) SetCode(code int) *Result {
	r.Code = code
	return r
}

func (r *Result) SetData(data interface{}) *Result {
	r.Data = data
	return r
}

func NewResult(message string) *Result {
	return &Result{
		Message: message,
		Success: true,
		Code:    0,
	}
}

func NewResultSuccess(data interface{}) *Result {
	return &Result{
		Message: "",
		Success: true,
		Code:    0,
		Data:    data,
	}
}
