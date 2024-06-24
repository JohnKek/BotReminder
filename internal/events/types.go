package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	UNKNOWN Type = iota
	MESSAGE
)

type Event struct {
	Type Type
	Text string
}
