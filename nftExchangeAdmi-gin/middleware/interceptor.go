package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	errors2 "nftExchangeAdmi-gin/errors"
	"nftExchangeAdmi-gin/types"
	"reflect"
	"runtime"
	"strconv"
)

type WrapHandler struct {
	Fun     interface{}             `json:"fun"`     // 业务处理函数
	Produce string                  `json:"produce"` // 数据格式，如 application/json、application/xml 等
	Params  map[string]reflect.Type `json:"params"`  // 参数映射，key 为参数名，value 为参数的类型
}

var methodMapping = map[string]WrapHandler{} // 方法映射表，用于存储所有注册的处理方法

// runtimeFuncName 获取函数的全名
func runtimeFuncName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

// BeanWrapHandler 将 WrapHandler 添加到 methodMapping，并对函数进行包装
func BeanWrapHandler(wrapHandler WrapHandler) gin.HandlerFunc {
	name := runtimeFuncName(wrapHandler.Fun)
	methodMapping[name] = wrapHandler
	return wrapHandlerFun(wrapHandler.Fun)
}

// wrapHandlerFun 包装处理函数，返回 gin.HandlerFunc
func wrapHandlerFun(handlerFunc interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		handlerValue := reflect.ValueOf(handlerFunc)
		handlerType := handlerValue.Type()

		// 确保处理函数是函数类型
		if handlerType.Kind() != reflect.Func {
			sendErrorResponse(c, http.StatusInternalServerError, "处理函数无效")
			return
		}

		// 获取处理函数全名
		funcFullName := runtimeFuncName(handlerFunc)
		wHandler, ok := methodMapping[funcFullName]
		if !ok {
			sendErrorResponse(c, http.StatusInternalServerError, "未找到对应的 WrapHandler")
			return
		}

		var args []reflect.Value
		params := wHandler.Params // 获取参数映射

		if len(params) > 0 { // 如果有参数需要处理
			produce := wHandler.Produce
			for paramName, paramType := range params {
				// 动态解析类型
				arg, err := resolveAndBindParam(c, paramType, produce, paramName)
				if err != nil {
					logrus.Error(err)
					return // 错误已在绑定函数中处理
				}
				args = append(args, arg)
			}
		}

		// 调用处理函数并获取返回值
		results := handlerValue.Call(args)
		handleResults(c, results)
	}
}

// resolveAndBindParam 解析参数类型并绑定值
func resolveAndBindParam(c *gin.Context, paramType reflect.Type, produce, paramName string) (reflect.Value, error) {
	var arg reflect.Value

	switch produce {
	case "application/json":
		arg = bindBodyParam(c, paramType)
	case "application/x-www-form-urlencoded", "multipart/form-data":
		arg = bindQueryParam(c, paramType, paramName)
	default:
		sendErrorResponse(c, http.StatusUnsupportedMediaType, "不支持的数据格式")
		return reflect.Value{}, fmt.Errorf("unsupported media type: %s", produce)
	}

	if !arg.IsValid() {
		return reflect.Value{}, fmt.Errorf("参数绑定失败: %s", paramName)
	}

	return arg, nil
}

// sendErrorResponse 统一的错误响应函数
func sendErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"message": message,
		"success": false,
	})
}

// bindQueryParam 绑定 GET/DELETE 请求参数
func bindQueryParam(c *gin.Context, inType reflect.Type, paramName string) reflect.Value {
	argVal := reflect.New(inType.Elem())

	if inType.Elem().Kind() == reflect.Struct {
		if err := c.ShouldBindQuery(argVal.Interface()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数: " + err.Error()})
			return reflect.Value{}
		}
		logrus.Infof("成功绑定结构体查询参数: %v", argVal.Interface())
	} else {
		queryValue := c.Query(paramName)
		if queryValue == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("缺少查询参数: %s", paramName)})
			return reflect.Value{}
		}

		if err := setBasicType(argVal.Elem(), queryValue); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("无效的查询参数 '%s': %v", paramName, err)})
			return reflect.Value{}
		}
		logrus.Infof("绑定查询参数 '%s' = '%v'", paramName, queryValue)
	}
	return argVal
}

// bindBodyParam 绑定 POST/PUT 请求参数
func bindBodyParam(c *gin.Context, inType reflect.Type) reflect.Value {
	if inType.Kind() == reflect.Ptr && inType.Elem().Kind() == reflect.Struct {
		argVal := reflect.New(inType.Elem())
		if err := c.ShouldBindJSON(argVal.Interface()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
			return reflect.Value{}
		}
		return argVal
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "POST/PUT 请求需要单个结构体参数"})
	return reflect.Value{}
}

// handleResults 处理返回值
func handleResults(c *gin.Context, results []reflect.Value) {
	var (
		output interface{}
		myErr  error
	)

	switch len(results) {
	case 1:
		// 单返回值认为是错误
		if !results[0].IsNil() {
			errInterface := results[0].Interface()
			myErr, _ = errInterface.(error)
		}
	case 2:
		// 双返回值，第一个为数据，第二个为错误
		if !results[0].IsNil() {
			output = results[0].Interface()
		}
		if !results[1].IsNil() {
			errInterface := results[1].Interface()
			myErr, _ = errInterface.(error)
		}
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "处理函数返回值必须为一或两个"})
		return
	}

	if myErr != nil {
		handleError(c, myErr)
		return
	}

	c.JSON(http.StatusOK, types.CommonResult{
		Code:    0,
		Message: "成功",
		Data:    output,
		Success: true,
	})
}

// handleError 处理错误
func handleError(c *gin.Context, err error) {
	var myErr *errors2.MyError
	if errors.As(err, &myErr) {
		formattedMessage := formatErrorMessage(myErr.Message, myErr.Param)
		logrus.WithFields(logrus.Fields{
			"code":    myErr.Code,
			"message": formattedMessage,
			"params":  myErr.Param,
			"path":    c.Request.URL.Path,
			"client":  c.ClientIP(),
		}).Error("业务错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    myErr.Code,
			"message": formattedMessage,
			"success": false,
		})
	} else {
		logrus.WithFields(logrus.Fields{
			"path":   c.Request.URL.Path,
			"client": c.ClientIP(),
		}).Errorf("内部服务错误: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "内部服务错误",
			"success": false,
		})
	}
}

// formatErrorMessage 格式化错误信息
func formatErrorMessage(template string, params []any) string {
	if len(params) > 0 {
		return fmt.Sprintf(template, params...)
	}
	return template
}

// setBasicType 设置基本类型的参数值
func setBasicType(value reflect.Value, queryValue string) error {
	switch value.Kind() {
	case reflect.String:
		value.SetString(queryValue)
	case reflect.Int:
		intVal, err := strconv.Atoi(queryValue)
		if err != nil {
			return err
		}
		value.SetInt(int64(intVal))
	default:
		return fmt.Errorf("不支持的类型: %s", value.Kind())
	}
	return nil
}
