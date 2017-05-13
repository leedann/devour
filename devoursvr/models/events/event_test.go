package events

import (
	"fmt"
	"testing"
)

func createNewEvent() *NewEvent {
	return &NewEvent{
		EventName:      "Event-Name",
		EventDesc:      "Event-Desciption",
		EventStartTime: "March 5, 2017 at 4:00PM (PST)",
		EventEndTime:   "March 5, 2017 at 8:00PM (PST)",
	}
}

func TestNewEventToEvent(t *testing.T) {
	ne := createNewEvent()
	e, err := ne.ToEvent("Potluck", "Casual", ne.EventStartTime, ne.EventEndTime)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(e)
}
