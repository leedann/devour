package events

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/leedann/devour/devoursvr/models/users"
)

//PGStore store stucture
type PGStore struct {
	DB *sql.DB
}

//InsertEvent inserts an event into a particular database
func (ps *PGStore) InsertEvent(newEvent *NewEvent, creator *users.User) (*Event, error) {
	mood, err := ps.GetMoodByName(newEvent.MoodType)
	if err != nil {
		return nil, err
	}
	eType, err := ps.GetTypeByName(newEvent.EventType)
	if err != nil {
		return nil, err
	}
	evt, err := newEvent.ToEvent(eType.ID, mood.ID)
	if err != nil {
		return nil, err
	}
	if evt == nil {
		return nil, fmt.Errorf(".ToEvent() returned nil")
	}
	tx, err := ps.DB.Begin()
	if err != nil {
		return nil, err
	}
	sql := `INSERT INTO events (EventTypeID, Name, Description, MoodTypeID, StartTime, EndTime) VALUES ($1, $2, $3, $4, $5, $6) returning id`
	row := tx.QueryRow(sql, eType.ID, evt.Name, evt.Description, mood.ID, evt.StartTime, evt.EndTime)
	err = row.Scan(&evt.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var attendance = &Attendance{}
	pendingStat, _ := ps.GetAttendanceStatusByName("Host")
	sql = `INSERT INTO event_attendance (UserID, EventID, StatusID) VALUES ($1, $2, $3) RETURNING id`
	row = tx.QueryRow(sql, creator.ID, evt.ID, pendingStat.ID)
	err = row.Scan(&attendance.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return evt, err
}

//GetEventByID gets the event by the id
func (ps *PGStore) GetEventByID(id EventID) (*Event, error) {
	var evt = &Event{}
	err := ps.DB.QueryRow(`SELECT * FROM events WHERE id = $1`, id).Scan(&evt.ID, &evt.TypeID, &evt.Name, &evt.Description, &evt.MoodTypeID, &evt.StartTime, &evt.EndTime)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return evt, nil
}

//GetTypeByName receives the type name and returns the whole type
func (ps *PGStore) GetTypeByName(eventType string) (*EventType, error) {
	var evtType = &EventType{}
	err := ps.DB.QueryRow(`SELECT * FROM event_type WHERE Name = $1`, eventType).Scan(&evtType.ID, &evtType.Name, &evtType.Description)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return evtType, nil
}

//GetTypeByID gets the event type be the id
func (ps *PGStore) GetTypeByID(id TypeID) (*EventType, error) {
	var evtType = &EventType{}
	err := ps.DB.QueryRow(`SELECT * FROM event_type WHERE id = $1`, id).Scan(&evtType.ID, &evtType.Name, &evtType.Description)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return evtType, nil
}

//GetAttendanceStatusByName gets the attendance status from the name
func (ps *PGStore) GetAttendanceStatusByName(status string) (*AttendanceStatus, error) {
	var atStat = &AttendanceStatus{}
	err := ps.DB.QueryRow(`SELECT * FROM event_attendance_status WHERE AttendanceStatus = $1`, status).Scan(&atStat.ID, &atStat.AttendanceStatus)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return atStat, nil
}

//GetAttendanceStatusByID gets the attendance status by its id
func (ps *PGStore) GetAttendanceStatusByID(id StatusID) (*AttendanceStatus, error) {
	var atStat = &AttendanceStatus{}
	err := ps.DB.QueryRow(`SELECT * FROM event_attendance_status WHERE id = $1`, id).Scan(&atStat.ID, &atStat.AttendanceStatus)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return atStat, nil
}

//GetMoodByName gets the mood type and returns the whole type
func (ps *PGStore) GetMoodByName(moodName string) (*MoodType, error) {
	var moodType = &MoodType{}
	err := ps.DB.QueryRow(`SELECT * FROM event_mood_type WHERE Name = $1`, moodName).Scan(&moodType.ID, &moodType.Name, &moodType.Description)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return moodType, nil
}

//GetMoodByID gets the mood type by its ID
func (ps *PGStore) GetMoodByID(id MoodTypeID) (*MoodType, error) {
	var moodType = &MoodType{}
	err := ps.DB.QueryRow(`SELECT * FROM event_mood_type WHERE id = $1`, id).Scan(&moodType.ID, &moodType.Name, &moodType.Description)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return moodType, nil
}

//InviteUserToEvent invites a user to a particular event
func (ps *PGStore) InviteUserToEvent(user *users.User, event *Event) (*Attendance, error) {
	var attendance = &Attendance{}
	attendance.EventID = event.ID
	attendance.UserID = user.ID
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}

	pendingStat, _ := ps.GetAttendanceStatusByName("Pending")
	attendance.StatusID = pendingStat.ID
	sql := `INSERT INTO event_attendance (UserID, EventID, StatusID) VALUES ($1, $2, $3) RETURNING id`
	//Receives ONE row from the database
	row := tx.QueryRow(sql, user.ID, event.ID, pendingStat.ID)
	//scans the value of ID returned from query INTO the user
	err = row.Scan(&attendance.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	return attendance, nil
}

//GetUserAttendanceStatus returns a particular users attendance status to a particular event
func (ps *PGStore) GetUserAttendanceStatus(user *users.User, event *Event) (*AttendanceStatus, error) {
	var atnStatus = &AttendanceStatus{}
	err := ps.DB.QueryRow(`
	SELECT S.ID, AttendanceStatus 
	FROM event_attendance E
	INNER JOIN event_attendance_status S ON E.StatusID = S.ID
	WHERE UserID = $1
	AND E.EventID = $2`, user.ID, event.ID).Scan(&atnStatus.ID, &atnStatus.AttendanceStatus)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return atnStatus, nil
}

//UpdateAttendanceStatus updates the status of a user to an event
func (ps *PGStore) UpdateAttendanceStatus(user *users.User, event *Event, status string) error {
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return err
	}

	stat, err := ps.GetAttendanceStatusByName(status)
	if err != nil {
		return err
	}

	sql := `UPDATE event_attendance SET StatusID = $1 WHERE userid = $2 AND EventID = $3`
	//executes the sql query
	_, err = tx.Exec(sql, stat.ID, user.ID, event.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//UpdateEventStart changes the starting time of the event (ONLY CREATOR)
func (ps *PGStore) UpdateEventStart(event *Event, newTime string) error {
	if newTime == "" {
		return nil
	}
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return err
	}
	start, err := time.Parse(longForm, newTime)
	if err != nil {
		return err
	}

	sql := `UPDATE events SET StartTime = $1 WHERE ID = $2`
	//executes the sql query
	_, err = tx.Exec(sql, start, event.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//UpdateEventEnd changes the ending time of the event (ONLY CREATOR)
func (ps *PGStore) UpdateEventEnd(event *Event, newTime string) error {
	if newTime == "" {
		return nil
	}
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return err
	}
	end, err := time.Parse(longForm, newTime)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	if end.Before(event.StartTime) {
		return fmt.Errorf("Cannot set a end time before the start time")
	}

	sql := `UPDATE events SET EndTime = $1 WHERE ID = $2`
	//executes the sql query
	_, err = tx.Exec(sql, end, event.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//UpdateEventMood changes the mood of the event (ONLY CREATOR)
func (ps *PGStore) UpdateEventMood(event *Event, mood string) error {
	if mood == "" {
		return nil
	}
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return err
	}

	moodID, err := ps.GetMoodByName(mood)
	if err != nil {
		return err
	}

	sql := `UPDATE events SET MoodTypeID = $1 WHERE ID = $2`
	//executes the sql query
	_, err = tx.Exec(sql, moodID.ID, event.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//UpdateEventType changes the type of the event (ONLY CREATOR)
func (ps *PGStore) UpdateEventType(event *Event, typeName string) error {
	if typeName == "" {
		return nil
	}
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return err
	}

	typeID, err := ps.GetTypeByName(typeName)
	if err != nil {
		return err
	}

	sql := `UPDATE events SET EventTypeID = $1 WHERE ID = $2`
	//executes the sql query
	_, err = tx.Exec(sql, typeID.ID, event.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//UpdateEventName changes the name of the Event (ONLY CREATOR)
func (ps *PGStore) UpdateEventName(event *Event, name string) error {
	if name == "" {
		return nil
	}
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return err
	}
	sql := `UPDATE events SET Name = $1 WHERE ID = $2`
	//executes the sql query
	_, err = tx.Exec(sql, name, event.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//UpdateEventDescription changes the description of the event (ONLY CREATOR)
func (ps *PGStore) UpdateEventDescription(event *Event, desc string) error {
	if desc == "" {
		return nil
	}
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return err
	}
	sql := `UPDATE events SET Description = $1 WHERE ID = $2`
	//executes the sql query
	_, err = tx.Exec(sql, desc, event.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//DeleteEvent deletes the event and all statuses pertaining to an event (ONLY CREATOR)
func (ps *PGStore) DeleteEvent(event *Event) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}

	//deleting attendance status
	//have to delete attendance first because they reference the event
	sql := `DELETE FROM event_attendance USING events WHERE event_attendance.EventID = $1`
	_, err = tx.Exec(sql, event.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	//deleting recipes
	//have to delete recipes first because they reference the event
	sql = `DELETE FROM recipe_suggestions USING events WHERE recipe_suggestions.EventID = $1`
	_, err = tx.Exec(sql, event.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	sql = `DELETE FROM events WHERE ID = $1`
	//executes the sql query
	_, err = tx.Exec(sql, event.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//RejectInvite deletes an invite for a particular user will have to use in conjuction with updating a status to not going
func (ps *PGStore) RejectInvite(event *Event, user *users.User) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}

	sql := `DELETE FROM event_attendance WHERE EventID = $1 AND UserID = $2`
	//executes the sql query
	_, err = tx.Exec(sql, event.ID, user.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//AddRecipeToEvent adds a recipe suggestion to an event
func (ps *PGStore) AddRecipeToEvent(event *Event, user *users.User, recipe string) (*RecipeSuggest, error) {
	var suggestion = &RecipeSuggest{}
	suggestion.EventID = event.ID
	suggestion.UserID = user.ID
	suggestion.Recipe = recipe
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}

	sql := `INSERT INTO recipe_suggestions (EventID, UserID, Recipe) VALUES ($1, $2, $3) RETURNING id`
	//Receives ONE row from the database
	row := tx.QueryRow(sql, event.ID, user.ID, recipe)
	//scans the value of ID returned from query INTO the user
	err = row.Scan(&suggestion.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	return suggestion, nil
}

//RemoveRecipeFromEvent removes a recipe suggestion from an event (ONLY CREATOR) and recipe suggester
func (ps *PGStore) RemoveRecipeFromEvent(event *Event, user *users.User, recipe string) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	sql := `DELETE FROM recipe_suggestions WHERE EventID = $1 AND UserID = $2 AND Recipe = $3`
	//executes the sql query
	_, err = tx.Exec(sql, event.ID, user.ID, recipe)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//GetAllRecipesInEvent gets all the recipes in a particular event
func (ps *PGStore) GetAllRecipesInEvent(event *Event) ([]string, error) {
	var allRecipes []string
	rows, err := ps.DB.Query(`
	SELECT recipe 
	FROM recipe_suggestions
	WHERE EventID = $1`, event.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var recipe = ""

		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&recipe); err != nil {
			return nil, err
		}

		//adds to array
		allRecipes = append(allRecipes, recipe)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allRecipes, nil
}

//GetAllUsersInEvent gets all the users that are attending a particular event -- comes back as a slice of users
func (ps *PGStore) GetAllUsersInEvent(event *Event) ([]*users.User, error) {
	var allUsers []*users.User
	rows, err := ps.DB.Query(`
	SELECT U.ID, U.FirstName, U.Email, U.LastName, U.PhotoURL 
	FROM event_attendance E 
	INNER JOIN Users U ON E.UserID = U.ID 
	WHERE E.EventID = $1`, event.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var usr = &users.User{}

		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&usr.ID, &usr.FirstName, &usr.Email, &usr.LastName, &usr.PhotoURL); err != nil {
			return nil, err
		}

		//adds to array
		allUsers = append(allUsers, usr)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allUsers, nil
}

// NEXT -- >  Get all events user invited to ("Pending")
//GetAllPendingEvents returns all of the pending events of a user (have to be user)
func (ps *PGStore) GetAllPendingEvents(user *users.User) ([]*Event, error) {
	var allEvts []*Event
	rows, err := ps.DB.Query(`
	SELECT E.* FROM events E 
	INNER JOIN event_attendance V ON E.ID = V.EventID 
	INNER JOIN event_attendance_status S ON V.StatusID = S.ID 
	WHERE V.UserID = $1 AND S.AttendanceStatus = 'Pending'`, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var evnt = &Event{}

		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&evnt.ID, &evnt.TypeID, &evnt.Name, &evnt.Description, &evnt.MoodTypeID, &evnt.StartTime, &evnt.EndTime); err != nil {
			return nil, err
		}

		//adds to array
		allEvts = append(allEvts, evnt)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allEvts, nil
}

// NEXT -- >  Get all events that have passed
//GetPastEvents gets all of the users past events
func (ps *PGStore) GetPastEvents(user *users.User) ([]*Event, error) {
	var allEvts []*Event
	rows, err := ps.DB.Query(`
	SELECT C.ID, C.EventTypeID, C.Name, C.Description, C.MoodTypeID, C.StartTime, C.EndTime
	FROM users A 
	INNER JOIN event_attendance B ON A.ID = B.UserID
	INNER JOIN events C ON B.EventID = C.ID
	INNER JOIN event_attendance_status D ON D.ID = B.StatusID
	WHERE EndTime < $1 
	AND (B.UserID = $2 
	AND D.AttendanceStatus = 'Pending' OR D.AttendanceStatus = 'Host')`, time.Now(), user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var evnt = &Event{}
		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&evnt.ID, &evnt.TypeID, &evnt.Name, &evnt.Description, &evnt.MoodTypeID, &evnt.StartTime, &evnt.EndTime); err != nil {
			return nil, err
		}

		//adds to array
		allEvts = append(allEvts, evnt)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allEvts, nil
}

// NEXT -- >  Get all events that are coming updates
//GetUpcomingEvents gets all of the events coming up
func (ps *PGStore) GetUpcomingEvents(user *users.User) ([]*Event, error) {
	var allEvts []*Event
	rows, err := ps.DB.Query(`
	SELECT C.ID, C.EventTypeID, C.Name, C.Description, C.MoodTypeID, C.StartTime, C.EndTime
	FROM users A 
	INNER JOIN event_attendance B ON A.ID = B.UserID
	INNER JOIN events C ON B.EventID = C.ID
	INNER JOIN event_attendance_status D ON D.ID = B.StatusID
	WHERE (EndTime > $1)
	AND (B.UserID = $2 
	AND D.AttendanceStatus = 'Pending' OR D.AttendanceStatus = 'Host')`, time.Now(), user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var evnt = &Event{}
		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&evnt.ID, &evnt.TypeID, &evnt.Name, &evnt.Description, &evnt.MoodTypeID, &evnt.StartTime, &evnt.EndTime); err != nil {
			return nil, err
		}

		//adds to array
		allEvts = append(allEvts, evnt)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allEvts, nil
}

// NEXT -- >  Get all hosted events user attending
//GetAllHostedEvents gets all of the events user is hosting
func (ps *PGStore) GetAllHostedEvents(user *users.User) ([]*Event, error) {
	var allEvts []*Event
	rows, err := ps.DB.Query(`
	SELECT E.* FROM events E 
	INNER JOIN event_attendance V ON E.ID = V.EventID 
	INNER JOIN event_attendance_status S ON V.StatusID = S.ID 
	WHERE V.UserID = $1 AND S.AttendanceStatus = 'Host'`, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var evnt = &Event{}

		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&evnt.ID, &evnt.TypeID, &evnt.Name, &evnt.Description, &evnt.MoodTypeID, &evnt.StartTime, &evnt.EndTime); err != nil {
			return nil, err
		}

		//adds to array
		allEvts = append(allEvts, evnt)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allEvts, nil
}

// NEXT -- >  Get all events user attending
//GetAllUserEvents gets all of the events the user is going to or hosting
func (ps *PGStore) GetAllUserEvents(user *users.User) ([]*Event, error) {
	var allEvts []*Event
	rows, err := ps.DB.Query(`
	SELECT E.* FROM events E 
	INNER JOIN event_attendance V ON E.ID = V.EventID 
	INNER JOIN event_attendance_status S ON V.StatusID = S.ID 
	WHERE V.UserID = $1 
	AND S.AttendanceStatus = 'Attending' OR S.AttendanceStatus = 'Host'`, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var evnt = &Event{}

		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&evnt.ID, &evnt.TypeID, &evnt.Name, &evnt.Description, &evnt.MoodTypeID, &evnt.StartTime, &evnt.EndTime); err != nil {
			return nil, err
		}

		//adds to array
		allEvts = append(allEvts, evnt)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allEvts, nil
}

// NEXT -- >  Get all friends going to an event (event page) (might not need but sure here..)
//GetAllFriendsInEvent gets all the friends that are attending a particular event -- comes back as a slice of users
func (ps *PGStore) GetAllFriendsInEvent(user *users.User, event *Event) ([]*users.User, error) {
	var allUsers []*users.User
	rows, err := ps.DB.Query(`
	SELECT U.ID, U.Email, U.FirstName, U.LastName, U.PhotoURL FROM event_attendance E 
	INNER JOIN Users U ON E.UserID = U.ID 
	INNER JOIN friends_list F ON U.ID = F.UserID 
	WHERE E.EventID = $1 AND F.UserID = $2`, event.ID, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var usr = &users.User{}
		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&usr.ID, &usr.Email, &usr.FirstName, &usr.LastName, &usr.PhotoURL); err != nil {
			return nil, err
		}

		//adds to array
		allUsers = append(allUsers, usr)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allUsers, nil
}
