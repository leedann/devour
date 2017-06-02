package handlers

import (
	"time"

	"github.com/leedann/devour/devoursvr/models/users"
)

//SessionState represents a structure with the user, client's host address that began the session and the time when the session began
type SessionState struct {
	BeganAt    time.Time
	ClientAddr string
	User       *users.User
}
