package types

type CommonResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Success bool   `json:"success"`
	Param   []any  `json:"param"` // 错误参数（可选）
}

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

// UserPO 定义用户请求参数
type UserPO struct {
	Name string `json:"name" binding:"required" `
	Age  int    `json:"age" binding:"required"`
}

// UserVO 定义用户响应数据
type UserVO struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
