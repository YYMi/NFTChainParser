package config

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

type (
	Gin struct {
		Port   string `yaml:"port"`
		Active string `yaml:"active" default:"prod"`
	}
	Cluster struct {
		Nodes []string `yaml:"nodes"`
	}
	Redis struct {
		Database int     `yaml:"database"`
		Password string  `yaml:"password"`
		Cluster  Cluster `yaml:"cluster"`
	}
	MySql struct {
		Tcp      string `yaml:"tcp"`
		DbName   string `yaml:"dbname"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	}
	DataSource struct {
		Debug   bool  `yaml:"debug"`
		Service MySql `yaml:"service"` // 业务数据库
		BcSys   MySql `yaml:"bcSys"`   // 安全数据库
		Bc      MySql `yaml:"bc"`      // 链数据库
	}
	Logger struct {
		Path     string `yaml:"path"`
		FileName string `yaml:"fileName"`
		Level    string `yaml:"level"`
		MaxSize  int    `yaml:"maxSize"`
	}
	Admin struct {
		Addresses string `yaml:"addresses"`
	}
	Executor struct {
		AppName string `yaml:"appname"`
		Port    string `yaml:"port"`
	}
	Xxljob struct {
		Admin    Admin    `yaml:"admin"`
		Executor Executor `yaml:"executor"`
	}

	Secure struct {
		Ignored struct { //白名单路径
			URLs []string `yaml:"urls"`
		} `yaml:"ignored"`
		Encrypted struct { // 加密名单路径
			Timestamp int64    `yaml:"timestamp"`
			URLs      []string `yaml:"urls"`
		}
	}

	JWT struct {
		TokenHeader string `yaml:"tokenHeader"` // JWT 存储的请求头
		Secret      string `yaml:"secret"`      // JWT 加解密使用的密钥
		Expiration  int64  `yaml:"expiration"`  // JWT 的超期限时间（秒）
		TokenHead   string `yaml:"tokenHead"`   // JWT 负载中拿到开头
	}
	MQ struct {
		Rabbitmq struct {
			Addresses         string `yaml:"addresses"`
			VirtualHost       string `yaml:"virtual-host"`
			Username          string `yaml:"username"`
			Password          string `yaml:"password"`
			PublisherConfirms bool   `yaml:"publisher-confirms"`
		} `yaml:"rabbitmq"`
	}
)
type Config struct {
	Name       string     `yaml:"name"`
	Port       int        `yaml:"port"`
	Redis      Redis      `yaml:"redis"`
	DataSource DataSource `yaml:"datasource"`
	Logger     Logger     `yaml:"logger"`
	Xxljob     Xxljob     `yaml:"xxljob"`
	Gin        Gin        `yaml:"gin"`
	Secure     Secure     `yaml:"secure"`
	JWT        JWT        `yaml:"jwt"`
	MQ         MQ         `yaml:"mq"`
}

// MustLoad 从指定的文件路径加载配置到给定的结构体中。
// 如果遇到任何错误，会记录错误并退出程序。
func MustLoad(path string, v interface{}) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("读取配置文件失败 %s: %v", path, err)
	}
	defer file.Close()

	// 读取文件内容
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("读取配置文件失败 %s: %v", path, err)
	}

	// 将 YAML 数据解码到提供的结构体中
	err = yaml.Unmarshal(data, v)
	if err != nil {
		log.Fatalf("解析 YAML 配置文件失败 %s: %v", path, err)
	}
}
