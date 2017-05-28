package events

import (
	"errors"

	"github.com/leedann/devour/devoursvr/models/users"
)

//ErrUserNotFound is returned when the requested user is not found in the store
var ErrUserNotFound = errors.New("event not found")

//Store represents an abstract store for model.User objects.
//This interface is used by the HTTP handlers to insert new users,
//get users, and update users. This interface can be implemented
//for any persistent database you want (e.g., MongoDB, PostgreSQL, etc.)
type Store interface {
	//Inserts a new event into the db
	InsertEvent(newEvent *NewEvent, creator *users.User) (*Event, error)

	//Helper to get the event type by the name
	GetTypeByName(eventType string) (*EventType, error)

	//Helper to get the attendance status by the name
	GetAttendanceStatusByName(status string) (*AttendanceStatus, error)

	//Helper to get the mood type by its name
	GetMoodByName(moodName string) (*MoodType, error)

	//Invites a user into an event
	InviteUserToEvent(user *users.User, event *Event) (*Attendance, error)

	//Getting the users attendance status of an event
	GetUserAttendanceStatus(user *users.User, event *Event) (*AttendanceStatus, error)

	//Updating a users attendance status
	UpdateAttendanceStatus(user *users.User, event *Event, status string)

	//Updating an events start time
	UpdateEventStart(event *Event, newTime string)

	//Updating the events end time
	UpdateEventEnd(event *Event, newTime string)

	//Updating the events mood type
	UpdateEventMood(event *Event, mood string)

	//Updating the events type
	UpdateEventType(event *Event, typeName string)

	//Updating the events name
	UpdateEventName(event *Event, name string)

	//Updating the events description
	UpdateEventDescription(event *Event, desc string)

	//Deleting the event (also deletes the attendance and recipes)
	DeleteEvent(event *Event) error

	//Rejects a user's invite to an event
	RejectInvite(event *Event, user *users.User) error

	//Adds a recipe to an event (user has to be the creator or the host)
	AddRecipeToEvent(event *Event, user *users.User, recipe string) (*RecipeSuggest, error)

	//Removes a recipe from an event (user has to be the creator or the host)
	RemoveRecipeFromEvent(event *Event, user *users.User, recipe string) error

	//Gets all of the recipes of a particular event
	GetAllRecipesInEvent(event *Event) ([]string, error)

	//Gets all of the users of a particular event
	GetAllUsersInEvent(event *Event) ([]*users.User, error)

	//Gets all of the pending events of a user
	GetAllPendingEvents(user *users.User) ([]*Event, error)

	//Gets all of the past events a user had attended or hosted
	GetPastEvents(user *users.User) ([]*Event, error)

	//Gets all of the upcoming events a user had attended or hosted
	GetUpcomingEvents(user *users.User) ([]*Event, error)

	//Gets all of the events that the user is hosting
	GetAllHostedEvents(user *users.User) ([]*Event, error)

	//Gets all of the users events that they are going to or hosting
	GetAllUserEvents(user *users.User) ([]*Event, error)

	//Gets all of the users friends in the event
	GetAllFriendsInEvent(user *users.User, event *Event) ([]*users.User, error)
}
