package eventflow

func InMemory() *MemoryFlow {
	return &MemoryFlow{
		aggregators: make(map[EventType]map[Aggregator]struct{}),
	}
}

type MemoryFlow struct {
	aggregators map[EventType]map[Aggregator]struct{}
}

func (b *MemoryFlow) Publish(event Event) error {
	if aggregators, ok := b.aggregators[event.Type]; ok {
		for aggregator := range aggregators {
			if err := aggregator.Apply(event); err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *MemoryFlow) Subscribe(aggregator Aggregator, types ...EventType) {
	// we can try registering multiple aggregators for given type
	// ensure registering just once for given type
	// initialize map if not created already.
	for _, t := range types {
		if _, ok := b.aggregators[t]; !ok {
			b.aggregators[t] = make(map[Aggregator]struct{})
		}

		// add this handler to proper collection
		b.aggregators[t][aggregator] = struct{}{}
	}
}
