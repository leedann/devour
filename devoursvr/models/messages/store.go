package messages

import (
	"errors"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"
)

//ErrMessageNotFound is returned when the requested message is not found in the store
var ErrMessageNotFound = errors.New("message not found")

//ErrChannelNotFound is returned when the requested channel is not found in the store
var ErrChannelNotFound = errors.New("channel not found")

//Store represents an abstract store for model.messages objects.
//This interface is used by the HTTP handlers to insert new channels
type Store interface {
	//GetAll returns channels that a user is part of and all public channels
	GetAll(id users.UserID) ([]*Channel, error)

	//GetRecent returns N messages from a channel
	GetRecent(ch *Channel) ([]*Message, error)

	//InsertChannel inserts a new channel into the store
	InsertChannel(newChannel *NewChannel, id users.UserID) (*Channel, error)

	//UpdateChannel applies updates to a channel
	UpdateChannel(updates *ChannelUpdates, ch *Channel) error

	//Gets a channel by its ID
	GetChannelByID(id *ChannelID) (*Channel, error)

	//Gets a Message by its ID
	GetMessageByID(id *MessageID) (*Message, error)

	//DeleteChannel deletes a channel and all of its messages
	DeleteChannel(ch *Channel) error

	//Add, adds a user to a channel
	Add(ch *Channel, user *users.User) error

	//Remove, removes a user from a channel
	Remove(ch *Channel, user *users.User) error

	//InsertMessage creates a new message
	InsertMessage(newMessage *NewMessage, id users.UserID) (*Message, error)

	//UpdateMessages updates an existing message
	UpdateMessage(updates *MessageUpdate, message *Message) error

	//DeleteMessage deletes an existing message
	DeleteMessage(message *Message) error
}
