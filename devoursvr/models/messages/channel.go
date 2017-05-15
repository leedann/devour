package messages

import (
	"time"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"
)

//ChannelID represents the type used in the Channel struct
type ChannelID interface{}

//Channel represents a channel in the database
type Channel struct {
	ID          ChannelID      `json:"id" bson:"_id"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
	CreatorID   users.UserID   `json:"creatorID"`
	Members     []users.UserID `json:"members"`
	Private     bool           `json:"private"`
}

//NewChannel represents a channel to be created as an official channel (above)
type NewChannel struct {
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Members     []users.UserID `json:"members,omitempty"`
	Private     bool           `json:"private,omitempty"`
}

//ChannelUpdates represtents the information that is allowed to change in channels
type ChannelUpdates struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

//ToChannel takes a new Channel and is converted into a channel and is returned
//requries the current user's ID to be passed in, the first member of a channel should be the creator
func (nc *NewChannel) ToChannel(user users.UserID) (*Channel, error) {
	ch := &Channel{}

	ch.Name = nc.Name
	ch.Description = nc.Description
	nc.Members = append(nc.Members, user)
	ch.Members = nc.Members
	ch.Private = nc.Private
	ch.CreatorID = user
	currentTime := time.Now().Local()
	created := currentTime.Format("01/02/2006 3:04pm (MST)")
	createdTime, err := time.Parse("01/02/2006 3:04pm (MST)", created)
	if err != nil {
		return nil, err
	}
	ch.CreatedAt = createdTime
	return ch, nil
}
