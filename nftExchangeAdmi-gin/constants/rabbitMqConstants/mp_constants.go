package rabbitMqConstants

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// ExchangeDirect Direct 类型交换机
// Direct 交换机将消息路由到绑定到交换机的队列中，路由规则基于消息的 routing key 完全匹配。
// 使用场景：需要精准控制消息流向，例如任务分发到特定队列
const (
	ExchangeDirect = "direct_exchange"
)

// ExchangeFanout Fanout 类型交换机
// Fanout 交换机会将消息广播到所有绑定的队列，无需考虑 routing key。
// 使用场景：广播消息到多个队列，例如日志广播或系统通知
const (
	ExchangeFanout = "fanout_exchange"
)

// ExchangeTopic 声明 Topic 类型交换机
// Topic 交换机会根据 routing key 模式匹配规则将消息路由到队列。
// routing key 支持通配符：`*` 匹配一个单词，`#` 匹配零个或多个单词。
// 使用场景：基于主题的消息路由，例如事件总线或日志分类。
const (
	ExchangeTopic = "topic_exchange"
)

// MQ 队列名称
const (
	QueueTaskA = "task_queue_a"
	QueueTaskB = "task_queue_b"
)

// MQ 路由键
const (
	RoutingKeyTaskA = "task_a"
	RoutingKeyTaskB = "task_b"
)

// DeclareExchange 声明所有交换机并绑定队列
func DeclareExchange(channel *amqp.Channel) error {
	// 声明 Direct 类型交换机并绑定队列
	if err := declareAndBindDirectExchange(channel); err != nil {
		return err
	}

	// 声明 Fanout 类型交换机
	if err := declareExchange(channel, ExchangeFanout, "fanout"); err != nil {
		return err
	}
	logrus.Debug("Fanout 类型交换机无需绑定队列，已跳过绑定")

	// 声明 Topic 类型交换机
	if err := declareExchange(channel, ExchangeTopic, "topic"); err != nil {
		return err
	}
	logrus.Debug("Topic 类型交换机不绑定队列，可根据业务需要动态绑定")

	logrus.Info("所有交换机声明并绑定队列成功")
	return nil
}

// declareAndBindDirectExchange 声明 Direct 交换机并绑定队列
func declareAndBindDirectExchange(channel *amqp.Channel) error {
	// 声明 Direct 交换机
	if err := declareExchange(channel, ExchangeDirect, "direct"); err != nil {
		return err
	}

	// 绑定队列到 Direct 交换机
	if err := bindQueue(channel, QueueTaskA, ExchangeDirect, RoutingKeyTaskA); err != nil {
		return err
	}
	if err := bindQueue(channel, QueueTaskB, ExchangeDirect, RoutingKeyTaskB); err != nil {
		return err
	}
	return nil
}

// declareExchange 声明交换机
func declareExchange(channel *amqp.Channel, exchangeName, exchangeType string) error {
	err := channel.ExchangeDeclare(
		exchangeName, // 交换机名称
		exchangeType, // 交换机类型
		true,         // 是否持久化
		false,        // 自动删除
		false,        // 内部使用
		false,        // 是否阻塞
		nil,          // 额外参数
	)
	if err != nil {
		logrus.WithError(err).Errorf("声明 %s 类型交换机失败: %s", exchangeType, exchangeName)
		return fmt.Errorf("声明交换机失败: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"exchange": exchangeName,
		"type":     exchangeType,
	}).Info("交换机声明成功")
	return nil
}

// bindQueue 声明队列并绑定到交换机
func bindQueue(channel *amqp.Channel, queueName, exchangeName, routingKey string) error {
	// 声明队列
	_, err := channel.QueueDeclare(
		queueName, // 队列名称
		true,      // 是否持久化
		false,     // 自动删除
		false,     // 是否独占
		false,     // 是否阻塞
		nil,       // 额外参数
	)
	if err != nil {
		logrus.WithError(err).Errorf("声明队列失败: %s", queueName)
		return fmt.Errorf("声明队列失败: %w", err)
	}

	// 绑定队列到交换机
	err = channel.QueueBind(
		queueName,    // 队列名称
		routingKey,   // 路由键
		exchangeName, // 交换机名称
		false,        // 是否阻塞
		nil,          // 额外参数
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"queue":    queueName,
			"exchange": exchangeName,
			"routing":  routingKey,
		}).Error("队列绑定到交换机失败")
		return fmt.Errorf("队列绑定到交换机失败: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"queue":    queueName,
		"exchange": exchangeName,
		"routing":  routingKey,
	}).Info("队列绑定到交换机成功")
	return nil
}

// 动态绑定队列到交换机示例
// 使用场景：根据业务需求动态绑定队列到 Fanout 或 Topic 类型交换机
func dynamicBindQueue(channel *amqp.Channel, queueName, exchangeName, routingKey string) error {
	logrus.Infof("动态绑定队列: %s 到交换机: %s [RoutingKey: %s]", queueName, exchangeName, routingKey)
	return bindQueue(channel, queueName, exchangeName, routingKey)
}
