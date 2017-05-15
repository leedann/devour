package messages

import (
	"testing"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"
)

func createNewChannel() *NewChannel {
	return &NewChannel{
		Name:        "ChannelTest",
		Description: "ChannelDesc",
		Members:     []users.UserID{"1"},
		Private:     false,
	}
}

func TestToChannel(t *testing.T) {
	nch := createNewChannel()
	_, err := nch.ToChannel(nch.Members[0])
	if err != nil {
		t.Errorf("Error creating channel %s", err)
	}
}
