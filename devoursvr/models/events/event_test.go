package events

import (
	"fmt"
	"testing"
)

func createNewEvent() *NewEvent {
	return &NewEvent{
		Name:        "Event-Name",
		Description: "Event-Desciption",
		StartTime:   "March 5, 2017 at 4:00PM (PST)",
		EndTime:     "March 5, 2017 at 8:00PM (PST)",
		EventType:   "Formal",
		MoodType:    "Silly",
	}
}

func TestNewEventToEvent(t *testing.T) {
	ne := createNewEvent()

	e, err := ne.ToEvent("1", "2", "3")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(e)
}
