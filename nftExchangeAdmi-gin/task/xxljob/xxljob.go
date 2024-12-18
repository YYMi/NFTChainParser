package xxljob

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xxl-job/xxl-job-executor-go"
	"nftExchangeAdmi-gin/config"
)

// 自定义日志器
type logger struct{}

var exec xxl.Executor

func (l *logger) Info(format string, a ...interface{}) {
	if format == "任务参数:%v" && a != nil {
		// 将参数 a[0] 转换为 *xxl.RunReq 类型
		param, ok := a[0].(*xxl.RunReq)
		if !ok {
			logrus.Error("参数类型转换失败，期望 *xxl.RunReq")
			return
		}

		// 拼接完整日志消息
		fullMessage := fmt.Sprintf("Xxjob func %s 任务参数 %s", param.ExecutorHandler, param.ExecutorParams)
		logrus.Info(fullMessage)
	}
}

func (l *logger) Error(format string, a ...interface{}) {
	logrus.Errorf(format, a...)
}
func InitXxjob(cof config.Xxljob) {
	// 配置 xxl-job 执行器
	exec = xxl.NewExecutor(
		xxl.ServerAddr(cof.Admin.Addresses),   // xxl-job 服务器地址
		xxl.RegistryKey(cof.Executor.AppName), // 执行器名称
		xxl.ExecutorPort(cof.Executor.Port),   // 默认端口，9999（可选）
		xxl.SetLogger(&logger{}),              // 使用自定义日志
	)

	// 初始化执行器
	exec.Init()
	exec.Use(customMiddleware)

	// 设置日志查看处理器
	exec.LogHandler(customLogHandle)

	// 注册任务处理器
	loadTask01()

	// 启动执行器
	go func() {
		if err := exec.Run(); err != nil {
			logrus.Fatalf("XXL-Job 执行器启动失败: %v", err)
		} else {
			logrus.Info("connect xxljob success")
		}
	}()
}

// 自定义日志处理器
func customLogHandle(req *xxl.LogReq) *xxl.LogRes {
	return &xxl.LogRes{Code: xxl.SuccessCode, Msg: "", Content: xxl.LogResContent{
		FromLineNum: req.FromLineNum,
		ToLineNum:   2,
		LogContent:  "这个是自定义日志handler",
		IsEnd:       true,
	}}
}

// 自定义中间件
func customMiddleware(tf xxl.TaskFunc) xxl.TaskFunc {
	return func(cxt context.Context, param *xxl.RunReq) string {
		logrus.Infof("I am a middleware start")
		// 使用 defer + recover 来捕获可能的 panic，防止程序崩溃
		defer func() {
			if err := recover(); err != nil {
				// 记录错误信息
				logrus.Errorf("任务执行时发生 panic：%v", err)
				// 根据需要，可以设置任务执行失败的返回值
			}
		}()
		res := tf(cxt, param)
		logrus.Infof("I am a middleware end")
		return res
	}
}
