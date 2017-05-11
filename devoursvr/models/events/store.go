package events

import "errors"

//ErrUserNotFound is returned when the requested user is not found in the store
var ErrUserNotFound = errors.New("event not found")

//Store represents an abstract store for model.User objects.
//This interface is used by the HTTP handlers to insert new users,
//get users, and update users. This interface can be implemented
//for any persistent database you want (e.g., MongoDB, PostgreSQL, etc.)
type Store interface {
}
