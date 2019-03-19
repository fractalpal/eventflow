package eventflow

type Bus interface {
	Publisher
	Subscriber
}

type Publisher interface {
	Publish(Event) error
}

type Subscriber interface {
	Subscribe(Aggregator, ...EventType)
}
