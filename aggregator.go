package eventflow

type Aggregator interface {
	Apply(Event) error
}
