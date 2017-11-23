package common

import (
	"encoding/json"
)

func init() {
	ResponseFactory[ErrInternalError] = &Response{
		Code: HTTP_INTERNAL_ERROR,
		Message: "Internal Error occurred, please retry later." +
			" If problem still exists, please contact us by ticket.",
	}
	ResponseFactory[ErrNotMatchCodeError] = &Response{
		Code:    HTTP_UNAUTHENTICATED,
		Message: "The request sever verify code not matched.",
	}
	ResponseFactory[ErrUserExistError] = &Response{
		Code:    HTTP_CONFLICT,
		Message: "The user you are trying to create already exists.",
	}
	ResponseFactory[ErrInvalidArgumentError] = &Response{
		Code:    HTTP_BAD_REQUEST,
		Message: "The argument is invalid.",
	}
	ResponseFactory[HTTP_OK] = &Response{
		Code:    HTTP_OK,
		Message: "OK",
	}
	ResponseFactory[HTTP_OK_CREATED] = &Response{
		Code:    HTTP_OK_CREATED,
		Message: "Created successfully!",
	}
	ResponseFactory[HTTP_METHOD_NOT_ALLOWED] = &Response{
		Code:    HTTP_METHOD_NOT_ALLOWED,
		Message: "The request method is not allowed",
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (res *Response) ToJson() []byte {
	data, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return data
}

var ResponseFactory = make(map[interface{}]*Response)
