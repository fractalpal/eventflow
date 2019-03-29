package eventflow

type EventType string

// Event contains serialized state of particular event that happened
// it is applied by aggregator to evaluate to current state (after this event)
type Event struct {
	// primary identifier for aggregate, like "SomeAggregate" or narrowed "<userId>:SomeAggregate"
	// should be indexed
	RowID string
	// additional columns for aggregate, should be used for segregation/grouping
	// Example: "entity_name" => "foo", "other_id" => "abcd123"
	// should be indexed because most queries would sort/group on them
	Columns map[string]interface{}
	// type of the event, used by aggregators on subscribe
	Type EventType
	// should be indexed because most queries would sort by timestamp
	Timestamp int64
	// serialized event payload
	Data []byte
	// serializer/mapper used to encode data
	Mapper string
}
