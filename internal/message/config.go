package message

type KafkaConfig struct {
	Brokers        []string
	TopicCreateDoc string
	TopicParseDoc  string
	ConsumerGroup  string
}
