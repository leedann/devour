package events

import (
	"time"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"
)

const longForm = "January 1, 2006 at 3:04PM (MST)"

//EventID defines the type for any event id (type, mood, normal)
type EventID interface{}

//StatusID defines the type for any status ids (host, accept, reject)
type StatusID interface{}

//Event represents gathering information
type Event struct {
	EventID         EventID   `json:"eventID" bson:"_eventID"`
	EventTypeID     EventID   `json:"eventTypeID"`
	EventName       string    `json:"eventName"`
	EventDesc       string    `json:"eventDesc"`
	EventMoodTypeID EventID   `json:"eventMoodTypeID"`
	EventStartTime  time.Time `json:"eventStartTime"`
	EventEndTime    time.Time `json:"eventEndTime"`
}

//NewEvent represents a new event-- recorder from user input
type NewEvent struct {
	EventName      string `json:"eventName"`
	EventDesc      string `json:"eventDesc"`
	EventStartTime string `json:"eventStartTime"`
	EventEndTime   string `json:"eventEndTime"`
}

//EventType represents the type of events -- birthday, potluck, etc
type EventType struct {
	EventTypeID   EventID `json:"eventTypeID"`
	EventTypeName string  `json:"eventTypeName"`
	EventTypeDesc string  `json:"eventTypeDesc"`
}

//EventMoodType represents the mood for an event -- casual, formal etc
type EventMoodType struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	DOB          string `json:"dob"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

//EventAttendance represents the status of an individual at an event (host, accept, reject)
type EventAttendance struct {
	EventAttendanceID EventID      `json:"eventAttendanceID"`
	EventID           EventID      `json:"eventID"`
	UserID            users.UserID `json:"userID"`
	StatusID          StatusID     `json:"statusID"`
}

//EventAttendanceStatus represents a type of status (accept, reject, host(s))
type EventAttendanceStatus struct {
	StatusID         StatusID `json:"statusID"`
	AttendanceStatus string   `json:"attendanceStatus"`
}

//ToEvent takes a new event and adds all the necessary information to make an Event
func (ne *NewEvent) ToEvent(eventType string, eventMood string, start string, end string) (*Event, error) {
	event := &Event{}
	//date time has to be formatted exactly like longform
	sTime, _ := time.Parse(longForm, start)
	eTime, _ := time.Parse(longForm, end)
	event.EventStartTime = sTime
	event.EventEndTime = eTime
	event.EventTypeID = eventType
	event.EventMoodTypeID = eventMood
	eventSetting(event, ne)
	return event, nil
}

func eventSetting(e *Event, ne *NewEvent) {
	e.EventName = ne.EventName
	e.EventDesc = ne.EventDesc
}
