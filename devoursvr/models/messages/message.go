package messages

import (
	"time"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"
)

//MessageID represents the type used in the Message struct
type MessageID interface{}

//Message is the message struct that appears in a channel
type Message struct {
	ID        MessageID    `json:"id" bson:"_id"`
	ChannelID ChannelID    `json:"channelID"`
	Body      string       `json:"body"`
	CreatedAt time.Time    `json:"createdAt"`
	CreatorID users.UserID `json:"creatorID"`
	EditedAt  time.Time    `json:"editedAt,omitempty"`
}

//NewMessage is the starter fields required for a message
type NewMessage struct {
	ChannelID ChannelID `json:"channelID"`
	Body      string    `json:"body"`
}

//MessageUpdate allows a user to update a pre existing message
type MessageUpdate struct {
	Body string `json:"body"`
}

//ToMessage turns a new message into a message
func (nm *NewMessage) ToMessage(user users.UserID) (*Message, error) {
	m := &Message{}
	m.ChannelID = nm.ChannelID
	m.Body = nm.Body
	m.CreatorID = user
	currentTime := time.Now().Local()
	created := currentTime.Format("01/02/2006 3:04pm (MST)")
	createdTime, err := time.Parse("01/02/2006 3:04pm (MST)", created)
	if err != nil {
		return nil, err
	}
	m.CreatedAt = createdTime

	return m, nil

}
