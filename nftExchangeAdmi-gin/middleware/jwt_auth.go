// middleware/jwt_auth.go

package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"nftExchangeAdmi-gin/config"
	"regexp"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	whiteListedPaths []string
	jwtConfig        config.JWT
	JwtStatus        bool = false
)

func InitJwtAuth(jwt2 config.JWT, secure config.Secure) {
	jwtConfig = jwt2
	whiteListedPaths = secure.Ignored.URLs
}

// JWTAuthMiddleware 返回一个 Gin 中间件函数，用于验证 JWT 令牌
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求路径
		requestPath := c.Request.URL.Path

		// 检查请求路径是否在白名单中
		if isPathWhitelisted(requestPath) {
			c.Next()
			return
		}

		// 从全局配置中获取 TokenHeader 和 TokenHead
		tokenHeader := jwtConfig.TokenHeader
		tokenHead := jwtConfig.TokenHead

		// 获取请求头中的 Token
		authHeader := c.GetHeader(tokenHeader)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "未提供 " + tokenHeader + " 请求头",
			})
			logrus.Error("请求方式权限未认证: 未提供 " + tokenHeader + "请求头")
			return
		}

		// 按空格分割，通常格式为 "<TokenHead> <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == tokenHead) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": tokenHeader + " 请求头格式错误",
			})
			return
		}

		// 使用全局配置中的密钥解析 JWT
		tokenString := parts[1]
		secretKey := jwtConfig.Secret

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法是否正确
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "无效或过期的令牌",
			})
			return
		}

		// 可选：将令牌中的声明保存到上下文中
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("claims", claims)
		}

		// 继续处理请求
		c.Next()
	}
}

// isPathWhitelisted 检查请求路径是否在白名单中，支持通配符和正则表达式
func isPathWhitelisted(requestPath string) bool {
	for _, pattern := range whiteListedPaths {
		// 将通配符模式转换为正则表达式
		regexPattern := wildcardToRegex(pattern)
		matched, _ := regexp.MatchString(regexPattern, requestPath)
		if matched {
			return true
		}
	}
	return false
}
