package handlers

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/info344-s17/challenges-leedann/apiserver/models/messages"
	"github.com/info344-s17/challenges-leedann/apiserver/sessions"
)

//ChannelsHandler will handle all requests made to the /v1/channels path.
func (ctx *Context) ChannelsHandler(w http.ResponseWriter, r *http.Request) {
	state := &SessionState{}
	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	err = ctx.SessionStore.Get(s, state)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	//state should now contain user if no error
	switch r.Method {
	case "GET":
		channels, err := ctx.MessageStore.GetAll(&state.User.ID)
		if err != nil {
			http.Error(w, "Error retrieving channels", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(channels)
	case "POST":
		newCh := &messages.NewChannel{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(newCh); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		ch, err := ctx.MessageStore.InsertChannel(newCh, &state.User.ID)
		if err != nil {
			http.Error(w, "Error inserting channel", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(ch)
	}
}

//SpecificChannelHandler handles all of the /v1/channels/id
func (ctx *Context) SpecificChannelHandler(w http.ResponseWriter, r *http.Request) {
	state := &SessionState{}
	//grabbing the session ID from the sessionsstore
	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	if err != nil {
		http.Error(w, "Error getting authenticated user", http.StatusInternalServerError)
		return
	}
	//getting the state and putting into state
	err = ctx.SessionStore.Get(s, state)
	if err != nil {
		http.Error(w, "Error getting session state", http.StatusInternalServerError)
		return
	}
	//get the url from the path
	_, channelID := path.Split(r.URL.String())
	//convert the id to a channelID interface
	var chid messages.ChannelID = channelID
	// retrieves the channel based on the path
	ch, err := ctx.MessageStore.GetChannelByID(&chid)
	if err != nil {
		http.Error(w, "Channel does not exist", http.StatusBadRequest)
		return
	}
	u, err := ctx.UserStore.GetByID(state.User.ID)
	if err != nil {
		http.Error(w, "Could not find user", http.StatusInternalServerError)
	}
	chid = ch.CreatorID
	uid := u.ID
	if err != nil {
		http.Error(w, "Error retrieving channel", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "GET":
		if ch.Private == true {
			for _, id := range ch.Members {
				if state.User.ID == id {
					//print messages
					messages, err := ctx.MessageStore.GetRecent(ch)
					if err != nil {
						http.Error(w, "Error retrieving messages", http.StatusInternalServerError)
						return
					}
					w.Header().Add("Content-Type", contentTypeJSONUTF8)
					encoder := json.NewEncoder(w)
					encoder.Encode(messages)
				}
			}
			//user was not found
			http.Error(w, "User is not allowed access in channel", http.StatusUnauthorized)
			return
		}
		//channel is public so can just show messages
		messages, err := ctx.MessageStore.GetRecent(ch)
		if err != nil {
			http.Error(w, "Error retrieving messages", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(messages)
	case "PATCH":
		if uid == chid {
			updates := &messages.ChannelUpdates{}
			decoder := json.NewDecoder(r.Body)
			//decode into updates
			if err := decoder.Decode(updates); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			err := ctx.MessageStore.UpdateChannel(updates, ch)
			if err != nil {
				http.Error(w, "Error updating the channel", http.StatusBadRequest)
				return
			}
			w.Header().Add("Content-Type", contentTypeJSONUTF8)
			channel, err := ctx.MessageStore.GetChannelByID(&chid)
			if err != nil {
				http.Error(w, "Error retrieving new Channel", http.StatusInternalServerError)
			}
			encoder := json.NewEncoder(w)
			encoder.Encode(channel)
		} else {
			http.Error(w, "User is not owner of channel", http.StatusUnauthorized)
			return
		}
	case "DELETE":
		if uid == chid {
			err := ctx.MessageStore.DeleteChannel(ch)
			if err != nil {
				http.Error(w, "error deleting channel", http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", contentTypeTextUTF8)
			w.Write([]byte("Channel has been successfully deleted"))
		} else {
			http.Error(w, "User is not the owner of this channel", http.StatusUnauthorized)
			return
		}
	case "LINK":
		if ch.Private == false {
			//add current user to list
			err := ctx.MessageStore.Add(ch, state.User)
			if err != nil {
				http.Error(w, "Error could not add user to store", http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", contentTypeTextUTF8)
			w.Write([]byte("User has been added to channel"))
		} else {
			if chid == uid {
				otherUID := r.Header.Get("Link")
				//get user from id
				otherUser, err := ctx.UserStore.GetByID(otherUID)
				if err != nil {
					http.Error(w, "Could not find user by that ID", http.StatusBadRequest)
					return
				}
				//add user to channel
				err = ctx.MessageStore.Add(ch, otherUser)
				if err != nil {
					http.Error(w, "Could not add user to channel", http.StatusInternalServerError)
					return
				}
				w.Header().Add("Content-Type", contentTypeTextUTF8)
				w.Write([]byte("User has been added to channel"))
			} else {
				http.Error(w, "User is not owner of channel", http.StatusUnauthorized)
				return
			}
		}
	case "UNLINK":
		if ch.Private == false {
			//remove current user from channel
			err := ctx.MessageStore.Remove(ch, state.User)
			if err != nil {
				http.Error(w, "Error removing user from channel", http.StatusBadRequest)
				return
			}
			w.Header().Add("Content-Type", contentTypeTextUTF8)
			w.Write([]byte("User has been removed from channel"))
		} else {
			otherUID := r.Header.Get("Link")
			//get user from id
			otherUser, err := ctx.UserStore.GetByID(otherUID)
			if err != nil {
				http.Error(w, "Could not find user by that ID", http.StatusBadRequest)
				return
			}
			err = ctx.MessageStore.Remove(ch, otherUser)
			if err != nil {
				http.Error(w, "Error removing user from channel", http.StatusBadRequest)
				return
			}
			w.Header().Add("Content-Type", contentTypeTextUTF8)
			w.Write([]byte("User has been removed from channel"))
		}
	}
}

//MessagesHandler inserts new messages
func (ctx *Context) MessagesHandler(w http.ResponseWriter, r *http.Request) {
	state := &SessionState{}
	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	err = ctx.SessionStore.Get(s, state)
	if err != nil {
		http.Error(w, "Error getting session", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		message := &messages.NewMessage{}
		if err := decoder.Decode(message); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		insertedMsg, err := ctx.MessageStore.InsertMessage(message, &state.User.ID)
		if err != nil {
			http.Error(w, "Error inserting new message", http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(insertedMsg)
	}
}

//SpecificMessageHandler updates messages and deletes
func (ctx *Context) SpecificMessageHandler(w http.ResponseWriter, r *http.Request) {
	state := &SessionState{}
	//grabbing the session ID from the sessionsstore
	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	if err != nil {
		http.Error(w, "Error getting authenticated user", http.StatusInternalServerError)
		return
	}
	//getting the state and putting into state
	err = ctx.SessionStore.Get(s, state)
	if err != nil {
		http.Error(w, "Error getting session state", http.StatusInternalServerError)
		return
	}
	//get the url from the path
	_, messageID := path.Split(r.URL.String())
	var mid messages.MessageID = messageID
	message, err := ctx.MessageStore.GetMessageByID(&mid)
	if err != nil {
		http.Error(w, "Error finding that message", http.StatusBadRequest)
		return
	}
	u, err := ctx.UserStore.GetByID(state.User.ID)
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	mid = message.CreatorID
	uid := u.ID

	if err != nil {
		http.Error(w, "Error retrieving Message", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "PATCH":
		if mid == uid {
			updates := &messages.MessageUpdate{}
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(updates); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			//update message
			err := ctx.MessageStore.UpdateMessage(updates, message)
			if err != nil {
				http.Error(w, "Error updating message", http.StatusInternalServerError)
				return
			}
			//have to retrieve the message again
			updated, err := ctx.MessageStore.GetMessageByID(&message.ID)
			w.Header().Add("Content-Type", contentTypeJSONUTF8)
			encoder := json.NewEncoder(w)
			encoder.Encode(updated)
		} else {
			http.Error(w, "User is not owner of channel", http.StatusUnauthorized)
			return
		}
	case "DELETE":
		if mid == uid {
			err := ctx.MessageStore.DeleteMessage(message)
			if err != nil {
				http.Error(w, "Error deleting message", http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", contentTypeTextUTF8)
			w.Write([]byte("Message has been removed from channel"))
		} else {
			http.Error(w, "User is not owner of channel", http.StatusUnauthorized)
			return
		}
	}
}
