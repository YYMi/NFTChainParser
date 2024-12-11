package controller

import (
	"github.com/gin-gonic/gin"
)

// ControllerInterface 是所有控制器都要实现的接口
type ControllerInterface interface {
	RegisterRoutes(router *gin.Engine)
}
