package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"nftExchangeAdmi-gin/controller"
	"reflect"
)

// InitRoutes 自动注册控制器
func InitRoutes(router *gin.Engine) {
	// 初始化你所有的控制器类型
	controllerTypes := []interface{}{
		&controller.UserController{},
		// 添加更多的控制器类型，比如 &controllers.ProductController{}
	}
	for _, ctrl := range controllerTypes {
		controllerValue := reflect.ValueOf(ctrl)
		method := controllerValue.MethodByName("RegisterRoutes")
		if method.IsValid() {
			args := []reflect.Value{reflect.ValueOf(router)}
			method.Call(args)
			logrus.Infof("Registered routes for controller: %T", ctrl)
		} else {
			logrus.Infof("No RegisterRoutes method found for controller: %T", ctrl)
		}
	}
}
