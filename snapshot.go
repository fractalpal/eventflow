package eventflow

// Snapshot contains serialized state of aggregates up to timestamp
// it can be used to speed up aggregate initialization
type Snapshot struct {
	// primary identifier for aggregate, like "SomeAggregate" or narrowed "<userId>:SomeAggregate"
	// should be indexed
	RowID string
	// should be indexed because most queries would sort by timestamp
	Timestamp int64
	// serialized state of aggregates
	Data []byte
	// serializer/mapper used to encode data
	Mapper string
}
