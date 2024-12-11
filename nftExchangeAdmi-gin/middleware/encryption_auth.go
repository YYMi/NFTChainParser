// middleware/encryption.go

package middleware

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"nftExchangeAdmi-gin/config"
	"nftExchangeAdmi-gin/util/crypto"
	"time"

	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
)

var (
	encryptedPaths []string
	privateKey     *rsa.PrivateKey
	publicKey      *rsa.PublicKey
	privateKeyStr  = "MIICXgIBAAKBgQDLlcZ1qkurVQrixC7+jYe85OtefcNdjlTZbC6dOygX/vS/UXbQMP0PTkePdSB1A/rsdW5pLnhG4XgyzV+ew3l1oVqR/qNWFJt3PzWfuPyqATArUZqKXR/bGPzLypkpijW9qkM2/tZBpuiO//I7L6hTAa47g1VUUh5XCno8j4Ju3wIDAQABAoGBAKiNCTFT27AvCYMzb8D6hj4CvUePEdd8Ro13/qPYXVp4kENxe8/kLy+j3KVOEKAwumdY1h+pBJWSiIRu+lKkfgJExCxDgExEuieJD5TrPoIjdT4ciFUvHu32Ja3Pk8a67M/bQrj87COMBuuXUvBZW5xf9sYBIV30HyJWlV2eA5hBAkEA/Ap3tEQZNClrZGDNQEyIfMttLjA2fMSzyWYADaivXM4QisOyOzlnj8FCViAVmlUbtRBhahAp/zwakJ61hj9soQJBAM7Ic79YcxQKb3LKraW8CaJdqudWaLunQZ0vTxpr1xtQLSdC6wMQ5EMMch6KLG6+yRYR4AtCpnIOTlI8eWr8q38CQQCtZwI+Ws/ATHLfZ54vC7inq0mdinwiUS6kdHG69ABaaEeHQOaBypfOlpbuLDrQqJdcdj2fPCm4uYBJWXfoOgChAkB5fcqLzMLoLm1mi3BU5NLgQ5pLCzjDsDRbztTyGBQtJwEwhR+hEIacYi6WhOYwNwXcYqo403MJpiAcLw3DbyqvAkEAgAllrOd8iDOuhMGQK+Efgd0OnRnwtSOhuY3niiA5npIlu9jBoqzWk0ujzEbgsxkwXtxRhqGfWIHifIFkyrF77g=="
	publicKeyStr   = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDLlcZ1qkurVQrixC7+jYe85OtefcNdjlTZbC6dOygX/vS/UXbQMP0PTkePdSB1A/rsdW5pLnhG4XgyzV+ew3l1oVqR/qNWFJt3PzWfuPyqATArUZqKXR/bGPzLypkpijW9qkM2/tZBpuiO//I7L6hTAa47g1VUUh5XCno8j4Ju3wIDAQAB"
	timestamp      int64
)

// InitEncryptionMiddleware 初始化加密中间件，加载 RSA 密钥和加密路径列表
func InitEncryptionMiddleware(cfg config.Secure) {
	encryptedPaths = cfg.Encrypted.URLs
	timestamp = cfg.Encrypted.Timestamp
	privateKey, _ = crypto.ParsePrivateKeyFromString(privateKeyStr)
	publicKey, _ = crypto.ParsePublicKeyFromString(publicKeyStr)
	logrus.Infof("私钥%s", privateKeyStr)
	logrus.Infof("公钥%s", publicKeyStr)

}

// EncryptionMiddleware 返回一个 Gin 中间件，用于处理加密和解密
// EncryptionMiddleware 返回一个 Gin 中间件，用于处理加密和解密
func EncryptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestPath := c.Request.URL.Path
		if isPathEncrypted(requestPath) {
			// 1. 读取加密的请求体
			encryptedBodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "读取请求体失败"})
				return
			}

			// 将请求体转换为字符串（Base64 编码的密文）
			encryptedBody := string(encryptedBodyBytes)
			key, err := crypto.EncryptWithPublicKey(encryptedBody, publicKey)
			if err != nil {
				logrus.Error(err)
			}
			logrus.Infof("公钥加密结果 %s", key)

			// 2. 使用私钥解密请求数据
			decryptedBody, err := crypto.DecryptWithPrivateKey(encryptedBody, privateKey)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "解密请求体失败"})
				return
			}

			// 3. 校验请求的时间戳
			var requestData struct {
				Timestamp int64           `json:"Timestamp"` // 请求的时间戳（毫秒）
				Data      json.RawMessage `json:"data"`      // 其他请求参数
			}

			// 解析解密后的 JSON 数据
			if err := json.Unmarshal([]byte(decryptedBody), &requestData); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "请求数据格式错误"})
				return
			}

			// 获取当前时间戳（毫秒）
			currentTimestamp := time.Now().UnixMilli()

			// 校验时间戳是否在有效范围
			if currentTimestamp-requestData.Timestamp > timestamp || requestData.Timestamp > currentTimestamp {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "请求时间无效"})
				return
			}
			// 打印请求的时间戳（转换为年月日-时分秒）
			requestTime := time.UnixMilli(requestData.Timestamp).Format("2006-01-02 15:04:05")
			logrus.Infof("请求有效时间校验通过，请求发起时间：%s", requestTime)

			// 4. 替换请求体为解密后的 Data 部分（如果 Data 存在）
			if len(requestData.Data) > 0 { // 只有当 Data 存在时替换请求体
				dataBody, err := json.Marshal(requestData.Data)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "处理请求数据失败"})
					return
				}
				c.Request.Body = io.NopCloser(bytes.NewBuffer(dataBody))
				// 修正 Content-Length
				c.Request.Header.Set("Content-Length", strconv.Itoa(len(dataBody)))
			}

			// 5. 包装响应写入器，捕获响应数据
			writer := &responseWriterWrapper{
				ResponseWriter: c.Writer,        // 使用原ResponseWriter初始化
				body:           &bytes.Buffer{}, // 初始化一个空的缓冲区
			}
			c.Writer = writer // 将上下文中的Writer替换为我们自定义的writer
			defer func() {
				writer.body.Reset() // 重置缓冲区以备后续请求使用
			}()
			c.Next()
			// 获取响应数据
			responseData := writer.body.Bytes()

			// 使用私钥生成签名
			signatureBase64, err := crypto.SignWithPrivateKey(string(responseData), privateKey)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "签名生成失败"})
				return
			}

			// 清空之前写入的响应
			c.Writer = writer.ResponseWriter // 重置 Writer 确保覆盖
			// 返回仅包含签名的响应
			c.Writer.Header().Set("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Write([]byte(`{"signature":"` + signatureBase64 + `"}`))
			// 中止后续逻辑
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// 定义responseWriterWrapper结构体，用于封装gin.ResponseWriter并添加缓冲区以存储响应体内容
type responseWriterWrapper struct {
	gin.ResponseWriter               // 继承gin.ResponseWriter接口
	body               *bytes.Buffer // 使用字节缓冲区存储响应体
}

// 重写Write方法，将响应体内容写入缓冲区
func (w *responseWriterWrapper) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

// isPathEncrypted 检查请求路径是否需要加密
func isPathEncrypted(requestPath string) bool {
	for _, pattern := range encryptedPaths {
		regexPattern := wildcardToRegex(pattern)
		matched, _ := regexp.MatchString(regexPattern, requestPath)
		if matched {
			return true
		}
	}
	return false
}
