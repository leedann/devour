package handlers

import (
	"github.com/leedann/devour/devoursvr/models/events"
	"github.com/leedann/devour/devoursvr/models/users"
	"github.com/leedann/devour/devoursvr/sessions"
)

//Context struct provides context to the session context
type Context struct {
	SessionKey   string
	SessionStore sessions.Store
	UserStore    users.Store
	EventStore   events.Store
}
