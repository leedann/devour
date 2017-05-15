package events

import (
	"time"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"
)

const longForm = "January 1, 2006 at 3:04PM (MST)"

// select eventtypedesc from event a inner join eventype b on a.eventtypeID = b.eventtypeID where a.descr = ""

//EventID defines the primary key of the event
type EventID interface{}

//StatusID defines the type for any status ids (host, accept, reject)
type StatusID interface{}

//TypeID represents the primary key of the types of events
type TypeID interface{}

//MoodTypeID represents the primary key of the mood of the event
type MoodTypeID interface{}

//AttendanceID represents the primary key of the attendance of a event
type AttendanceID interface{}

//Event represents gathering information
type Event struct {
	ID          EventID      `json:"id" bson:"_id"`
	TypeID      TypeID       `json:"typeID"`
	CreatorID   users.UserID `json:"userID"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	MoodTypeID  MoodTypeID   `json:"moodTypeID"`
	StartTime   time.Time    `json:"startAt"`
	EndTime     time.Time    `json:"endAt"`
}

//NewEvent represents a new event-- recorder from user input
type NewEvent struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

//EventType represents the type of events -- birthday, potluck, etc
type EventType struct {
	ID          EventID `json:"id" bson:"_id"`
	Name        string  `json:"typeName"`
	Description string  `json:"typeDescription"`
}

//MoodType represents the mood for an event -- casual, formal etc
type MoodType struct {
	ID          MoodTypeID `json:"id" bson:"_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}

//Attendance represents the status of an individual at an event (host, accept, reject)
type Attendance struct {
	ID       AttendanceID `json:"id" bson:"_id"`
	EventID  EventID      `json:"eventID"`
	UserID   users.UserID `json:"userID"`
	StatusID StatusID     `json:"statusID"`
}

//AttendanceStatus represents a type of status (accept, reject, host(s))
type AttendanceStatus struct {
	ID               StatusID `json:"id" bson:"_id"`
	AttendanceStatus string   `json:"attendanceStatus"`
}

//ToEvent takes a new event and adds all the necessary information to make an Event
func (ne *NewEvent) ToEvent(eventType EventID, eventMood MoodTypeID, user users.UserID) (*Event, error) {
	event := &Event{}
	//date time has to be formatted exactly like longform

	start := ne.StartTime.Format(longForm)
	end := ne.EndTime.Format(longForm)
	nStart, err := time.Parse(longForm, start)
	if err != nil {
		return nil, err
	}
	nEnd, err := time.Parse(longForm, end)
	if err != nil {
		return nil, err
	}
	event.StartTime = nStart
	event.EndTime = nEnd
	event.Name = ne.Name
	event.Description = ne.Description
	event.MoodTypeID = eventMood
	event.TypeID = eventType
	event.CreatorID = user
	return event, nil
}
