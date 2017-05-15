package messages

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"
	"github.com/lib/pq"
)

//ErrUserExists returns an error that the user is already within the targetted channel
var ErrUserExists = errors.New("user already exists in channel")

//PGStore store stucture
type PGStore struct {
	DB *sql.DB
}

func toMemStruct(mems []uint8) []users.UserID {
	//takes the far array and converts to a user.UID array
	//construct the whole thing (since uid splits the brackets)
	counter := 0
	constructor := ""
	var chmembers []users.UserID
	for i := range mems {
		//represents an end bracket in uint8 format
		if mems[i] == 125 {
			counter++
			chmembers = append(chmembers, constructor)
			constructor = ""
		}
		//represents a comma, not done with members array
		if mems[i] == 44 {
			chmembers = append(chmembers, constructor)
			constructor = ""
		}
		//Only want id number not { }
		if mems[i] != 125 && mems[i] != 123 && mems[i] != 44 {
			constructor += string(mems[i])
		}
	}
	return chmembers
}

//GetChannelByID returns the channel with the specified id
func (ps *PGStore) GetChannelByID(id *ChannelID) (*Channel, error) {
	var channel = &Channel{}
	var mems []uint8
	err := ps.DB.QueryRow(`SELECT * FROM channels WHERE ID = $1`, id).Scan(&channel.ID, &channel.Name, &channel.Description, &channel.CreatedAt, &channel.CreatorID, &mems, &channel.Private)
	channel.Members = toMemStruct(mems)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return channel, nil
}

//GetMessageByID returns the message with the specified id
func (ps *PGStore) GetMessageByID(id *MessageID) (*Message, error) {
	var message = &Message{}
	err := ps.DB.QueryRow(`SELECT * FROM messages WHERE ID = $1`, id).Scan(&message.ID, &message.ChannelID, &message.Body, &message.CreatedAt, &message.CreatorID, &message.EditedAt)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return message, nil
}

//GetAll returns channels that a user is part of and all public channels
func (ps *PGStore) GetAll(id users.UserID) ([]*Channel, error) {
	var channels []*Channel

	//Query the database to return multiple rows
	rows, err := ps.DB.Query(`SELECT id, name, description, createdAt, creatorID, members, private FROM channels`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var ch = &Channel{}
		//scans values into User struct; error returned if scan unsuccessful
		//user id array is of type userID struct so we need to put it a temp uint8 array
		var mems []uint8
		if err := rows.Scan(&ch.ID, &ch.Name, &ch.Description, &ch.CreatedAt, &ch.CreatorID, &mems, &ch.Private); err != nil {
			return nil, err
		}
		//takes the int array and converts to a user.UID array
		chmembers := toMemStruct(mems)
		//adds the new users.UID array struct to the channel
		ch.Members = chmembers
		//adds to array
		channels = append(channels, ch)
	}

	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return channels, nil
}

//InsertChannel inserts a new channel into the store
func (ps *PGStore) InsertChannel(newChannel *NewChannel, creator users.UserID) (*Channel, error) {
	ch, err := newChannel.ToChannel(creator)
	if err != nil {
		return nil, err
	}
	if nil == ch {
		return nil, fmt.Errorf(".ToChannel() returned nil")
	}
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}

	sql := `INSERT INTO channels (name, description, createdAt, creatorID, members, private) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	// Receives one row from the db
	row := tx.QueryRow(sql, ch.Name, ch.Description, ch.CreatedAt, ch.CreatorID, pq.Array(ch.Members), ch.Private)
	//scans the value of ID returned from query INTO the channel
	err = row.Scan(&ch.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	return ch, nil
}

//UpdateChannel applies updates to a channel
func (ps *PGStore) UpdateChannel(updates *ChannelUpdates, ch *Channel) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	sql := `UPDATE channels SET Name = $1, Description = $2 WHERE ID = $3`
	//executes the sql query
	_, err = tx.Exec(sql, updates.Name, updates.Description, ch.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//DeleteChannel deletes a channel and all of its messages
func (ps *PGStore) DeleteChannel(ch *Channel) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}

	//deleting messages
	//have to delete messages first because they reference the channel
	//Could have delete on cascade as well
	sql := `DELETE FROM messages USING channels WHERE messages.ChannelID = $1`

	_, err = tx.Exec(sql, ch.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	sql = `DELETE FROM channels WHERE ID = $1`
	//executes the sql query
	_, err = tx.Exec(sql, ch.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//Add adds a user to a channel
func (ps *PGStore) Add(ch *Channel, user *users.User) error {
	//start transaction to see if user is already in target
	var found = ""
	// var checker int
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	//finds the user within the members array
	query := `SELECT ID FROM channels WHERE $1 = ANY(Members) AND ID = $2`
	err = ps.DB.QueryRow(query, user.ID, ch.ID).Scan(&found)
	// if an error comes back-- we WANT errNoRows in order to add
	if err == sql.ErrNoRows {
		tx.Commit()
		//start transaction
		tx, err = ps.DB.Begin()
		if err != nil {
			return err
		}
		//if the user exists in the array then this will add a duplicate
		//so the previous code prevents duplicating
		query = `UPDATE channels SET members = array_append(members, $1) WHERE ID = $2`
		//executes the sql query
		_, err = tx.Exec(query, user.ID, ch.ID)
		//err if could not exec, rollback transaction
		if err != nil {
			tx.Rollback()
			return err
		}
		//commits-- connection no longer reserved
		tx.Commit()
		return nil
	}
	tx.Rollback()
	if err != nil {
		return err
	}
	return ErrUserExists
}

//Remove removes a user from a channel
func (ps *PGStore) Remove(ch *Channel, user *users.User) error {
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	sql := `UPDATE channels SET Members = array_remove(Members, $1) WHERE ID = $2`
	_, err = tx.Exec(sql, user.ID, ch.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//InsertMessage creates a new message
func (ps *PGStore) InsertMessage(newMessage *NewMessage, user users.UserID) (*Message, error) {
	m, err := newMessage.ToMessage(user)
	if err != nil {
		return nil, err
	}
	if nil == m {
		return nil, fmt.Errorf(".ToMessage() returned nil")
	}

	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}

	sql := `INSERT INTO messages (ChannelID, Body, CreatedAt, CreatorID, EditedAt) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	// Receives one row from the db
	row := tx.QueryRow(sql, m.ChannelID, m.Body, m.CreatedAt, m.CreatorID, m.EditedAt)
	//scans the value of ID returned from query INTO the message
	err = row.Scan(&m.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	return m, nil
}

//GetRecent returns N messages from a channel
func (ps *PGStore) GetRecent(ch *Channel) ([]*Message, error) {
	var messages []*Message

	//Query the database to return multiple rows
	rows, err := ps.DB.Query(`SELECT * FROM messages WHERE channelID = $1 ORDER BY createdAt DESC LIMIT 500`, ch.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var m = &Message{}
		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&m.ID, &m.ChannelID, &m.Body, &m.CreatedAt, &m.CreatorID, &m.EditedAt); err != nil {
			return nil, err
		}
		//adds to array
		messages = append(messages, m)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return messages, nil
}

//UpdateMessage updates an existing message
func (ps *PGStore) UpdateMessage(updates *MessageUpdate, message *Message) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	currentTime := time.Now().Local()
	created := currentTime.Format("01/02/2006 3:04pm (MST)")
	createdTime, err := time.Parse("01/02/2006 3:04pm (MST)", created)

	sql := `UPDATE messages SET Body = $1, EditedAt = $2 WHERE ID = $3`
	//executes the sql query
	_, err = tx.Exec(sql, updates.Body, createdTime, message.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//DeleteMessage deletes an existing message
func (ps *PGStore) DeleteMessage(message *Message) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	sql := `DELETE FROM messages WHERE ID = $1`
	//executes the sql query
	_, err = tx.Exec(sql, message.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}
