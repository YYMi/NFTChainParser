package middleware

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"nftExchangeAdmi-gin/config"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
)

var rootPath string
var logDir string

// CustomFormatter 是一个自定义的日志格式化器
type CustomFormatter struct{}

// 获取当前文件的路径
func getCurrentFilePath() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return file
}

// 获取当前文件所在的目录
func getCurrentDirectory() string {
	filePath := getCurrentFilePath()
	return filepath.Dir(filePath)
}

// 获取 Goroutine ID
func getGoroutineID() string {
	stack := debug.Stack()
	lines := strings.Split(string(stack), "\n")
	if len(lines) > 1 {
		parts := strings.Fields(lines[0])
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return "unknown"
}

// 自定义日志格式化方法
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 格式化时间戳
	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")

	// 获取文件、行号和函数名
	file := entry.Caller.File
	line := entry.Caller.Line

	// 确定根路径，简化路径显示
	if rootPath == "" {
		currentDir := getCurrentDirectory()
		currentDir = strings.Replace(currentDir, "/", ".", -1)
		split := strings.Split(currentDir, ".")
		rootPath = split[len(split)-2]
	}

	// 将绝对路径转换为相对路径
	split := strings.Split(file, rootPath)
	var baseFilePath string
	if len(split) == 1 {
		baseFilePath = split[0]
	} else {
		baseFilePath = split[1]
	}

	// 去掉第一个 '/'
	if strings.HasPrefix(baseFilePath, "/") {
		baseFilePath = baseFilePath[1:]
	}

	// 获取 Goroutine ID
	goroutineID := getGoroutineID()

	// 构建日志行
	logLine := fmt.Sprintf("[ %s ]-[ Go %s ][ %s ][ %s : %d ] %s\n",
		timestamp, goroutineID, entry.Level.String(), baseFilePath, line, entry.Message)
	return []byte(logLine), nil
}

// 初始化日志系统
func InitLogger(_log config.Logger) {
	// 使用自定义格式化器
	logrus.SetFormatter(&CustomFormatter{})

	// 设置日志级别
	level := strings.ToLower(_log.Level)
	switch level {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	// 设置日志轮转器的配置
	logDir = filepath.Join(_log.Path, _log.FileName)
	logger := &lumberjack.Logger{
		Filename:   filepath.Join(logDir, _log.FileName+".log"),
		MaxSize:    _log.MaxSize,
		MaxBackups: 0,
		MaxAge:     180,
		LocalTime:  true,
	}

	// 设置报告调用者信息
	logrus.SetReportCaller(true)

	// 创建一个 MultiWriter，包含 os.Stdout 和 logger
	mw := io.MultiWriter(os.Stdout, logger)
	logrus.SetOutput(mw)

	// 重定向 os.Stdout 和 os.Stderr
	os.Stdout = newMultiWriter(os.Stdout, logger)
	os.Stderr = newMultiWriter(os.Stderr, logger)
}

// newMultiWriter 创建一个 io.Writer，将输出同时写入原始的 os.Stdout/os.Stderr 和日志文件
func newMultiWriter(origWriter *os.File, logger io.Writer) *os.File {
	reader, writer, err := os.Pipe()
	if err != nil {
		logrus.Fatalf("无法创建管道: %v", err)
	}
	go func() {
		multiWriter := io.MultiWriter(origWriter, logger)
		_, err := io.Copy(multiWriter, reader)
		if err != nil {
			logrus.Errorf("复制输出时发生错误: %v", err)
		}
	}()
	return writer
}
