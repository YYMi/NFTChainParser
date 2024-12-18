package xxljob

import (
	"context"
	"github.com/xxl-job/xxl-job-executor-go"
	"nftExchangeAdmi-gin/constants/rabbitMqConstants"
	"nftExchangeAdmi-gin/mq/rabbitmq"
)

func loadTask01() {
	exec.RegTask("transactionNotificationHandler", TransactionNotificationHandler)

}
func TransactionNotificationHandler(ctx context.Context, param *xxl.RunReq) string {
	err := rabbitmq.PublishMessage(rabbitMqConstants.ExchangeDirect, rabbitMqConstants.RoutingKeyTaskA, "测试123")
	if err != nil {
		return "error: invalid param format"
	}
	return "success"
}
