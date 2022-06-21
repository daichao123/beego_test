package utils

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"sync"
)

var (
	err              error
	client           sarama.Client
	producer         sarama.AsyncProducer
	wg               sync.WaitGroup
	kfLock           sync.Mutex
	producerInstance *Producer
)

type Producer struct {
	producer  sarama.AsyncProducer
	showDebug bool
}

func init() {
	_ = Instance()
}

// Instance 单例模式
func Instance() *Producer {
	if producerInstance != nil {
		return producerInstance
	}
	kfLock.Lock()
	defer kfLock.Unlock()
	if producerInstance != nil {
		return producerInstance
	}
	kafkaAddr, _ := beego.AppConfig.String("KAFKA_ADDR")
	showDebug, _ := beego.AppConfig.Bool("SHOW_DEBUG")
	producerInstance = NewProducer([]string{kafkaAddr}, showDebug)
	return producerInstance
}

// NewProducer 实例化生产者对象
func NewProducer(kafkaAddr []string, showDebug bool) *Producer {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Errors = true
	client, err = sarama.NewClient(kafkaAddr, config)
	if err != nil {
		panic(err)
	}
	producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// config.Producer.Return.Errors = true 后一定要监听这个chan，默认大小256 如果满了就阻塞掉
		for err := range producer.Errors() {
			logs.Info("[kafka] [error] err=%s", err)
		}
	}()
	instance := &Producer{
		producer:  producer,
		showDebug: showDebug,
	}
	return instance
}

func (p *Producer) NewMessage(topic string, value interface{}) error {
	var k_value sarama.Encoder
	if p.showDebug {
		logs.Info("[kafka] [info] topic=%s, value = %v", topic, value)
	}
	switch value.(type) {
	case string:
		k_value = sarama.StringEncoder(value.(string))
	case int:
		k_value = sarama.ByteEncoder{byte(value.(int))}
	case byte:
		k_value = sarama.ByteEncoder{value.(byte)}
	default:
		return errors.New("value type error")
	}
	message := &sarama.ProducerMessage{Topic: topic, Value: k_value}
	producer.Input() <- message
	return nil
}

func (p *Producer) ShowDebug() {
	p.showDebug = true
}

func (p *Producer) Close() {
	p.producer.AsyncClose()
}
