package handlers

import "net/http"

//EventsHandler handling of ALL of the events
func (ctx *Context) EventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
	}
}

//SpecificEventsHandler handling of just a single event
func (ctx *Context) SpecificEventsHandler(w http.ResponseWriter, r *http.Request) {

}

//EventAttendanceHandler handling of the attendance of events
func (ctx *Context) EventAttendanceHandler(w http.ResponseWriter, r *http.Request) {

}

//EventRecipesHandler handling of the recipes of a event
func (ctx *Context) EventRecipesHandler(w http.ResponseWriter, r *http.Request) {

}
