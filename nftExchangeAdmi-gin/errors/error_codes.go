package errors

// ErrorCode 定义错误代码结构
type ErrorCode struct {
	Code    int64
	Message string
}

// 定义错误代码常量
var (
	SUCCESS         = ErrorCode{Code: 0, Message: "成功"}
	FAILED          = ErrorCode{Code: 500, Message: "失败"}
	VALIDATE_FAILED = ErrorCode{Code: 404, Message: "参数检验失败"}
	UNAUTHORIZED    = ErrorCode{Code: 401, Message: "暂未登录或token已经过期"}
	FORBIDDEN       = ErrorCode{Code: 403, Message: "没有相关权限"}
	RESUBMIT        = ErrorCode{Code: 409, Message: "重复提交"}
	OTHER           = ErrorCode{Code: -1, Message: "未知错误"}
)
