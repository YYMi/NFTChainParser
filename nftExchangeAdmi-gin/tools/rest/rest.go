package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"math"
	"nftExchangeAdmi-gin/config"
	"nftExchangeAdmi-gin/middleware"
	"nftExchangeAdmi-gin/router"
	"time"
)

// 定义全局变量
var engine *gin.Engine
var port string = ":8080"

// InitGin 初始化 Gin 框架
func InitGin(cfg config.Gin) {
	// 根据配置设置模式
	switch cfg.Active {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test", "qa":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	engine = gin.New() // 生产环境，创建一个新的 Gin 实例

	port = ":" + cfg.Port

	engine.Use(setupLogrusLogger())
	//// 添加自定义的 CORS 中间件，处理跨域请求
	engine.Use(middleware.CorsMiddleware())
	engine.Use(middleware.JWTAuthMiddleware())
	engine.Use(middleware.EncryptionMiddleware())
	// Swagger 路由
	//  注册路由
	router.InitRoutes(engine)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// setupLogrusLogger 设置 Gin 使用 logrus 作为日志记录器
func setupLogrusLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求的开始时间
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		// 请求信息
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		// 示例日志打印，添加了向上取整功能
		latencyInMs := math.Ceil(float64(latency.Milliseconds())) // 将时间向上取整，单位是毫秒
		logrus.Infof("[%3d] [%vms] [%s] [%s] %v",
			statusCode,
			latencyInMs,
			clientIP,
			method,
			path,
		)
	}
}

// Start 启动 Gin 服务
func Start() {
	if engine == nil {
		// 如果 engine 未初始化，则记录错误日志并退出程序
		logrus.Fatal("Router is not initialized. Call InitGin() first.")
	}
	// 启动服务器并监听指定端口
	err := engine.Run(port)
	if err != nil { // 启动失败，记录错误日志并退出程序
		logrus.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}
