package rabbitmq

import (
	"encoding/json"
	"fmt"
	"nftExchangeAdmi-gin/config"
	"nftExchangeAdmi-gin/constants/rabbitMqConstants"
	"reflect"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	Connection *amqp.Connection
	Channel    *amqp.Channel
	mu         sync.Mutex // 用于保护全局变量的并发安全
)

// NewRabbitMQ 创建并初始化 RabbitMQ 客户端
// _mq 参数包含 RabbitMQ 的配置信息（地址、虚拟主机、用户名、密码等）
func NewRabbitMQ(_mq config.MQ) {
	amqpURL := fmt.Sprintf("amqp://%s:%s@%s/%s",
		_mq.Rabbitmq.Username,
		_mq.Rabbitmq.Password,
		_mq.Rabbitmq.Addresses,
		_mq.Rabbitmq.VirtualHost,
	)

	for {
		conn, err := amqp.DialConfig(amqpURL, amqp.Config{
			Heartbeat: 10 * time.Second, // 启用心跳机制
		})
		if err != nil {
			logrus.WithError(err).Error("无法连接到 RabbitMQ，正在重试...")
			time.Sleep(5 * time.Second) // 重试间隔
			continue
		}

		ch, err := conn.Channel()
		if err != nil {
			logrus.WithError(err).Error("无法打开 RabbitMQ 通道，正在重试...")
			_ = conn.Close()
			time.Sleep(5 * time.Second) // 重试间隔
			continue
		}

		// 启用发布确认机制
		if _mq.Rabbitmq.PublisherConfirms {
			if err := ch.Confirm(false); err != nil {
				logrus.WithError(err).Error("启用发布确认机制失败，正在重试...")
				_ = ch.Close()
				_ = conn.Close()
				time.Sleep(5 * time.Second)
				continue
			}
			logrus.Info("RabbitMQ 发布确认机制已启用")
		}

		// 初始化交换机和队列绑定
		if err := rabbitMqConstants.DeclareExchange(ch); err != nil {
			logrus.WithError(err).Error("初始化 RabbitMQ 交换机和队列绑定失败，正在重试...")
			_ = ch.Close()
			_ = conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		// 保存连接和通道到全局变量（需加锁以保证并发安全）
		mu.Lock()
		Connection = conn
		Channel = ch
		mu.Unlock()

		// 添加消费者
		if err := AddMQMessageConsumer(rabbitMqConstants.ExchangeDirect, rabbitMqConstants.QueueTaskA, TaskQueueA); err != nil {
			logrus.WithError(err).Error("添加消费者 TaskQueueA 失败，正在重试...")
			_ = ch.Close()
			_ = conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		if err := AddMQMessageConsumer(rabbitMqConstants.ExchangeDirect, rabbitMqConstants.QueueTaskB, TaskQueueB); err != nil {
			logrus.WithError(err).Error("添加消费者 TaskQueueB 失败，正在重试...")
			_ = ch.Close()
			_ = conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		// 启动监听连接关闭的 goroutine
		go monitorConnection(conn, _mq)

		logrus.Info("RabbitMQ 客户端初始化成功")
		break
	}
}

// monitorConnection 监听 RabbitMQ 连接关闭事件，并自动尝试重连
func monitorConnection(conn *amqp.Connection, _mq config.MQ) {
	closeChan := make(chan *amqp.Error)
	conn.NotifyClose(closeChan)

	err := <-closeChan
	if err != nil {
		logrus.WithError(err).Warn("RabbitMQ 连接断开，正在尝试重连...")
		NewRabbitMQ(_mq)
	}
}

// PublishMessage 推送消息到指定交换机
// exchangeName：交换机名称
// routingKey：路由键，用于将消息路由到特定队列
// message：要推送的消息内容，可以是字符串、结构体、map、数组等
func PublishMessage(exchangeName, routingKey string, message interface{}) error {
	if Channel == nil {
		logrus.Error("RabbitMQ 通道未初始化")
		return fmt.Errorf("RabbitMQ 通道未初始化")
	}

	// 将消息序列化为 JSON
	messageBody, err := serializeMessage(message)
	if err != nil {
		logrus.WithError(err).Error("消息序列化失败")
		return fmt.Errorf("消息序列化失败: %w", err)
	}

	// 调试日志
	logrus.WithFields(logrus.Fields{
		"exchange":   exchangeName,
		"routingKey": routingKey,
		"message":    string(messageBody),
	}).Info("准备推送消息")

	// 设置 mandatory 参数为 true，确保消息被正确路由
	err = Channel.Publish(
		exchangeName, // 交换机名称
		routingKey,   // 路由键
		true,         // mandatory: 确保消息路由到队列
		false,        // immediate: 已弃用
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBody,
		},
	)
	if err != nil {
		logrus.Error("推送错误信息 %s", err.Error())
		logrus.WithError(err).Errorf("推送消息失败: %s -> %s", exchangeName, routingKey)
	}

	logrus.Info("消息推送成功")
	return nil
}

// serializeMessage 将消息序列化为 JSON 字节数组
func serializeMessage(message interface{}) ([]byte, error) {
	if message == nil {
		return nil, fmt.Errorf("消息内容不能为空")
	}

	switch reflect.TypeOf(message).Kind() {
	case reflect.String:
		// 如果是字符串，使用 json.Marshal 包装成 JSON 格式
		return json.Marshal(message) // 会自动添加双引号
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		// 如果是结构体、map、slice 或 array 类型，序列化为 JSON
		return json.Marshal(message)
	default:
		// 其他类型转为字符串，再使用 json.Marshal 格式化为 JSON
		return json.Marshal(fmt.Sprintf("%v", message))
	}
}

// AddMQMessageConsumer 添加队列消费者
// exchangeName：交换机名称
// queueName：队列名称
// handler：消息处理函数，需为 func(参数类型 *T) 的形式，参数类型需与消息 JSON 结构对应
func AddMQMessageConsumer(exchangeName, queueName string, handler interface{}) error {
	if Channel == nil {
		return fmt.Errorf("RabbitMQ 通道未初始化")
	}

	// 启动消费者
	megs, err := Channel.Consume(
		queueName, // 队列名称
		"",        // 消费者名称（空字符串则由 RabbitMQ 自动生成）
		true,      // 自动确认消息
		false,     // 非独占
		false,     // 不禁止同一连接的多个消费者消费同一队列
		false,     // 当没有消息时不阻塞
		nil,       // 额外参数
	)
	if err != nil {
		logrus.WithError(err).Errorf("启动消费者失败: %s", queueName)
		return fmt.Errorf("启动消费者失败: %w", err)
	}

	// 使用 goroutine 异步处理消息
	go func() {
		for d := range megs {
			handlerValue := reflect.ValueOf(handler)
			handlerType := handlerValue.Type()

			// 检查 handler 是否为函数且只有一个入参
			if handlerType.Kind() != reflect.Func || handlerType.NumIn() != 1 {
				logrus.Error("handler 必须是一个具有单一参数的函数")
				continue
			}

			argType := handlerType.In(0)
			// 入参必须为指针类型
			if argType.Kind() != reflect.Ptr {
				logrus.Error("handler 的参数必须是指针类型，例如 *string 或 *YourStruct")
				continue
			}

			// 创建与入参类型对应的实例
			argInstance := reflect.New(argType.Elem()).Interface()

			// 反序列化消息体为 handler 参数类型
			if err := json.Unmarshal(d.Body, &argInstance); err != nil {
				logrus.WithError(err).Errorf("消息反序列化失败: %s", string(d.Body))
				continue
			}

			// 调用 handler 处理消息
			handlerValue.Call([]reflect.Value{reflect.ValueOf(argInstance)})
		}
	}()
	logrus.Infof("交换机 %s 消费者 %s 已经添加", exchangeName, queueName)
	return nil
}
