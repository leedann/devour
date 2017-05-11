package handlers

import (
	"encoding/json"
	"net/http"

	"time"

	"github.com/info344-s17/challenges-leedann/apiserver/models/users"
	"github.com/info344-s17/challenges-leedann/apiserver/sessions"
)

const (
	charsetUTF8         = "charset=utf-8"
	contentTypeJSON     = "application/json"
	contentTypeJSONUTF8 = contentTypeJSON + "; " + charsetUTF8
	contentTypeTextUTF8 = "text/plain; " + charsetUTF8
)

//UserHandler allows users to sign up or gets all users
func (ctx *Context) UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		newuser := &users.NewUser{}
		if err := decoder.Decode(newuser); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		err := newuser.Validate()
		if err != nil {
			http.Error(w, "User not valid", http.StatusBadRequest)
			return
		}
		usr, _ := ctx.UserStore.GetByEmail(newuser.Email)
		if usr != nil {
			http.Error(w, "Email Already Exists", http.StatusBadRequest)
			return
		}

		usr, _ = ctx.UserStore.GetByUserName(newuser.UserName)
		if usr != nil {
			http.Error(w, "Username Already Exists", http.StatusBadRequest)
			return
		}
		user, err := ctx.UserStore.Insert(newuser)
		state := &SessionState{
			BeganAt:    time.Now(),
			ClientAddr: r.RequestURI,
			User:       user,
		}
		_, err = sessions.BeginSession(ctx.SessionKey, ctx.SessionStore, state, w)
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	case "GET":
		users, err := ctx.UserStore.GetAll()
		if err != nil {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(users)
	}
}

//SessionsHandler allows existing users to sign in
func (ctx *Context) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		creds := &users.Credentials{}
		if err := decoder.Decode(creds); err != nil {
			http.Error(w, "Error in Credentials", http.StatusBadRequest)
			return
		}
		u, err := ctx.UserStore.GetByEmail(creds.Email)
		if err != nil {
			http.Error(w, "Email not found", http.StatusUnauthorized)
			return
		}
		err = u.Authenticate(creds.Password)
		if err != nil {
			http.Error(w, "Error authenticating user", http.StatusUnauthorized)
			return
		}
		state := &SessionState{
			BeganAt:    time.Now(),
			ClientAddr: r.RequestURI,
			User:       u,
		}
		_, err = sessions.BeginSession(ctx.SessionKey, ctx.SessionStore, state, w)
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(u)
	} else {
		http.Error(w, "Error with request", http.StatusBadRequest)
		return
	}
}

//SessionsMineHandler allows authenticated users to sign out
func (ctx *Context) SessionsMineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		s, err := sessions.GetSessionID(r, ctx.SessionKey)
		if err != nil {
			http.Error(w, "Could not find authenticated user", http.StatusInternalServerError)
			return
		}
		err = ctx.SessionStore.Delete(s)
		if err != nil {
			http.Error(w, "Error ending session", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("User has been signed out"))
	} else {
		http.Error(w, "Error with request", http.StatusBadRequest)
		return
	}
}

//UsersMeHandler Get the session state
func (ctx *Context) UsersMeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		state := &SessionState{}
		s, err := sessions.GetSessionID(r, ctx.SessionKey)
		if err != nil {
			http.Error(w, "Could not find user", http.StatusInternalServerError)
			return
		}
		err = ctx.SessionStore.Get(s, state)
		if err != nil {
			http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(state.User)
	case "PATCH":
		// get the current authenticated user
		state := &SessionState{}
		s, err := sessions.GetSessionID(r, ctx.SessionKey)
		if err != nil {
			http.Error(w, "Could not find user", http.StatusInternalServerError)
			return
		}
		err = ctx.SessionStore.Get(s, state)
		if err != nil {
			http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
			return
		}
		//get updates and apply to user
		decoder := json.NewDecoder(r.Body)
		updated := &users.UserUpdates{}
		if err := decoder.Decode(updated); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
		}
		err = ctx.UserStore.Update(updated, state.User)
		if err != nil {
			http.Error(w, "Error updating user", http.StatusInternalServerError)
		}
	}

}
