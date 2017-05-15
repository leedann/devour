package messages

import (
	"database/sql"
	"testing"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"

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
		UserName:     "mrtester",
		FirstName:    "test",
		LastName:     "tester",
	}
	nu2 := &users.NewUser{
		Email:        "best@best.com",
		Password:     "password",
		PasswordConf: "password",
		UserName:     "mrbester",
		FirstName:    "best",
		LastName:     "bester",
	}

	newCh := &NewChannel{
		Name:        "Channel1",
		Description: "ChDesc",
		Private:     false,
	}
	nch2 := &NewChannel{
		Name:        "Channel2",
		Description: "ChDesc2",
		Private:     true,
	}

	//reset the auto increment counter and clears previous test users in the DB

	_, err = psdb.Exec("DELETE FROM users")
	_, err = psdb.Exec("DELETE FROM channels")
	_, err = psdb.Exec("DELETE FROM messages")

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
	//Make a new channel
	ch, err := store.InsertChannel(newCh, user.ID)
	if err != nil {
		t.Errorf("error inserting channel: %v\n", err)
	}
	//Make a new channel
	nch, err := store.InsertChannel(nch2, user.ID)
	if err != nil {
		t.Errorf("error inserting channel: %v\n", err)
	}

	//this use should have two channels in allCh
	allCh, err := store.GetAll(user.ID)
	if err != nil {
		t.Errorf("error getting channels: %v\n", err)
	}
	if len(allCh) < 2 {
		t.Errorf("did not receive all channels expected 2 but got %d", len(allCh))
	}

	//getting all channels that are available to user2 (not a mem)
	allCh, err = store.GetAll(user2.ID)
	if err != nil {
		t.Errorf("error getting channels: %v\n", err)
	}
	if len(allCh) < 1 {
		t.Errorf("did not receive all channels expected 2 but got %d", len(allCh))
	}
	//add user2 to members of Channel1
	err = store.Add(ch, user2)
	if err != nil {
		t.Errorf("error adding user to channel %v\n", err)
	}
	// try insert member again -- shouldnt work
	err = store.Add(ch, user2)
	if err != ErrUserExists {
		t.Errorf("expected error user exists but got %v\n", err)
	}
	_, err = store.GetChannelByID(&ch.ID)
	if err != nil {
		t.Errorf("error getting channel %v", err)
	}
	//removes user from chanel
	err = store.Remove(ch, user2)
	if err != nil {
		t.Errorf("error removing user from channel %v\n", err)
	}
	updates := &ChannelUpdates{
		Name:        "NEW2",
		Description: "NEWDesc2",
	}
	//updating channel
	err = store.UpdateChannel(updates, nch)
	updatedCh, err := store.GetChannelByID(&nch.ID)
	if err != nil {
		t.Errorf("error updating channel: %v", err)
	}
	if updatedCh.Name != updates.Name {
		t.Errorf("error changing name expected %s but got %s", updates.Name, nch.Name)
	}
	if updatedCh.Description != updates.Description {
		t.Errorf("error changing name expected %s but got %s", updates.Description, nch.Description)
	}
	newMessage := &NewMessage{
		ChannelID: updatedCh.ID,
		Body:      "testMess",
	}
	messageUpdate := &MessageUpdate{
		Body: "NEWMESSAGE",
	}

	m, err := store.InsertMessage(newMessage, user.ID)
	if err != nil {
		t.Errorf("Error inserting message to channel %v\n", err)
	}
	if m.ChannelID != updatedCh.ID {
		t.Errorf("Error message channel does not match expected %d but got %d", ch.ID, m.ChannelID)
	}

	if m.Body != newMessage.Body {
		t.Errorf("Error message body does not match expected %s but got %s", newMessage.Body, m.Body)
	}
	_, err = store.GetMessageByID(&m.ID)

	messages, err := store.GetRecent(updatedCh)
	if err != nil {
		t.Errorf("Error retrieving messages: %v\n", err)
	}
	//update message
	err = store.UpdateMessage(messageUpdate, m)
	if err != nil {
		t.Errorf("Error updating message %v\n", err)
	}
	messages, err = store.GetRecent(updatedCh)
	if err != nil {
		t.Errorf("Error retrieving messages: %v\n", err)
	}
	if messages[0].Body != messageUpdate.Body {
		t.Errorf("Error updating message expected %s but got %s", messages[0].Body, messageUpdate.Body)
	}
	newMessage = &NewMessage{
		ChannelID: updatedCh.ID,
		Body:      "testMess2",
	}
	m, _ = store.InsertMessage(newMessage, user.ID)
	newMessage = &NewMessage{
		ChannelID: updatedCh.ID,
		Body:      "MESSAGE3",
	}
	_, err = store.InsertMessage(newMessage, user2.ID)
	if err != nil {
		t.Errorf("Error inserting message %v", err)
	}
	//Deleting message
	err = store.DeleteMessage(m)
	if err != nil {
		t.Errorf("Error deleting message %v", err)
	}

	//delete channels and messages in them
	err = store.DeleteChannel(updatedCh)
	if err != nil {
		t.Errorf("Error deleting channel %v", err)
	}

	_, err = psdb.Exec("ALTER SEQUENCE users_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE channels_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE messages_id_seq RESTART")
	_, err = psdb.Exec("DELETE FROM users")
	_, err = psdb.Exec("DELETE FROM channels")
	_, err = psdb.Exec("DELETE FROM messages")
}
