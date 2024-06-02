package entity

// Event entity table model
type Event struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Origin      string `json:"origin"`
	Amount      int    `json:"amount"`
	Destination string `json:"destination"`
}

// EventRepository is the interface for the Event repository
type EventRepository interface {
	Create(event *Event) (*Event, error)
	DeleteAll() error
}
