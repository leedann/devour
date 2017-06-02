package handlers

import (
	"encoding/json"
	"net/http"
	"path"

	"time"

	"github.com/leedann/devour/devoursvr/models/users"
	"github.com/leedann/devour/devoursvr/notification"
	"github.com/leedann/devour/devoursvr/sessions"
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
		user, err := ctx.UserStore.Insert(newuser)
		state := &SessionState{
			BeganAt:    time.Now(),
			ClientAddr: r.RequestURI,
			User:       user,
		}
		_, err = sessions.BeginSession(ctx.SessionKey, ctx.SessionStore, state, w)

		_, err = ctx.UserStore.CreateLikesList(user)
		_, err = ctx.UserStore.CreateGroceryList(user)

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

//UserDietHandler handles the diets of a user
//Getting the users diets, adding diet, removing diet,
func (ctx *Context) UserDietHandler(w http.ResponseWriter, r *http.Request) {
	//gets the current user
	state := &SessionState{}
	s, err := sessions.GetSessionID(r, ctx.SessionKey)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	err = ctx.SessionStore.Get(s, state)
	if err != nil {
		http.Error(w, "Error getting sessionID", http.StatusInternalServerError)
		return
	}
	user := state.User
	switch r.Method {
	case "GET":
		diets, err := ctx.UserStore.GetUserDiet(user)
		if err != nil {
			http.Error(w, "Error fetching diets", http.StatusInternalServerError)
			return
		}
		var dietNames []string
		for _, val := range diets {
			diet, _ := ctx.UserStore.GetDietByID(val.DietTypeID)

			dietNames = append(dietNames, diet.Name)
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(dietNames)
	case "POST": //Adds a diet to the user -- requires a diet name
		decoder := json.NewDecoder(r.Body)
		newDiet := &users.NewDiet{}
		if err := decoder.Decode(newDiet); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		//Add all diets
		for _, val := range newDiet.Diets {
			_, err := ctx.UserStore.AddDiet(user, val)
			if err != nil {
				http.Error(w, "Error adding diet", http.StatusBadRequest)
			}
		}
		//Get all the diets to return
		allDiets, err := ctx.UserStore.GetUserDiet(user)
		if err != nil {
			http.Error(w, "Unable to get diets", http.StatusInternalServerError)
		}
		var dietNames []string
		for _, val := range allDiets {
			diet, _ := ctx.UserStore.GetDietByID(val.DietTypeID)
			dietNames = append(dietNames, diet.Name)
		}
		//Recipe events adding recipe FROM event
		ntfy := &notification.DietEvent{
			EventType: notification.NewDiet,
			Message:   dietNames,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(allDiets)
	}
}

//SpecificDietHandler removing a diet
func (ctx *Context) SpecificDietHandler(w http.ResponseWriter, r *http.Request) {
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
	_, dietName := path.Split(r.URL.String())
	user := state.User
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	if r.Method == "DELETE" {
		err := ctx.UserStore.RemoveDiet(user, dietName)
		if err != nil {
			http.Error(w, "Unable to remove that diet", http.StatusBadRequest)
		}
		nameArr := []string{dietName}
		//Removal of a diet
		ntfy := &notification.DietEvent{
			EventType: notification.RemoveDiet,
			Message:   nameArr,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("Diet has been removed from the user"))
	}
}

//UserAllergyHandler handles the allergies of a user
//Getting the users allergies, adding an allergy, removing an allergy
func (ctx *Context) UserAllergyHandler(w http.ResponseWriter, r *http.Request) {
	//gets the current user
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
	user := state.User
	switch r.Method {
	case "GET":
		allergies, err := ctx.UserStore.GetUserAllergy(user)
		if err != nil {
			http.Error(w, "Error fetching allergies", http.StatusInternalServerError)
			return
		}
		var allergyNames []string
		for _, val := range allergies {
			allergy, _ := ctx.UserStore.GetAllergyByID(val.AllergyTypeID)
			allergyNames = append(allergyNames, allergy.Name)
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(allergyNames)
	case "POST": //Adds a allergy to the user -- requires a allergy name
		decoder := json.NewDecoder(r.Body)
		newAllergy := &users.NewAllergy{}
		if err := decoder.Decode(newAllergy); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		//Add all allegies
		for _, val := range newAllergy.Allergies {
			_, err := ctx.UserStore.AddAllergy(user, val)
			if err != nil {
				http.Error(w, "Error adding allergy", http.StatusBadRequest)
			}
		}
		//Get all the allergy to return
		allAllergies, err := ctx.UserStore.GetUserAllergy(user)
		if err != nil {
			http.Error(w, "Unable to get allergies", http.StatusInternalServerError)
		}
		var allergyNames []string
		for _, val := range allAllergies {
			allergy, _ := ctx.UserStore.GetAllergyByID(val.AllergyTypeID)
			allergyNames = append(allergyNames, allergy.Name)
		}
		//Recipe events adding allergy
		ntfy := &notification.AllergyEvent{
			EventType: notification.NewAllergy,
			Message:   allergyNames,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(allergyNames)
	case "DELETE": //Deletes a allergy, requires an allergy name
		decoder := json.NewDecoder(r.Body)
		specificAllergy := &users.SpecificAllergy{}
		if err := decoder.Decode(specificAllergy); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		err := ctx.UserStore.RemoveAllergy(user, specificAllergy.Allergy)
		if err != nil {
			http.Error(w, "Unable to remove that diet", http.StatusBadRequest)
		}
		nameArr := []string{specificAllergy.Allergy}
		//Removal of a allergy
		ntfy := &notification.AllergyEvent{
			EventType: notification.RemoveAllergy,
			Message:   nameArr,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("Allergy has been removed from the user"))
	}
}

//SpecificAllergyHandler removing allergy
func (ctx *Context) SpecificAllergyHandler(w http.ResponseWriter, r *http.Request) {
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
	_, allergyName := path.Split(r.URL.String())
	user := state.User
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	if r.Method == "DELETE" {
		err := ctx.UserStore.RemoveAllergy(user, allergyName)
		if err != nil {
			http.Error(w, "Unable to remove that allergy", http.StatusBadRequest)
		}
		nameArr := []string{allergyName}
		//Removal of a allergy
		ntfy := &notification.AllergyEvent{
			EventType: notification.RemoveAllergy,
			Message:   nameArr,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("Allergy has been removed from the user"))
	}
}

//UserRecipesHandler handles all of the users bookmarked recipes
//getting bookmarked recipes, adding bookmarked recipes, deleting bookmark recipes
func (ctx *Context) UserRecipesHandler(w http.ResponseWriter, r *http.Request) {
	//gets the current user
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
	user := state.User
	switch r.Method {
	case "GET":
		recipes, err := ctx.UserStore.GetUserBook(user)
		if err != nil {
			http.Error(w, "Error fetching recipes", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(recipes.Recipes)
	}
}

//SpecificFavRecipeHandler handles a particular favorite recipe
func (ctx *Context) SpecificFavRecipeHandler(w http.ResponseWriter, r *http.Request) {
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
	_, recipeName := path.Split(r.URL.String())
	user := state.User
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "POST":
		//Adds the recipe
		err := ctx.UserStore.AddToBook(user, recipeName)
		if err != nil {
			http.Error(w, "Error adding favorite recipe", http.StatusInternalServerError)
		}
		//Recipe events adding recipe to book
		ntfy := &notification.RecipesEvent{
			EventType: notification.NewBook,
			Message:   recipeName,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("Recipe has been added from the book"))
	case "DELETE":
		err := ctx.UserStore.DeleteFromBook(user, recipeName)
		if err != nil {
			http.Error(w, "Unable to remove that recipe", http.StatusBadRequest)
		}
		//Recipe events adding recipe to book
		ntfy := &notification.RecipesEvent{
			EventType: notification.RemoveBook,
			Message:   recipeName,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("Recipe has been removed from the book"))
	}
}

//UserFriendsHandler handles the friends of a user
//Adding friends, getting friends, deleting friends
func (ctx *Context) UserFriendsHandler(w http.ResponseWriter, r *http.Request) {
	//gets the current user
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
	me := state.User
	switch r.Method {
	case "GET":
		allFriends, err := ctx.UserStore.GetUserFriendsList(me)
		if err != nil {
			http.Error(w, "Error fetching friends", http.StatusInternalServerError)
			return
		}
		var friends []*users.User
		for _, val := range allFriends {
			usr, _ := ctx.UserStore.GetByID(val.FriendID)

			friends = append(friends, usr)
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(friends)
	}
}

//SpecificFriendHandler handles the user's interaction with a particular friend
func (ctx *Context) SpecificFriendHandler(w http.ResponseWriter, r *http.Request) {
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
	_, friendEmail := path.Split(r.URL.String())
	friend, err := ctx.UserStore.GetByEmail(friendEmail)
	if err != nil {
		http.Error(w, "Error getting friend", http.StatusInternalServerError)
		return
	}
	user := state.User
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "POST":
		//Adds the friend
		friendsList, err := ctx.UserStore.AddFriend(user, friend)
		if err != nil {
			http.Error(w, "Error adding friend", http.StatusInternalServerError)
		}
		frnd, _ := ctx.UserStore.GetByID(friendsList.FriendID)

		//Adding friend
		ntfy := &notification.UserEvent{
			EventType: notification.AddFriend,
			Message:   frnd,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(frnd)
	case "DELETE":
		err := ctx.UserStore.DeleteFriend(user, friend)
		if err != nil {
			http.Error(w, "Unable to remove that friend", http.StatusBadRequest)
		}
		//Deleting friend
		ntfy := &notification.UserEvent{
			EventType: notification.RemoveFriend,
			Message:   friend,
		}
		ctx.Notifier.Notify(ntfy)
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("friend has been removed from the user"))
	}
}

//UserFavoritesHandler handles the favorite friends of a user
//Adding favorites, getting favorites, deleting favorites
func (ctx *Context) UserFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	//gets the current user
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
	user := state.User
	switch r.Method {
	case "GET":
		allFriends, err := ctx.UserStore.GetUserFavFriends(user)
		if err != nil {
			http.Error(w, "Error fetching friends", http.StatusInternalServerError)
			return
		}
		var friends []*users.User
		for _, val := range allFriends {
			usr, _ := ctx.UserStore.GetByID(val.FriendID)

			friends = append(friends, usr)
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(friends)
	}
}

//SpecificFavFriendHandler handles the user's interaction with a particular favorite friend
func (ctx *Context) SpecificFavFriendHandler(w http.ResponseWriter, r *http.Request) {
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
	_, friendEmail := path.Split(r.URL.String())
	friend, err := ctx.UserStore.GetByEmail(friendEmail)
	if err != nil {
		http.Error(w, "Error getting friend", http.StatusInternalServerError)
		return
	}
	user := state.User
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	if r.Method == "PATCH" {
		//get updates and apply to user
		decoder := json.NewDecoder(r.Body)
		updated := &users.UpdateFavoriteFriend{}
		if err := decoder.Decode(updated); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
		}
		switch updated.Method {
		case "add":
			err := ctx.UserStore.AddFavFriend(user, friend)
			if err != nil {
				http.Error(w, "Error adding favorite friend", http.StatusInternalServerError)
				return
			}
			//Adding favorite friend
			ntfy := &notification.UserEvent{
				EventType: notification.AddFavFriend,
				Message:   friend,
			}
			ctx.Notifier.Notify(ntfy)
			w.Header().Add("Content-Type", contentTypeTextUTF8)
			w.Write([]byte("friend has been added as favorite"))
		case "remove":
			err := ctx.UserStore.RemoveFavFriend(user, friend)
			if err != nil {
				http.Error(w, "Error adding favorite friend", http.StatusInternalServerError)
				return
			}
			//removing favorite friend
			ntfy := &notification.UserEvent{
				EventType: notification.RemoveFavFriend,
				Message:   friend,
			}
			ctx.Notifier.Notify(ntfy)
			w.Header().Add("Content-Type", contentTypeTextUTF8)
			w.Write([]byte("friend removed from favorites"))
		default:
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
	}
}

//UserGroceriesHandler handles the users grocery list
//adding to list, deleting list, getting groceries
func (ctx *Context) UserGroceriesHandler(w http.ResponseWriter, r *http.Request) {
	//gets the current user
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
	user := state.User
	switch r.Method {
	case "GET":
		groceries, err := ctx.UserStore.GetUserGroceries(user)
		if err != nil {
			http.Error(w, "Error fetching groceries", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(groceries.Ingredients)
	case "POST": //Adds a grocery to users list
		decoder := json.NewDecoder(r.Body)
		newGrocery := &users.SpecificGrocery{}
		if err := decoder.Decode(newGrocery); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		//Adds grocery
		err := ctx.UserStore.AddToGrocery(user, newGrocery.Grocery)
		if err != nil {
			http.Error(w, "Error adding grocery", http.StatusInternalServerError)
		}
		//Get all the diets to return
		allGroceries, err := ctx.UserStore.GetUserGroceries(user)
		if err != nil {
			http.Error(w, "Unable to get groceries", http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(allGroceries.Ingredients)
	case "DELETE": //Adds a grocery to users list
		decoder := json.NewDecoder(r.Body)
		newGrocery := &users.SpecificGrocery{}
		if err := decoder.Decode(newGrocery); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		//Delete grocery
		err := ctx.UserStore.DeleteFromGrocery(user, newGrocery.Grocery)
		if err != nil {
			http.Error(w, "Error removing grocery", http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", contentTypeTextUTF8)
		w.Write([]byte("removed grocery from list"))
	}
}
