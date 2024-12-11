package tools

import (
	"nftExchangeAdmi-gin/config"
	"nftExchangeAdmi-gin/db/mySql"
	"nftExchangeAdmi-gin/db/redis"
	"nftExchangeAdmi-gin/middleware"
	"nftExchangeAdmi-gin/task/xxljob"
	"nftExchangeAdmi-gin/tools/rest"
)

func LoadComponent(c config.Config) {
	logger := c.Logger
	if logger.Path != "" {
		middleware.InitLogger(logger)
	}
	source := c.DataSource
	if source.Debug {
		mySql.InitBebug(source.Debug)
	}
	if source.Bc.Tcp != "" {
		mySql.BcDbInit(source.Bc)
	}
	if source.BcSys.Tcp != "" {
		mySql.BcSysDInit(source.BcSys)
	}
	if source.Service.Tcp != "" {
		mySql.ServiceDbInit(source.Service)
	}
	red := c.Redis
	if red.Password != "" {
		redis.CreateRedisClusterClient(red)
	}
	if c.Xxljob.Admin.Addresses != "" {
		xxljob.InitXxjob(c.Xxljob)
	}
	jwt := c.JWT
	secure := c.Secure
	middleware.InitJwtAuth(jwt, secure)
	middleware.InitEncryptionMiddleware(secure)
	gin := c.Gin
	rest.InitGin(gin)
}
