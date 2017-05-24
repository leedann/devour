package events

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/leedann/devour/devoursvr/models/users"

	_ "github.com/lib/pq"
)

type AllStores struct {
	MessageStore Store
	UserStore    users.Store
}

//TestPostgresStore tests the dockerized PGStore
func TestPostgresStore(t *testing.T) {
	//Preparing a Postgres data abstraction for later use
	psdb, err := sql.Open("postgres", "user=pgstest dbname=pg2 sslmode=disable")
	if err != nil {
		t.Errorf("error starting db: %v", err)
	}
	//Creates the store structure
	store := &PGStore{
		DB: psdb,
	}

	usrStore := &users.PGStore{
		DB: psdb,
	}
	//Pings the DB-- establishes a connection to the db
	err = psdb.Ping()
	if err != nil {
		t.Errorf("error pinging db %v", err)
	}

	newUser := &users.NewUser{
		Email:        "test@test.com",
		Password:     "password",
		PasswordConf: "password",
		DOB:          "12/12/1990",
		FirstName:    "test",
		LastName:     "tester",
	}
	nu2 := &users.NewUser{
		Email:        "best@best.com",
		Password:     "password",
		PasswordConf: "password",
		DOB:          "12/20/2000",
		FirstName:    "best",
		LastName:     "bester",
	}

	//reset the auto increment counter and clears previous test users in the DB
	_, err = psdb.Exec("ALTER SEQUENCE users_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE user_diet_type_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE user_allergy_type_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE grocery_list_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE user_like_list_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE friends_list_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE event_attendance_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE events_id_seq RESTART")
	_, err = psdb.Exec("DELETE FROM users")
	_, err = psdb.Exec("DELETE FROM user_diet_type")
	_, err = psdb.Exec("DELETE FROM user_allergy_type")
	_, err = psdb.Exec("DELETE FROM grocery_list")
	_, err = psdb.Exec("DELETE FROM user_like_list")
	_, err = psdb.Exec("DELETE FROM friends_list")
	_, err = psdb.Exec("DELETE FROM event_attendance")
	_, err = psdb.Exec("DELETE FROM events")

	//start of insert
	user, err := usrStore.Insert(newUser)
	if err != nil {
		t.Errorf("error inserting user: %v\n", err)
	}
	//means that ToUser() probably was not implemented correctly
	if nil == user {
		t.Fatalf("Nil returned from store.Insert()\n")
	}
	//start of insert
	user2, err := usrStore.Insert(nu2)
	if err != nil {
		t.Errorf("error inserting user: %v\n", err)
	}
	//means that ToUser() probably was not implemented correctly
	if nil == user2 {
		t.Fatalf("Nil returned from store.Insert()\n")
	}

	newEvt := &NewEvent{
		Name:        "testEVENT",
		Description: "testDescription",
		StartTime:   "March 5, 2017 at 4:00pm (PST)",
		EndTime:     "March 5, 2017 at 7:00pm (PST)",
		EventType:   "Formal",
		MoodType:    "Fancy",
	}

	//insert event
	evt, err := store.InsertEvent(newEvt, user)
	if err != nil {
		t.Errorf("error inserting new event %v\n", err)
	}
	if evt.Name != "testEVENT" {
		t.Errorf("error making event expected creator %s but got %s", "testEvent", evt.Name)
	}

	//invite user to the event
	atn, err := store.InviteUserToEvent(user2, evt)
	if err != nil {
		t.Errorf("error inviting user to event %v\n", err)
	}

	//Getting user attendance status
	atnStat, err := store.GetUserAttendanceStatus(user2, evt)
	if err != nil {
		t.Errorf("Error getting user's attendance status")
	}

	if atnStat.AttendanceStatus != "Pending" {
		t.Errorf("Error getting the attendance status expected %s but got %s", "Pending", atnStat.AttendanceStatus)
	}

	if atn.StatusID != atnStat.ID {
		t.Errorf("Error getting the correct attendance status ID expected %d but got %d", atn.StatusID, atnStat.ID)
	}

	//Updating attendance status
	err = store.UpdateAttendanceStatus(user2, evt, "Attending")
	if err != nil {
		t.Errorf("Error getting an updated attendance status")
	}

	//Getting the updated attendance status
	atnStat, err = store.GetUserAttendanceStatus(user2, evt)
	if err != nil {
		t.Errorf("Error getting user's attendance status")
	}
	if atnStat.AttendanceStatus != "Attending" {
		t.Errorf("Error getting the correct UPDATED status: expected Attending but got %s", atnStat.AttendanceStatus)
	}
	fmt.Println(evt)

	//updating event stuff
	err = store.UpdateEventName(evt, "UpdatedTestName")
	if err != nil {
		t.Errorf("Error updating event name %v", err)
	}

	err = store.UpdateEventDescription(evt, "UpdatedDescription")
	if err != nil {
		t.Errorf("Error updating event description %v", err)
	}

	err = store.UpdateEventMood(evt, "Focused")
	if err != nil {
		t.Errorf("Error updating event mood %v", err)
	}

	err = store.UpdateEventType(evt, "Other")
	if err != nil {
		t.Errorf("Error updating event type %v", err)
	}

	err = store.UpdateEventEnd(evt, "March 6, 2017 at 12:00pm (PST)")
	if err != nil {
		t.Errorf("Error updating event end %v", err)
	}

	err = store.UpdateEventStart(evt, "March 1, 2017 at 2:20pm (PST)")
	if err != nil {
		t.Errorf("Error updating event start %v", err)
	}

	upEvents, err := store.GetAllHostedEvents(user)
	if err != nil {
		t.Errorf("Error getting all of the users store: %v", err)
	}
	if upEvents[0].Name != "UpdatedTestName" {
		t.Errorf("Error updating stuffs %v", err)
	}

	err = store.DeleteEvent(evt)
	if err != nil {
		t.Errorf("error deleting event %v", err)
	}

}
