package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"nftExchangeAdmi-gin/types"
)

func TaskQueueA(mes *string) {
	logrus.Infof(" 获取到执行器推送的消息 %s", *mes)
}
func TaskQueueB(mes *types.User) {

}
