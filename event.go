package eventflow

import (
	"database/sql/driver"
	"errors"
	"time"
)

type EventType string

// TODO do we want this as interface ?
type Event struct {
	ID     string
	Type   EventType
	Time   time.Time
	Data   []byte
	Mapper string
}

type Events []Event

func (e *Events) Append(event Event) {
	*e = append(*e, event)
}

// Value - Implementation of valuer for database/sql
func (e EventType) Value() (driver.Value, error) {
	return string(e), nil
}

// Scan - Implement the database/sql scanner interface
func (e *EventType) Scan(value interface{}) error {
	// if value is nil, false
	if value == nil {
		return errors.New("scan value is nil")
	}
	if bv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := bv.(string); ok {
			// set the value of the pointer yne to YesNoEnum(v)
			*e = EventType(v)
			return nil
		}
	}
	// otherwise, return an error
	return errors.New("failed to scan EventType")
}
