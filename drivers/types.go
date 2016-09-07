package drivers

import "time"

type Database interface {
	Connection() (Connection, error)
}

type Connection interface {
	Close() error
	Collection(name string) (Collection, error)
}

type Collection interface {
	Name() string
	Append(event Event) (Event, error)
	Get(query Query) ([]Event, error)
}

type Query interface {
	ID() string
	Multi() bool
	Type() *string
	After() *time.Time
	Before() *time.Time
}

type Event interface {
	ID() string
	Type() string
	Version() uint64
	When() time.Time
	Payload() interface{}
}
