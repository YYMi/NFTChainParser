package controller

import (
	"github.com/gin-gonic/gin"
	"nftExchangeAdmi-gin/errors"
	"nftExchangeAdmi-gin/middleware"
	"nftExchangeAdmi-gin/service"
	"nftExchangeAdmi-gin/types"
	"reflect"
)

type UserController struct {
	userService *service.UserService
}

// RegisterRoutes 注册 NFT 集合相关的路由
func (ctrl *UserController) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/chainLink/Name") // 11 用户管理
	{
		// 获得用户
		group.GET("/users", middleware.BeanWrapHandler(middleware.WrapHandler{
			Fun:     ctrl.GetAllProducts,
			Produce: "multipart/form-data",
			Params: map[string]reflect.Type{
				"name": reflect.TypeOf((*string)(nil)), // 指针类型
				"aget": reflect.TypeOf((*int)(nil)),    // 指针类型
			},
		}))

		// 获得所有的用户列表
		group.POST("/userList", middleware.BeanWrapHandler(middleware.WrapHandler{
			Fun:     ctrl.UserList,
			Produce: "application/json",
			Params:  map[string]reflect.Type{"po": reflect.TypeOf((*types.UserPO)(nil))}, // 无参数
		}))
	}
}

// GetAllProducts
// @Summary 获得用户
// @Tags 11 用户管理
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "请求参数-1"
// @Param aget formData int true "请求参数-2"
// @Success 200 {object} types.UserVO
// @Router /chainLink/Name/users [GET]
func (ctrl *UserController) GetAllProducts(name *string, aget *int) (*types.UserVO, *errors.MyError) {
	if name == nil || aget == nil {
		return nil, &errors.MyError{
			Code:    400,
			Message: "缺少必要参数",
		}
	}
	// 调用服务层逻辑
	return ctrl.userService.GetAllProducts(name, aget)
}
// UserList
// @Summary 获得所有的用户列表
// @Tags 11 用户管理
// @Accept json
// @Produce json
// @Param po body types.UserPO true  "请求参数-1"
// @Success 200 {object} bool
// @Router /chainLink/Name/userList [POST]
func (ctrl *UserController) UserList(po *types.UserPO) (*bool, *errors.MyError) {
	// 假设服务层方法返回一个布尔值，表示是否成功
	success := true
	return &success, nil
}
