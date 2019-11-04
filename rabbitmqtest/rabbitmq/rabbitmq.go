package rabbitmq

import (
	"github.com/streadway/amqp"
)


var (
	BTCKEY = "BTC"
	RabbitmqURL= "localhost:5672"
	RabbitmqUsername = "guest"
	RabbitmqPasswd = "guest"
	BTCExchange = "BTCExchange"
	BTCExitQueue = "BTCExitQueue"
	BTCBalanceQueue = "BTCBalanceQueue"
	BTCCenoddsQueue = "BTCCenoddsQueue"
	BTCExitRK = "BTCExitRK"
	BTCBalanceRK = "BTCBalanceRK"
	BTCCenoddsRK = "BTCCenoddsRK"
	Exchanges = []string{BTCExchange,}
	ExchangeMap = map[string]string{
		BTCKEY:BTCExchange,
	}
	BTCQueues = []string{BTCExitQueue,BTCBalanceQueue,BTCCenoddsQueue,}
	BTCQueueRK = map[string]string{
		BTCExitQueue:BTCExitRK,
		BTCBalanceQueue:BTCBalanceRK,
		BTCCenoddsQueue:BTCCenoddsRK,
	}
	BindMaps = map[string][]string{
		BTCExchange:BTCQueues,
	}

	ExchangeRK = map[string]map[string]string{
		BTCExchange:BTCQueueRK,
	}

	BanlanceRKMap = map[string]string{
		BTCKEY:BTCBalanceRK,
	}

	ExitRKMap = map[string]string{
		BTCKEY:BTCExitRK,
	}
	CenoddsRKMap = map[string]string{
		BTCKEY:BTCCenoddsRK,
	}
)




var (
	webconn *amqp.Connection
	WebExitChannel *amqp.Channel
	WebBalanceChannel *amqp.Channel
	WebCenoddsChannel *amqp.Channel

	BTCconn *amqp.Connection
	BTCExitChannel *amqp.Channel
	BTCBalanceChannel *amqp.Channel
	BTCCenoddsChannel *amqp.Channel

	SychanExitMap = map[string]*amqp.Channel{
		BTCKEY:BTCExitChannel,
	}

	SychanBalanceMap =  map[string]*amqp.Channel{
		BTCKEY:BTCBalanceChannel,
	}

	SychanCenoddsMap = map[string]*amqp.Channel {
		BTCKEY:BTCCenoddsChannel,
	}
)


//所以必须web要先启动 或者可以先手动创建好

//重复创建没关系 exchange和queue 但是毕竟是阻塞读取的 还是创建不同的进程init好了
//web3个channel负责往不同标的物的发送消息 只要带路由key就可以了
func Webinit() {
	WebConnection()
}


func WebConnection() (err error){
	url := "amqp://" + RabbitmqUsername + ":" + RabbitmqPasswd + "@" + RabbitmqURL + "/"

	if webconn == nil {
		var mqConfig = amqp.Config{ChannelMax: 10}
		webconn, err = amqp.DialConfig(url, mqConfig)

		if err != nil {
			return err
		}

		WebExitChannel, err = webconn.Channel()
		WebBalanceChannel, err = webconn.Channel()
		WebCenoddsChannel, err = webconn.Channel()

		if err != nil {
			return err
		}


		for _,e := range Exchanges {
			WebExitChannel.ExchangeDeclare(e,"direct",true,false,false,false,nil)
			WebBalanceChannel.ExchangeDeclare(e,"direct",true,false,false,false,nil)
			WebCenoddsChannel.ExchangeDeclare(e,"direct",true,false,false,false,nil)
		}
	}
	return nil
}

//已经创建交换机 只管创建channel和队列 再绑定exchange就可以了
func BTCinit() (err error) {
	url := "amqp://" + RabbitmqUsername + ":" + RabbitmqPasswd + "@" + RabbitmqURL + "/"

	if BTCconn == nil {
		var mqConfig = amqp.Config{ChannelMax: 10}
		BTCconn, err = amqp.DialConfig(url, mqConfig)

		if err != nil {
			return err
		}

		BTCExitChannel, err = BTCconn.Channel()
		BTCBalanceChannel, err = BTCconn.Channel()
		BTCCenoddsChannel, err = BTCconn.Channel()

		if err != nil {
			return err
		}

		BTCExitChannel.ExchangeDeclare(BTCExchange,"direct",true,false,false,false,nil)
		BTCExitChannel.QueueDeclare(BTCExitQueue,true,false,false,false,nil)
		BTCExitChannel.QueueBind(BTCExitQueue,BTCExitRK,BTCExchange,false,nil)
		BTCExitChannel.Qos(1, 0, true)

		BTCBalanceChannel.ExchangeDeclare(BTCExchange,"direct",true,false,false,false,nil)
		BTCBalanceChannel.QueueDeclare(BTCBalanceQueue,true,false,false,false,nil)
		BTCBalanceChannel.QueueBind(BTCBalanceQueue,BTCBalanceRK,BTCExchange,false,nil)
		BTCBalanceChannel.Qos(1, 0, true)

		BTCCenoddsChannel.ExchangeDeclare(BTCExchange,"direct",true,false,false,false,nil)
		BTCCenoddsChannel.QueueDeclare(BTCCenoddsQueue,true,false,false,false,nil)
		BTCCenoddsChannel.QueueBind(BTCCenoddsQueue,BTCCenoddsRK,BTCExchange,false,nil)
		BTCCenoddsChannel.Qos(1, 0, true)
	}
	return nil
}



