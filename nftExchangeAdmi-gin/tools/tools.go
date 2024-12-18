package tools

import (
	"nftExchangeAdmi-gin/config"
	"nftExchangeAdmi-gin/db/mySql"
	"nftExchangeAdmi-gin/db/redis"
	"nftExchangeAdmi-gin/middleware"
	"nftExchangeAdmi-gin/mq/rabbitmq"
	"nftExchangeAdmi-gin/task/xxljob"
	"nftExchangeAdmi-gin/tools/rest"
)

// LoadComponent 加载系统所需的各个组件
// 根据配置初始化日志、数据库、Redis、任务调度、JWT、RabbitMQ等组件
func LoadComponent(c config.Config) {
	// 初始化日志组件
	logger := c.Logger
	if logger.Path != "" {
		middleware.InitLogger(logger)
	}

	// 初始化数据库组件
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

	// 初始化 Redis 组件
	red := c.Redis
	if red.Password != "" {
		redis.CreateRedisClusterClient(red)
	}

	// 初始化 XXL-JOB 分布式任务调度组件
	if c.Xxljob.Admin.Addresses != "" {
		xxljob.InitXxjob(c.Xxljob)
	}

	// 初始化中间件（JWT 和加密）
	jwt := c.JWT
	secure := c.Secure
	middleware.InitJwtAuth(jwt, secure)
	middleware.InitEncryptionMiddleware(secure)

	// 初始化 Gin 框架
	gin := c.Gin
	rest.InitGin(gin)

	// 初始化 RabbitMQ 组件
	mq := c.MQ
	if mq.Rabbitmq.Addresses != "" {
		rabbitmq.NewRabbitMQ(mq)
	}
}
