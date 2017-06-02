package events

import (
	"time"

	"github.com/leedann/devour/devoursvr/models/users"
)

const longForm = "January 2, 2006 at 3:04pm (MST)"

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

//SuggestionID represents the primary key of the suggestion of a recipe for an event
type SuggestionID interface{}

//RecipeSuggest shows who suggested which recipe
type RecipeSuggest struct {
	ID      SuggestionID `json:"id" bson:"_id"`
	EventID EventID      `json:"eventid"`
	UserID  users.UserID `json:"userId"`
	Recipe  string       `json:"recipeName"`
}

//DietAllergies all the diets and allergies in event
type DietAllergies struct {
	Allergies []string `json:"allergies"`
	Diets     []string `json:"diets"`
}

//RecipeAdd is the struct required to add a recipe to event
type RecipeAdd struct {
	Recipe string `json:"recipeName"`
}

//Event represents gathering information
type Event struct {
	ID          EventID    `json:"id" bson:"_id"`
	TypeID      TypeID     `json:"typeID"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	MoodTypeID  MoodTypeID `json:"moodTypeID"`
	StartTime   time.Time  `json:"startAt"`
	EndTime     time.Time  `json:"endAt"`
}

//FmtEvent represents gathering information however readable for the client
type FmtEvent struct {
	ID          EventID   `json:"id" bson:"_id"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	MoodType    string    `json:"moodType"`
	StartTime   time.Time `json:"startAt"`
	EndTime     time.Time `json:"endAt"`
}

//NewEvent represents a new event-- recorder from user input
type NewEvent struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	EventType   string `json:"type"`
	MoodType    string `json:"mood"`
}

//Invitation represents the invitation to or removal from event
type Invitation struct {
	Email string `json:"email"`
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

//UpdateAttendance is the struct to update an event
type UpdateAttendance struct {
	EventID          string `json:"eventid"`
	AttendanceStatus string `json:"attendanceStatus"`
}

//ToEvent takes a new event and adds all the necessary information to make an Event
func (ne *NewEvent) ToEvent(eventType TypeID, eventMood MoodTypeID) (*Event, error) {
	event := &Event{}
	//date time has to be formatted exactly like longform

	nStart, err := time.Parse(longForm, ne.StartTime)
	if err != nil {
		return nil, err
	}
	nEnd, err := time.Parse(longForm, ne.EndTime)
	if err != nil {
		return nil, err
	}
	event.StartTime = nStart
	event.EndTime = nEnd
	event.Name = ne.Name
	event.Description = ne.Description
	event.MoodTypeID = eventMood
	event.TypeID = eventType
	return event, nil
}

// go in and search for event id by name and mood id by name
// insert the event to the database
// returns an event
// using event.EventTypeID and event.MoodTypeID --> sql calls get by ID for each
// construct JSON with those values
// var evnt2={
// id: 2,
// Name: "Dinner at Danny's",
// Description: "Come hang out at my house! No need to bring anything",
// Hosting: false,
// Time: "6:00pm",
// StartTime: d2,
// EndTime: d2,
// EventType: "Business",
// MoodType: "Casual",
// }
