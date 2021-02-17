package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer //声明一个全局的连接kafka的生产者client
)

func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	//tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	return
}
func SendTiKafka(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder(data)
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)

}

func main() {

}
