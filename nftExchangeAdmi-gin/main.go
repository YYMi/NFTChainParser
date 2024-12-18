package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"nftExchangeAdmi-gin/config"
	_ "nftExchangeAdmi-gin/docs"
	"nftExchangeAdmi-gin/tools"
	"nftExchangeAdmi-gin/tools/rest"
)

// @title NFT Exchange Admin API
// @version 1.0
// @description API documentation for NFT Exchange Admin

// @host 192.168.10.66:19808
// @BasePath /

var configFile = flag.String("f", "/Users/yuyong/Documents/Golang/nftExchangeAdmi-gin/etc/nftexchangeadmin-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	config.MustLoad(*configFile, &c)
	tools.LoadComponent(c)
	rest.Start()
}
func createEthAddress() {
	// 1. 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("生成私钥失败: %v", err)
	}
	// 2. 获取私钥的十六进制表示
	privateKeyBytes := crypto.FromECDSA(privateKey) // 将私钥转换为字节数组
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	fmt.Printf("生成的私钥: %s\n", privateKeyHex)

	// 3. 生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法将公钥转换为 ECDSA")
	}

	// 4. 获取以太坊地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("生成的以太坊地址: %s\n", address)
}
