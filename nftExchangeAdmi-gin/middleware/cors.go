// middleware/cors.go

package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsMiddleware 返回一个 Gin 中间件函数，用于处理跨域请求
func CorsMiddleware() gin.HandlerFunc {
	// 使用自定义配置创建 CORS 中间件
	return cors.New(cors.Config{
		// 允许的请求源列表，可以添加多个域名
		AllowOrigins: []string{"*"}, // 出于安全考虑，建议指定具体的域名，例如：http://example.com

		// 允许的 HTTP 方法列表
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},

		// 允许的请求头列表
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},

		// 设置浏览器可以暴露的响应头
		ExposeHeaders: []string{"Content-Length", "Content-Type"},

		// 是否允许携带凭证信息（如 Cookie）
		AllowCredentials: true,

		// 预检请求的缓存时间，表示在此时间段内，无需再发送预检请求
		MaxAge: 12 * time.Hour,
	})
}
