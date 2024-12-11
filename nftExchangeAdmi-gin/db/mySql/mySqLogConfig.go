package mySql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// LogrusWriter 用于将 GORM 的日志写入到 logrus
type LogrusWriter struct{}

// Printf 实现 logger.Writer 接口
var debug = false

func InitBebug(_debug bool) {
	debug = _debug
}

// Printf 实现 logger.Writer 接口
func (lw *LogrusWriter) Printf(format string, args ...interface{}) {
	if len(args) < 4 {
		// 确保参数数量足够，避免越界错误
		//logrus.Warn("日志参数不足")
		logrus.Error(args)
		return
	}

	// 提取日志参数
	logPath := fmt.Sprintf("%v", args[0]) // 文件路径

	// 尝试获取相对路径
	relativePath := logPath
	if idx := strings.Index(logPath, "NftExchange"); idx != -1 {
		relativePath = logPath[idx:]
	}
	if !debug {
		logrus.Info(fmt.Sprintf("SQL Query Execution Path: %s", relativePath))
		return

	} else {
		runtime := fmt.Sprintf("%v", args[1]) // 运行时间
		result := fmt.Sprintf("%v", args[2])  // 查询结果
		query := fmt.Sprintf("%v", args[3])   // 查询语句
		// 自定义日志输出格式
		output := fmt.Sprintf(
			" [ %s ]\n\tRUNTIME-[%sms]\n\tRESULT-[%s]\n\t%s",
			relativePath,
			runtime,
			result,
			query,
		)
		// 使用 logrus 输出日志
		logrus.Info(output)
	}
}
