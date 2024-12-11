package main

import (
	"flag"
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
