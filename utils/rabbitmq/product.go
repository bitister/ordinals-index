package rabbitmq

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

func TestProduct() {
	InitProductRabbitmq()
	rund()
}

func InitProductRabbitmq() *RabbitPool {
	if instancePool == nil {
		oncePool.Do(func() {
			instancePool = NewProductPool()
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

func rund() {

	// var wg sync.WaitGroup

	//wg.Add(1)
	//go func() {
	//	fmt.Println("aaaaaaaaaaaaaaaaaaaaaa")
	//	defer wg.Done()
	//	runtime.SetMutexProfileFraction(1)  // 开启对锁调用的跟踪
	//	runtime.SetBlockProfileRate(1)      // 开启对阻塞操作的跟踪
	//	err:= http.ListenAndServe("0.0.0.0:8080", nil)
	//	fmt.Println(err)
	//}()

	for i := 0; i < 100000; i++ {
		// wg.Add(1)
		go func(num int) {
			// defer wg.Done()
			data := GetRabbitMqDataFormat(beego.AppConfig.String("rabbitmq::transaction_exchange"), EXCHANGE_TYPE_DIRECT, beego.AppConfig.String("rabbitmq::transaction_queue"), beego.AppConfig.String("rabbitmq::transaction_router"), fmt.Sprintf("这里是数据%d", num))
			err := instancePool.Push(data, 0)
			if err != nil {
				fmt.Println(err)
				panic("===")
			}
			fmt.Println("===", num)
		}(i)
	}

	// wg.Wait()
	time.Sleep(60 * time.Second)
}
