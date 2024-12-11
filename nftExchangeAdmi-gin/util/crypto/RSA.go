package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

// GenerateRSAKeyPair 生成 RSA 密钥对
// 参数 bits: int - 密钥的位数，只能是 1024、2048、3072 或 4096
// 返回 *rsa.PrivateKey: 生成的私钥
// 返回 *rsa.PublicKey: 生成的公钥（从私钥中派生）
// 返回 error: 错误信息（如果有）
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	validBits := []int{1024, 2048, 3072, 4096}
	isValid := false
	for _, b := range validBits {
		if bits == b {
			isValid = true
			break
		}
	}
	if !isValid {
		return nil, nil, fmt.Errorf("无效的密钥长度。支持的长度：%v", validBits)
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// EncryptWithPublicKey 使用公钥加密
// 参数 msg: string - 需要加密的明文消息（字符串）
// 参数 pub: *rsa.PublicKey - 用于加密的 RSA 公钥
// 返回 string: 加密后的密文（Base64 编码）
// 返回 error: 错误信息（如果有）
func EncryptWithPublicKey(msg string, pub *rsa.PublicKey) (string, error) {
	// 将字符串转换为字节数组
	msgBytes := []byte(msg)

	// 计算 RSA 可加密的最大块大小
	keySize := pub.Size()        // 公钥模长（字节数）
	maxChunkSize := keySize - 11 // PKCS#1 v1.5 填充占用 11 字节

	var encryptedBytes []byte
	for start := 0; start < len(msgBytes); start += maxChunkSize {
		end := start + maxChunkSize
		if end > len(msgBytes) {
			end = len(msgBytes)
		}

		// 对每个块进行加密
		chunk := msgBytes[start:end]
		encryptedChunk, err := rsa.EncryptPKCS1v15(rand.Reader, pub, chunk)
		if err != nil {
			return "", err
		}

		encryptedBytes = append(encryptedBytes, encryptedChunk...)
	}

	// 对加密后的数据进行 Base64 编码
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

// DecryptWithPrivateKey 使用私钥解密
// 参数 ciphertextBase64: string - 需要解密的密文（Base64 编码）
// 参数 private: *rsa.PrivateKey - 用于解密的 RSA 私钥
// 返回 string: 解密后的明文
// 返回 error: 错误信息（如果有）
func DecryptWithPrivateKey(ciphertextBase64 string, private *rsa.PrivateKey) (string, error) {
	// 将 Base64 编码的密文解码为字节数组
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", err
	}

	// 计算 RSA 可解密的块大小
	keySize := private.Size()
	var decryptedBytes []byte
	for start := 0; start < len(ciphertext); start += keySize {
		end := start + keySize
		if end > len(ciphertext) {
			end = len(ciphertext)
		}

		// 对每个块进行解密
		chunk := ciphertext[start:end]
		decryptedChunk, err := rsa.DecryptPKCS1v15(rand.Reader, private, chunk)
		if err != nil {
			return "", err
		}

		decryptedBytes = append(decryptedBytes, decryptedChunk...)
	}

	return string(decryptedBytes), nil
}

// SignWithPrivateKey 使用私钥对消息进行签名（相当于加密）
// 参数 msg: string - 需要签名的明文消息（字符串）
// 参数 private: *rsa.PrivateKey - 用于签名的 RSA 私钥
// 返回 string: 签名后的消息（Base64 编码）
// 返回 error: 错误信息（如果有）
func SignWithPrivateKey(msg string, private *rsa.PrivateKey) (string, error) {
	// 将字符串转换为字节数组
	msgBytes := []byte(msg)

	// 对消息进行 SHA-256 哈希处理，生成消息摘要
	hash := sha256.New()
	hash.Write(msgBytes)
	digest := hash.Sum(nil)

	// 使用私钥对消息摘要进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA256, digest)
	if err != nil {
		return "", err
	}

	// 对签名结果进行 Base64 编码
	encoded := base64.StdEncoding.EncodeToString(signature)
	return encoded, nil
}

// DecryptWithPublicKey 使用公钥验证签名（相当于解密）
// 参数 signatureBase64: string - 需要验证的签名（Base64 编码）
// 参数 msg: string - 原始的明文消息（字符串）
// 参数 pub: *rsa.PublicKey - 用于验证签名的 RSA 公钥
// 返回 error: 如果签名验证失败则返回错误信息
func DecryptWithPublicKey(signatureBase64 string, msg string, pub *rsa.PublicKey) error {
	// 将 Base64 编码的签名解码为字节数组
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return err
	}

	// 将消息字符串转换为字节数组并生成 SHA-256 哈希
	msgBytes := []byte(msg)
	hash := sha256.New()
	hash.Write(msgBytes)
	digest := hash.Sum(nil)

	// 使用公钥验证签名，如果签名不匹配则返回错误
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, digest, signature)
	if err != nil {
		return err
	}
	return nil
}

// RsaPrivateKeyToString 将私钥转换为字符串
func RsaPrivateKeyToString(private *rsa.PrivateKey) string {
	privateDer := x509.MarshalPKCS1PrivateKey(private)
	privateBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateDer,
	}
	key := string(pem.EncodeToMemory(&privateBlock))
	// 去掉头部和尾部标识符
	keyLines := strings.Split(key, "\n")
	keyBody := strings.Join(keyLines[1:len(keyLines)-2], "")
	return keyBody
}

// RsaPublicKeyToString 将公钥转换为字符串
func RsaPublicKeyToString(pub *rsa.PublicKey) string {
	pubDER, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		panic(err)
	}
	pubBlock := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubDER,
	}
	key := string(pem.EncodeToMemory(&pubBlock))
	keyLines := strings.Split(key, "\n")
	keyBody := strings.Join(keyLines[1:len(keyLines)-2], "")
	return keyBody
}

// ParsePrivateKeyFromString 解析 PEM 格式的私钥字符串为 rsa.PrivateKey 对象
// 参数 keyStr: string - 包含 PEM 格式的私钥字符串
// 返回: *rsa.PrivateKey, error
func ParsePrivateKeyFromString(keyStr string) (*rsa.PrivateKey, error) {
	// 去除首尾的空白字符
	keyStr = strings.TrimSpace(keyStr)

	// 如果私钥字符串不包含 "BEGIN PRIVATE KEY" 和 "END PRIVATE KEY"，就补充这些标识符
	if !strings.HasPrefix(keyStr, "-----BEGIN PRIVATE KEY-----") {
		keyStr = "-----BEGIN PRIVATE KEY-----\n" + keyStr
	}
	if !strings.HasSuffix(keyStr, "-----END PRIVATE KEY-----") {
		keyStr = keyStr + "\n-----END PRIVATE KEY-----"
	}

	// 解析 PEM 格式的私钥块
	block, _ := pem.Decode([]byte(keyStr))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}

	// 尝试解析 PKCS1 格式的私钥
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return private, nil
	}

	// 如果 PKCS1 解析失败，则尝试解析 PKCS8 格式的私钥
	privateInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err // 如果解析 PKCS8 也失败，则返回错误
	}

	// 转换为 rsa.PrivateKey 类型
	if rsaPrivateKey, ok := privateInterface.(*rsa.PrivateKey); ok {
		return rsaPrivateKey, nil
	}

	return nil, errors.New("private key is not an RSA key")
}

// ParsePublicKeyFromString 解析 PEM 格式的公钥字符串为 *rsa.PublicKey 对象
// 参数 keyStr: string - 包含 PEM 格式的公钥字符串
// 返回: *rsa.PublicKey, error
func ParsePublicKeyFromString(keyStr string) (*rsa.PublicKey, error) {
	// 去除首尾的空白字符
	keyStr = strings.TrimSpace(keyStr)

	// 检查并补充 PEM 标识符
	if !strings.HasPrefix(keyStr, "-----BEGIN") {
		keyStr = "-----BEGIN PUBLIC KEY-----\n" + keyStr
	}
	if !strings.HasSuffix(keyStr, "-----END PUBLIC KEY-----") {
		keyStr = keyStr + "\n-----END PUBLIC KEY-----"
	}

	// 解析 PEM 格式的公钥块
	block, _ := pem.Decode([]byte(keyStr))
	if block == nil {
		return nil, errors.New("无法解析 PEM 块中的公钥")
	}

	var pubInterface interface{}
	var err error

	switch block.Type {
	case "PUBLIC KEY":
		// 尝试解析 PKIX 或 PKCS#8 格式的公钥（一般情况下的公钥）
		pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, errors.New("无法解析 PKIX 公钥：" + err.Error())
		}
	case "RSA PUBLIC KEY":
		// 尝试解析 PKCS#1 格式的公钥
		pubInterface, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, errors.New("无法解析 PKCS#1 公钥：" + err.Error())
		}
	case "CERTIFICATE":
		// 如果是证书，尝试从证书中提取公钥
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, errors.New("无法解析证书：" + err.Error())
		}
		pubInterface = cert.PublicKey
	default:
		return nil, errors.New("不支持的公钥类型：" + block.Type)
	}

	// 类型断言为 *rsa.PublicKey
	if pub, ok := pubInterface.(*rsa.PublicKey); ok {
		return pub, nil
	} else {
		return nil, errors.New("公钥不是 RSA 类型")
	}
}
