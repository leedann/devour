package users

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

//TestPostgresStore tests the dockerized PGStore
func TestPostgresStore(t *testing.T) {
	//Preparing a Postgres data abstraction for later use
	psdb, err := sql.Open("postgres", "user=pgstest dbname=devourpg sslmode=disable")
	if err != nil {
		t.Errorf("error starting db: %v", err)
	}
	//Creates the store structure
	store := &PGStore{
		DB: psdb,
	}
	//Pings the DB-- establishes a connection to the db
	err = psdb.Ping()
	if err != nil {
		t.Errorf("error pinging db %v", err)
	}

	newUser := &NewUser{
		Email:        "test@test.com",
		Password:     "password",
		PasswordConf: "password",
		DOB:          "12/12/1990",
		FirstName:    "test",
		LastName:     "tester",
	}

	//reset the auto increment counter and clears previous test users in the DB
	_, err = psdb.Exec("ALTER SEQUENCE users_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE user_diet_type_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE user_allergy_type_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE grocery_list_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE user_like_list_id_seq RESTART")
	_, err = psdb.Exec("ALTER SEQUENCE friends_list_id_seq RESTART")
	_, err = psdb.Exec("DELETE FROM users")
	_, err = psdb.Exec("DELETE FROM user_diet_type")
	_, err = psdb.Exec("DELETE FROM user_allergy_type")
	_, err = psdb.Exec("DELETE FROM grocery_list")
	_, err = psdb.Exec("DELETE FROM user_like_list")
	_, err = psdb.Exec("DELETE FROM friends_list")

	if err != nil {
		t.Errorf("could not delete table: %v\n", err)
	}
	//start of insert
	user, err := store.Insert(newUser)
	if err != nil {
		t.Errorf("error inserting user: %v\n", err)
	}
	//means that ToUser() probably was not implemented correctly
	if nil == user {
		t.Fatalf("Nil returned from store.Insert()\n")
	}

	//getting user from ID of previous inserted user
	user2, err := store.GetByID(user.ID)
	if err != nil {
		t.Errorf("error finding user by ID: %v\n", err)
	}
	//found something but didnt match
	if user.ID != user2.ID {
		t.Errorf("ID of user retrieved by ID does not match: expected %s but got %s\n", user.ID, user2.ID)
	}
	//getting any user with the given email
	user2, err = store.GetByEmail(newUser.Email)
	if err != nil {
		t.Errorf("error getting user by email: %v\n", err)
	}
	if user.ID != user2.ID {
		t.Errorf("ID of user retreived by Email does not match: expected %s but got %s\n", user.ID, user2.ID)
	}

	update := &UserUpdates{
		FirstName: "UPDATED Test",
		LastName:  "UPDATED Tester",
	}
	//updates the store with fields in update
	if err = store.Update(update, user); err != nil {
		t.Errorf("Error updating user: %v\n", err)
	}

	//reaquire the user -- by now user ought to have updated fields
	user, err = store.GetByID(user.ID)
	if err != nil {
		t.Errorf("error finding user by ID: %v\n", err)
	}

	if user.FirstName != update.FirstName {
		t.Errorf("FirstName field not updated: expected `%s` but got `%s`\n", update.FirstName, user.FirstName)
	}
	if user.LastName != update.LastName {
		t.Errorf("LastName field not updated: expected `%s` but got `%s`\n", update.LastName, user.LastName)
	}

	//gets all users in an array
	all, err := store.GetAll()
	if err != nil {
		t.Errorf("Error getting all users: %v\n", err)
	}
	if len(all) != 1 {
		t.Errorf("incorrect length of all users: expected %d but got %d\n", 1, len(all))
	}
	if all[0].ID != user.ID {
		t.Errorf("ID of user retrieved by all does not match: expected %s but got %s\n", user.ID, all[0].ID)
	}

	//TESTING THE DIETS AND ALLERGIES

	_, err = store.GetDietByName("Vegan")
	if err != nil {
		t.Errorf("error getting diet: %v\n", err)
	}

	_, err = store.GetAllergyByName("Dairy")
	if err != nil {
		t.Errorf("error getting allergy: %v\n", err)
	}

	diets := []string{"Vegan", "Pescetarian", "Vegetarian"}
	allergies := []string{"Dairy", "Egg", "Gluten", "Soy"}

	newDiets, err := store.InsertDiet(user, diets)
	if err != nil {
		t.Errorf("error inserting diets for user %v\n", err)
	}

	newDiets2, err := store.GetUserDiet(user)
	if err != nil {
		t.Errorf("error getting user's diet %v\n", err)
	}

	for i, v := range newDiets {
		if (v.UserID != newDiets2[i].UserID) && (v.DietTypeID != newDiets2[i].DietTypeID) {
			t.Errorf("error comparing diets expected user %s and got %s, expected diet %s and got diet %s", v.UserID, newDiets2[i].UserID, v.DietTypeID, newDiets2[i].DietTypeID)
		}
	}

	_, err = store.AddDiet(user, "Everything")
	if err != nil {
		t.Errorf("error adding user's diet %v\n", err)
	}
	newDiets2, err = store.GetUserDiet(user)
	if err != nil {
		t.Errorf("error getting user's diet %v\n", err)
	}
	if len(newDiets2) != 4 {
		t.Errorf("error adding new diet, expected %d but got %d", 4, len(newDiets2))
	}

	err = store.RemoveDiet(user, "Everything")
	if err != nil {
		t.Errorf("error removing user's diet %v\n", err)
	}
	newDiets2, err = store.GetUserDiet(user)
	if err != nil {
		t.Errorf("error getting user's diet %v\n", err)
	}
	if len(newDiets2) != 3 {
		t.Errorf("error removing new diet, expected %d but got %d", 4, len(newDiets2))
	}

	newAllergies, err := store.InsertAllergies(user, allergies)
	if err != nil {
		t.Errorf("error getting user's allergies %v\n", err)
	}

	newAllergies2, err := store.GetUserAllergy(user)
	if err != nil {
		t.Errorf("error receiving users allergies %v\n", err)
	}
	for i, v := range newAllergies {
		if (v.UserID != newAllergies2[i].UserID) && (v.AllergyTypeID != newAllergies2[i].AllergyTypeID) {
			t.Errorf("error comparing diets expected user %s and got %s, expected diet %s and got diet %s", v.UserID, newAllergies2[i].UserID, v.AllergyTypeID, newAllergies2[i].AllergyTypeID)
		}
	}

	_, err = store.AddAllergy(user, "None")
	if err != nil {
		t.Errorf("error adding user's allergy %v\n", err)
	}
	newAllergies2, err = store.GetUserAllergy(user)
	if err != nil {
		t.Errorf("error getting user's allergy %v\n", err)
	}
	if len(newAllergies2) != 5 {
		t.Errorf("error adding new allergy, expected %d but got %d", 5, len(newAllergies2))
	}

	err = store.RemoveAllergy(user, "None")
	if err != nil {
		t.Errorf("error removing user's allergy %v\n", err)
	}
	newAllergies2, err = store.GetUserAllergy(user)
	if err != nil {
		t.Errorf("error getting user's allergy %v\n", err)
	}
	if len(newAllergies2) != 4 {
		t.Errorf("error adding new allergy, expected %d but got %d", 4, len(newAllergies2))
	}

	// groceries := []string{"Pepper", "Cognac", "Rice", "Garlic", "Onion"}

	err = store.AddToBook(user, "Eggs")
	if err != nil {
		t.Errorf("error adding a favorite %v\n", err)
	}
	err = store.AddToBook(user, "Eggs")
	if err != nil {
		t.Errorf("error adding a favorite %v\n", err)
	}

	err = store.AddToGrocery(user, "Eggs")
	if err != nil {
		t.Errorf("error adding a grocery %v\n", err)
	}
	err = store.AddToGrocery(user, "Eggs")
	if err != nil {
		t.Errorf("error adding a grocery %v\n", err)
	}

	_, err = store.GetUserGroceries(user)
	if err != nil {
		t.Errorf("error getting user grocery list %v\n", err)
	}

	_, err = store.GetUserBook(user)
	if err != nil {
		t.Errorf("error getting user favorites %v\n", err)
	}

	err = store.DeleteFromBook(user, "Eggs")
	if err != nil {
		t.Errorf("error deleting recipe from list %v\n", err)
	}

	err = store.AddToBook(user, "Eggs")
	if err != nil {
		t.Errorf("error adding a favorite %v\n", err)
	}

	err = store.DeleteFromGrocery(user, "Eggs")
	if err != nil {
		t.Errorf("error deleting recipe from list %v\n", err)
	}

	err = store.AddToGrocery(user, "Eggs")
	if err != nil {
		t.Errorf("error adding a grocery %v\n", err)
	}

	friend := &NewUser{
		Email:        "best@best.com",
		Password:     "password",
		PasswordConf: "password",
		DOB:          "12/12/2000",
		FirstName:    "best",
		LastName:     "friend",
	}

	nu, err := store.Insert(friend)
	if err != nil {
		t.Errorf("error creating a friend %v\n", err)
	}

	_, err = store.AddFriend(user, nu)
	if err != nil {
		t.Errorf("error adding a friend %v\n", err)
	}
	err = store.AddFavFriend(user, nu)
	if err != nil {
		t.Errorf("error adding favorit %v\n", err)
	}

	anotha, err := store.Insert(friend)
	if err != nil {
		t.Errorf("error creating a friend %v\n", err)
	}

	lastFriend, err := store.Insert(friend)
	if err != nil {
		t.Errorf("error creating a friend %v\n", err)
	}
	_, err = store.AddFriend(user, anotha)
	if err != nil {
		t.Errorf("error adding a friend %v\n", err)
	}
	_, err = store.AddFriend(user, lastFriend)
	if err != nil {
		t.Errorf("error adding a friend %v\n", err)
	}

	list, err := store.GetUserFriendsList(user)
	if err != nil {
		t.Errorf("error getting my friends %v\n", err)
	}
	newList, err := store.GetUserFavFriends(user)
	if err != nil {
		t.Errorf("error getting my friends %v\n", err)
	}

	if len(list) != 3 {
		t.Errorf("error getting the correct amt of friends: expected %d but got %d", 3, len(list))
	}

	if len(newList) != 1 {
		t.Errorf("error getting the correct amt of friends: expected %d but got %d", 3, len(list))
	}

	err = store.DeleteFriend(user, anotha)
	if err != nil {
		t.Errorf("error deleting friend: %v\n", err)
	}

	list, err = store.GetUserFriendsList(user)
	if err != nil {
		t.Errorf("error getting my friends %v\n", err)
	}

	if len(list) != 2 {
		t.Errorf("error getting the correct amt of friends: expected %d but got %d", 2, len(list))
	}

	err = store.RemoveFavFriend(user, nu)

	newList, err = store.GetUserFavFriends(user)
	if err != nil {
		t.Errorf("error getting my friends %v\n", err)
	}

	if len(newList) != 0 {
		t.Errorf("error getting the correct amt of friends: expected %d but got %d", 3, len(list))
	}
	// _, err = psdb.Exec("DELETE FROM users")
	// _, err = psdb.Exec("DELETE FROM user_diet_type")
	// _, err = psdb.Exec("DELETE FROM user_allergy_type")
	// _, err = psdb.Exec("DELETE FROM grocery_list")
}
