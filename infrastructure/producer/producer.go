package producer

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/saufiroja/blog-microservice/users/config"
)

type IProducer interface {
	Publish(topic string, message []byte) error
}

type Producer struct {
	saramaConf *sarama.Config
	producer   sarama.SyncProducer
	conf       *config.AppConfig
}

func NewProducer(conf *config.AppConfig) IProducer {
	saramaConf := sarama.NewConfig()
	saramaConf.Producer.Return.Successes = true
	saramaConf.Producer.Return.Errors = true

	url := fmt.Sprintf("%s:%s", conf.MessageBroker.Host, conf.MessageBroker.Port)

	producer, err := sarama.NewSyncProducer([]string{url}, saramaConf)
	if err != nil {
		panic(err)
	}

	return &Producer{
		saramaConf: saramaConf,
		producer:   producer,
		conf:       conf,
	}
}

func (p *Producer) Publish(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
