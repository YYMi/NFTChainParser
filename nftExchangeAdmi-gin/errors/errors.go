package errors

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

// MyError 自定义错误类型
type MyError struct {
	Code    int64  `json:"code"`    // 错误代码
	Message string `json:"message"` // 错误信息
	Param   []any  `json:"param"`   // 错误参数（可选）
}

// Error 实现 error 接口，返回错误的详细信息
func (e *MyError) Error() string {
	if len(e.Param) > 0 {
		return fmt.Sprintf("Code %d: %s, Params: %v", e.Code, e.Message, e.Param)
	}
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

// logError 打印错误信息和调用堆栈
func logError(err *MyError) {
	stack := captureStackTrace(3) // 捕获堆栈信息
	logrus.Error("Error occurred: %s Stack trace: %s", err.Error(), stack)
}

// captureStackTrace 捕获调用堆栈信息
func captureStackTrace(skip int) string {
	var stackTrace string
	pcs := make([]uintptr, 30) // 最多获取 10 层调用栈
	n := runtime.Callers(skip, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		stackTrace += fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
	return stackTrace
}

// SysError 跑出系统错误
func SysError(err error) *MyError {
	return &MyError{
		Code:    500,
		Message: "FAILED", // 包含系统错误信息
	}
}

// Error 抛出自定义错误代码
func Error(err string) *MyError {
	return &MyError{
		Code:    500,
		Message: err,
	}
}

// ErrorResult 根据 ErrorCode 生成错误
func ErrorResult(eCode ErrorCode) *MyError {
	return &MyError{
		Code:    eCode.Code,
		Message: eCode.Message,
	}
}

// ErrorResultParam 根据 ErrorCode 和参数生成错误
func ErrorResultParam(eCode ErrorCode, param []any) *MyError {
	return &MyError{
		Code:    eCode.Code,
		Message: eCode.Message,
		Param:   param,
	}
}
