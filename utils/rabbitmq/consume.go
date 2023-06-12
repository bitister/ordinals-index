package rabbitmq

import (
	"fmt"

	"github.com/astaxie/beego"
)

func TestConsume() {
	InitConsumeRabbitmq()
	Consume()
}

func InitConsumeRabbitmq() *RabbitPool {
	if instancePool == nil {
		oncePool.Do(func() {
			instancePool = NewConsumePool()
			//instanceConsumePool.SetMaxConsumeChannel(100)
			ip := beego.AppConfig.String("rabbitmq::ip")
			port, _ := beego.AppConfig.Int("rabbitmq::port")
			username := beego.AppConfig.String("rabbitmq::username")
			password := beego.AppConfig.String("rabbitmq::password")
			err := instancePool.Connect(ip, port, username, password)
			// err := instanceConsumePool.ConnectVirtualHost("192.168.1.169", 5672, "temptest", "test123456", "/temptest1")
			if err != nil {
				fmt.Println(err)
			}
		})
	}
	return instancePool
}

func Consume() {
	nomrl := &ConsumeReceive{
		ExchangeName: beego.AppConfig.String("rabbitmq::transaction_exchange"), //交换机名称
		ExchangeType: EXCHANGE_TYPE_DIRECT,
		Route:        beego.AppConfig.String("rabbitmq::transaction_router"), //router名称
		QueueName:    beego.AppConfig.String("rabbitmq::transaction_queue"),  //队列名称
		IsTry:        true,                                                   //是否重试
		IsAutoAck:    false,                                                  //自动消息确认
		MaxReTry:     5,                                                      //最大重试次数
		EventFail: func(code int, e error, data []byte) {
			fmt.Printf("error:%s", e)
		},
		EventSuccess: func(data []byte, header map[string]interface{}, retryClient RetryClientInterface) bool { //如果返回true 则无需重试
			_ = retryClient.Ack()
			fmt.Printf("data:%s\n", string(data))
			return true
		},
	}
	instancePool.RegisterConsumeReceive(nomrl)
	err := instancePool.RunConsume()
	if err != nil {
		fmt.Println(err)
	}
}
