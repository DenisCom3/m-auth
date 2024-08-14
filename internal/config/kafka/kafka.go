package kafka

import "github.com/IBM/sarama"

type Kafka struct {
	Consumer Consumer `yaml:"consumer"`
}

type Consumer struct {
	GroupId    string   `yaml:"group-id"`
	BrokerList []string `yaml:"brokers"`
}

func (c Consumer) Brokers() []string {
	return c.BrokerList
}

func (c Consumer) GroupID() string {
	return c.GroupId
}

func (c Consumer) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}
