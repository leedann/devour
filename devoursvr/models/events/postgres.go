package events

import (
	"database/sql"
	"fmt"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"
)

//PGStore store stucture
type PGStore struct {
	DB *sql.DB
}

//InsertEvent inserts an event into a particular database
func (ps *PGStore) InsertEvent(newEvent *NewEvent, creator *users.UserID, eventType string, moodType string) (*Event, error) {
	mood, err := ps.GetMoodByName(moodType)
	if err != nil {
		return nil, err
	}
	eType, err := ps.GetTypeByName(eventType)
	if err != nil {
		return nil, err
	}
	evt, err := newEvent.ToEvent(eType.ID, mood.ID, creator)
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
	row := tx.QueryRow(sql, evt.TypeID, evt.Name, evt.Description, evt.MoodTypeID, evt.StartTime, evt.EndTime)
	err = row.Scan(&evt.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return evt, err
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

//GetMoodByName gets the mood type and returns the whole type
func (ps *PGStore) GetMoodByName(moodName string) (*MoodType, error) {
	var moodType = &MoodType{}
	err := ps.DB.QueryRow(`SELECT * FROM event_mood_type WHERE Name = $1`, moodName).Scan(&moodType.ID, &moodType.Name, &moodType.Description)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return moodType, nil
}
