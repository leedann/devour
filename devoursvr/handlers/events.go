package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/leedann/devour/devoursvr/models/events"
	"github.com/leedann/devour/devoursvr/models/users"
	"github.com/leedann/devour/devoursvr/notification"
	"github.com/leedann/devour/devoursvr/sessions"
)

//allEvents is a structure to be encoded to the client containing all events, past and upcoming
type allEvents struct {
	AllEvents      []*events.FmtEvent `json:"allEvents"`
	UpcomingEvents []*events.FmtEvent `json:"upcomingEvents"`
	PastEvents     []*events.FmtEvent `json:"pastEvents"`
}

//singleEvent is the struct to be encoded to the client containing users, friends, and recipes
type singleEvent struct {
	Host         *users.User           `json:"host"`
	Event        *events.FmtEvent      `json:"event"`
	AllUsers     []*users.User         `json:"allUsers"`
	Friends      []*users.User         `json:"friends"`
	Recipes      []string              `json:"recipes"`
	Restrictions *events.DietAllergies `json:"restrictions"`
}

//FormatEvent turns an event into better readability
func (ctx *Context) formatEvent(evt *events.Event) (*events.FmtEvent, error) {
	formatted := &events.FmtEvent{}
	formatted.ID = evt.ID
	formatted.StartTime = evt.StartTime
	formatted.EndTime = evt.EndTime
	formatted.Name = evt.Name
	formatted.Description = evt.Description
	mood, err := ctx.EventStore.GetMoodByID(evt.MoodTypeID)
	if err != nil {
		return nil, fmt.Errorf("Error converting mood by id")
	}
	typeName, err := ctx.EventStore.GetTypeByID(evt.TypeID)
	if err != nil {
		return nil, fmt.Errorf("Error converting type by id")
	}
	formatted.MoodType = mood.Name
	formatted.Type = typeName.Name
	return formatted, nil
}

//FormatArr takes an array of events and returns a formatted version of it
func (ctx *Context) formatArr(evts []*events.Event) ([]*events.FmtEvent, error) {
	var formatted []*events.FmtEvent
	for _, val := range evts {
		newEvt, err := ctx.formatEvent(val)
		if err != nil {
			return nil, err
		}
		formatted = append(formatted, newEvt)
	}
	return formatted, nil
}

//not getting emails
//HideUserInfo hides all of the non pertinent user details
func (ctx *Context) hideUserInfo(usr *users.User) *users.User {
	newUser := &users.User{}
	newUser.Email = usr.Email
	newUser.FirstName = usr.FirstName
	newUser.LastName = usr.LastName
	newUser.PhotoURL = usr.PhotoURL
	return newUser
}

//HideAllUserInfo hides all of the non pertinent user details and returns a user array
func (ctx *Context) hideAllUserInfo(usr []*users.User) []*users.User {
	var allUsers []*users.User
	for _, val := range usr {
		newUser := ctx.hideUserInfo(val)
		allUsers = append(allUsers, newUser)
	}
	return allUsers
}

//EventsHandler handling of ALL of the events
func (ctx *Context) EventsHandler(w http.ResponseWriter, r *http.Request) {
	//should get the current user regardless
	//gets the current user
	state := &SessionState{}
	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	err = ctx.SessionStore.Get(s, state)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	user := state.User
	switch r.Method {
	//should get all, past, and upcoming
	case "GET":
		allEvt, err := ctx.EventStore.GetAllUserEvents(user)
		if err != nil {
			http.Error(w, "Error getting all events of user", http.StatusInternalServerError)
			return
		}
		upEvt, err := ctx.EventStore.GetUpcomingEvents(user)
		if err != nil {
			http.Error(w, "Error getting all upcoming events of user", http.StatusInternalServerError)
			return
		}
		pastEvt, err := ctx.EventStore.GetPastEvents(user)
		if err != nil {
			http.Error(w, "Error getting all past events of user", http.StatusInternalServerError)
			return
		}
		fmtAll, err := ctx.formatArr(allEvt)
		fmtUpcoming, err := ctx.formatArr(upEvt)
		fmtPast, err := ctx.formatArr(pastEvt)
		if err != nil {
			http.Error(w, "Error formatting events", http.StatusInternalServerError)
			return
		}
		usrEvents := &allEvents{
			AllEvents:      fmtAll,
			UpcomingEvents: fmtUpcoming,
			PastEvents:     fmtPast,
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(usrEvents)
	//creation of an event
	case "POST":
		decoder := json.NewDecoder(r.Body)
		newEvent := &events.NewEvent{}
		if err := decoder.Decode(newEvent); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		evt, err := ctx.EventStore.InsertEvent(newEvent, user)
		if err != nil {
			http.Error(w, "Error inserting new event", http.StatusInternalServerError)
			return
		}
		//creating a new event
		ntfy := &notification.EvtEvent{
			EventType: notification.NewEv,
			Message:   evt,
		}
		//send event to websocket
		ctx.Notifier.Notify(ntfy)

		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(evt)
	}
}

//SpecificEventsHandler handling of just a single event
func (ctx *Context) SpecificEventsHandler(w http.ResponseWriter, r *http.Request) {
	//should get the current user regardless
	//gets the current user
	state := &SessionState{}
	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	err = ctx.SessionStore.Get(s, state)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	//get the url from the path
	_, eventID := path.Split(r.URL.String())
	var eid events.EventID = eventID
	event, err := ctx.EventStore.GetEventByID(eid)
	if err != nil {
		http.Error(w, "Error getting the event", http.StatusBadRequest)
		return
	}
	user := state.User
	hostConfirm, err := ctx.EventStore.GetUserAttendanceStatus(user, event)
	if err != nil {
		http.Error(w, "Could not find a valid invite for that user", http.StatusBadRequest)
		return
	}
	//If the user is not invited to event they cannot see info, or change info regardless
	switch r.Method {
	//Gets all friends going, people in event, recipes
	case "GET":
		allPeople, err := ctx.EventStore.GetAllUsersInEvent(event)
		if err != nil {
			http.Error(w, "Error getting all users in event", http.StatusInternalServerError)
			return
		}
		friends, err := ctx.EventStore.GetAllFriendsInEvent(user, event)
		if err != nil {
			http.Error(w, "Error getting all friends in event", http.StatusInternalServerError)
			return
		}
		recipes, err := ctx.EventStore.GetAllRecipesInEvent(event)
		if err != nil {
			http.Error(w, "Error getting all recipes in event", http.StatusInternalServerError)
			return
		}
		ad := &events.DietAllergies{}
		var allergyNames []string
		var dietNames []string
		for _, val := range allPeople {
			allergies, err := ctx.UserStore.GetUserAllergy(val)
			if err != nil {
				http.Error(w, "Error fetching allergies", http.StatusInternalServerError)
				return
			}
			for _, lrg := range allergies {
				allergy, _ := ctx.UserStore.GetAllergyByID(lrg.AllergyTypeID)
				allergyNames = append(allergyNames, allergy.Name)
			}
			diets, err := ctx.UserStore.GetUserDiet(val)
			if err != nil {
				http.Error(w, "Error fetching diets", http.StatusInternalServerError)
				return
			}
			for _, dname := range diets {
				diet, _ := ctx.UserStore.GetDietByID(dname.DietTypeID)
				dietNames = append(dietNames, diet.Name)
			}
		}
		ad.Allergies = allergyNames
		ad.Diets = dietNames

		fmtPeople := ctx.hideAllUserInfo(allPeople)
		fmtFriends := ctx.hideAllUserInfo(friends)
		fmtEvent, _ := ctx.formatEvent(event)
		host, err := ctx.EventStore.GetHost(event)
		if err != nil {
			http.Error(w, "Error getting host of event", http.StatusInternalServerError)
			return
		}
		eventInfo := &singleEvent{
			Host:         host,
			Event:        fmtEvent,
			AllUsers:     fmtPeople,
			Friends:      fmtFriends,
			Recipes:      recipes,
			Restrictions: ad,
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(eventInfo)
	//Updates Name, Description, Type, Mood, StartTime, EndTime (end cant be before start, user has to be owner of evt)
	case "PATCH":
		if hostConfirm.AttendanceStatus != "Host" {
			http.Error(w, "Cannot make changes to an event you are not hosting", http.StatusBadRequest)
			return
		}
		decoder := json.NewDecoder(r.Body)
		//reusing this struct
		newEvent := &events.NewEvent{}
		if err := decoder.Decode(newEvent); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if err := ctx.EventStore.UpdateEventName(event, newEvent.Name); err != nil {
			http.Error(w, "Error updating name", http.StatusBadRequest)
			return
		}
		if err := ctx.EventStore.UpdateEventDescription(event, newEvent.Description); err != nil {
			http.Error(w, "Error updating description", http.StatusBadRequest)
			return
		}
		if err := ctx.EventStore.UpdateEventMood(event, newEvent.MoodType); err != nil {
			http.Error(w, "Error updating mood", http.StatusBadRequest)
			return
		}
		if err := ctx.EventStore.UpdateEventType(event, newEvent.EventType); err != nil {
			http.Error(w, "Error updating type", http.StatusBadRequest)
			return
		}
		if err := ctx.EventStore.UpdateEventStart(event, newEvent.StartTime); err != nil {
			http.Error(w, "Error updating start time", http.StatusBadRequest)
			return
		}
		if err := ctx.EventStore.UpdateEventEnd(event, newEvent.EndTime); err != nil {
			http.Error(w, "Error updating end time", http.StatusBadRequest)
			return
		}

		updatedEvent, err := ctx.EventStore.GetEventByID(eid)
		if err != nil {
			http.Error(w, "Error getting updating event", http.StatusInternalServerError)
		}
		fmtedEvent, err := ctx.formatEvent(updatedEvent)
		if err != nil {
			http.Error(w, "Error getting formatting event", http.StatusInternalServerError)
		}
		//creating a new event
		ntfy := &notification.FmtEvtEvent{
			EventType: notification.NewEv,
			Message:   fmtedEvent,
		}
		//send event to websocket
		ctx.Notifier.Notify(ntfy)

		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(fmtedEvent)
	//Invites the user to the event -- only owner
	case "LINK":
		if hostConfirm.AttendanceStatus != "Host" {
			http.Error(w, "Cannot make changes to an event you are not hosting", http.StatusBadRequest)
			return
		}
		friendEmail := r.Header.Get("Link")
		friend, err := ctx.UserStore.GetByEmail(friendEmail)
		if err != nil {
			http.Error(w, "Error finding that user", http.StatusBadRequest)
			return
		}
		_, err = ctx.EventStore.InviteUserToEvent(friend, event)
		if err != nil {
			http.Error(w, "Unable to add user to event", http.StatusInternalServerError)
			return
		}

		//add user to list returns a user
		ntfy := &notification.UserEvent{
			EventType: notification.InviteEv,
			Message:   friend,
		}
		//send event to websocket
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("User has been invited to the event"))
	//Removes the user from an event -- only if user has gotten the invite and the user removing friend is host
	case "UNLINK":
		if hostConfirm.AttendanceStatus != "Host" {
			http.Error(w, "Cannot make changes to an event you are not hosting", http.StatusBadRequest)
			return
		}
		friendEmail := r.Header.Get("Link")
		friend, err := ctx.UserStore.GetByEmail(friendEmail)
		if err != nil {
			http.Error(w, "Error finding that user", http.StatusBadRequest)
			return
		}
		_, err = ctx.EventStore.GetUserAttendanceStatus(friend, event)
		//checks to see if the user was invited
		if err != nil {
			http.Error(w, "User is not invited to this event", http.StatusBadRequest)
			return
		}
		//friend is already invited and user is the host of an event
		err = ctx.EventStore.RejectInvite(event, friend)
		if err != nil {
			http.Error(w, "Error removing friend from event", http.StatusInternalServerError)
			return
		}
		//Removing user from event aka rejection
		ntfy := &notification.UserEvent{
			EventType: notification.RejectEv,
			Message:   friend,
		}
		//send event to websocket
		ctx.Notifier.Notify(ntfy)

		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("User has been removed from the event"))
	//Deletes the event -- only owner
	case "DELETE":
		if hostConfirm.AttendanceStatus != "Host" {
			http.Error(w, "Cannot make changes to an event you are not hosting", http.StatusBadRequest)
			return
		}
		if err := ctx.EventStore.DeleteEvent(event); err != nil {
			http.Error(w, "Error deleting the event", http.StatusInternalServerError)
			return
		}
		//deleting Event
		ntfy := &notification.EvtEvent{
			EventType: notification.DeleteEv,
			Message:   event,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("Event successfully deleted"))
	}
}

//EventAttendanceHandler handling of the attendance of events
func (ctx *Context) EventAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	//gets the current user
	state := &SessionState{}
	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	err = ctx.SessionStore.Get(s, state)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	user := state.User
	switch r.Method {
	//Gets all of the user's pending events
	case "GET":
		events, err := ctx.EventStore.GetAllPendingEvents(user)
		if err != nil {
			http.Error(w, "Error getting user's pending events", http.StatusInternalServerError)
			return
		}
		fmted, err := ctx.formatArr(events)
		if err != nil {
			http.Error(w, "Error formatting events", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(fmted)
	//Attending, Not Attending (if not attending reject invite) -- it is already pending
	case "PATCH":
		decoder := json.NewDecoder(r.Body)
		//reusing this struct
		attendanceUpdate := &events.UpdateAttendance{}
		if err := decoder.Decode(attendanceUpdate); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		var eid events.EventID = attendanceUpdate.EventID
		evt, err := ctx.EventStore.GetEventByID(eid)
		if err != nil {
			http.Error(w, "Error getting event", http.StatusBadRequest)
			return
		}
		_, err = ctx.EventStore.GetUserAttendanceStatus(user, evt)
		//err if the user comes back with no rows or an error
		if err != nil {
			http.Error(w, "User is not invited to this event", http.StatusBadRequest)
			return
		}
		//check to see if user is host
		status, err := ctx.EventStore.GetUserAttendanceStatus(user, evt)
		if err != nil {
			http.Error(w, "Could not find attendance status of user", http.StatusInternalServerError)
			return
		}
		//invalid, user cannot change status if they are the host
		if status.AttendanceStatus == "Host" {
			http.Error(w, "Invalid: User cannot change status if hosting", http.StatusBadRequest)
			return
		}
		fmtEvt, err := ctx.formatEvent(evt)
		if err != nil {
			http.Error(w, "Error formatting events", http.StatusInternalServerError)
			return
		}
		//user is invited to the event, check if the update is rejection
		if attendanceUpdate.AttendanceStatus == "Not Attending" {
			err = ctx.EventStore.RejectInvite(evt, user)
			if err != nil {
				http.Error(w, "Error updating attendance status", http.StatusInternalServerError)
				return
			}
			//Changing the attendance of event -- will have to update pending events
			ntfy := &notification.FmtEvtEvent{
				EventType: notification.RejectEv,
				Message:   fmtEvt,
			}
			ctx.Notifier.Notify(ntfy)
		} else {
			err = ctx.EventStore.UpdateAttendanceStatus(user, evt, attendanceUpdate.AttendanceStatus)
			if err != nil {
				http.Error(w, "Error updating attendance status", http.StatusInternalServerError)
				return
			}
			//Changing the attendance of event -- will have to update pending events
			ntfy := &notification.FmtEvtEvent{
				EventType: notification.UpdateAttendance,
				Message:   fmtEvt,
			}
			ctx.Notifier.Notify(ntfy)
		}
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("Attendance successfully updated"))
	}

}

//EventRecipesHandler handling of the recipes of a event
func (ctx *Context) EventRecipesHandler(w http.ResponseWriter, r *http.Request) {
	//gets the current user
	state := &SessionState{}
	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	err = ctx.SessionStore.Get(s, state)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	user := state.User
	_, eventID := path.Split(r.URL.String())
	var eid events.EventID = eventID
	event, err := ctx.EventStore.GetEventByID(eid)
	//check if user is invited to the event
	_, err = ctx.EventStore.GetUserAttendanceStatus(user, event)
	if err != nil {
		http.Error(w, "Cannot add recipe to uninvited event", http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	recipeName := &events.RecipeAdd{}
	if err := decoder.Decode(recipeName); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "POST":
		recipe, err := ctx.EventStore.AddRecipeToEvent(event, user, recipeName.Recipe)
		if err != nil {
			http.Error(w, "Error adding recipe to the event", http.StatusInternalServerError)
			return
		}
		//Recipe events adding recipe FROM event
		ntfy := &notification.RecipesEvent{
			EventType: notification.AddRecipe,
			Message:   recipeName.Recipe,
		}
		ctx.Notifier.Notify(ntfy)

		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(recipe)
	case "DELETE":
		err = ctx.EventStore.RemoveRecipeFromEvent(event, user, recipeName.Recipe)
		if err != nil {
			http.Error(w, "Error deleting recipe from event", http.StatusInternalServerError)
			return
		}
		//Recipe events removing recipe FROM event
		ntfy := &notification.RecipesEvent{
			EventType: notification.RemoveRecipe,
			Message:   recipeName.Recipe,
		}
		ctx.Notifier.Notify(ntfy)

		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("Recipe successfully removed"))
	}
}
