package mySql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"nftExchangeAdmi-gin/config"
	"time"
)

var (
	BcDb *gorm.DB
)

func BcDbInit(sql config.MySql) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", sql.UserName, sql.Password, sql.Tcp, sql.DbName)
	var err error
	BcDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			&LogrusWriter{},
			logger.Config{
				SlowThreshold: time.Second, // 慢查询阈值
				LogLevel:      logger.Info, // 日志级别
				Colorful:      true,        // 彩色打印
			},
		),
	})
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}
	logrus.Info("connect database success")
}
