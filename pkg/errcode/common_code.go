/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: common_code
 * @Version: 1.0.0
 * @Date: 2023/3/15 08:09
 */

package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success                   = NewError(0, "success")
	ServerError               = NewError(10000000, "server internal error")
	InvalidParams             = NewError(10000001, "params error")
	NotFound                  = NewError(10000002, "not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "authentication failure, not found AppKey and AppSecert")
	UnauthorizedTokenError    = NewError(10000004, "authentication failure, Token error")
	UnauthorizedTokenTimeout  = NewError(10000005, "authentication failure, Token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "authentication failure, Token generate error")
	TooManyRequests           = NewError(10000007, "too many requests")
)

var codes = map[int]string{}

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("error code exists %s, please replacr other one.", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e // copy
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests

	}
	return http.StatusInternalServerError
}
