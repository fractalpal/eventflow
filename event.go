package eventflow

import (
	"database/sql/driver"
	"errors"
	"time"
)

type EventType string

type Event interface {
	EventAggregatorID() string
	EventType() EventType
	EventTime() time.Time
	EventData() []byte
	EventMapper() string
}

var _ Event = &BaseEvent{}

type BaseEvent struct {
	AggregatorID string
	Type         EventType
	Time         time.Time
	Data         []byte
	Mapper       string
}

func (e *BaseEvent) EventAggregatorID() string {
	return e.AggregatorID
}

func (e *BaseEvent) EventType() EventType {
	return e.Type
}

func (e *BaseEvent) EventTime() time.Time {
	return e.Time
}

func (e *BaseEvent) EventData() []byte {
	return e.Data
}

func (e *BaseEvent) EventMapper() string {
	return e.Mapper
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
