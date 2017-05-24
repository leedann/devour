package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

//PGStore store stucture
type PGStore struct {
	DB *sql.DB
}

//GetDietByName returns the users diet by the name
func (ps *PGStore) GetDietByName(dietName string) (*DietType, error) {
	var dietType = &DietType{}
	err := ps.DB.QueryRow(`SELECT * FROM diet_type WHERE Name = $1`, dietName).Scan(&dietType.ID, &dietType.Name, &dietType.Description)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return dietType, nil
}

//GetAllergyByName gets all of the allergies of the users by name
func (ps *PGStore) GetAllergyByName(allergyName string) (*AllergyType, error) {
	var allergyType = &AllergyType{}
	err := ps.DB.QueryRow(`SELECT * FROM allergy_type WHERE Name = $1`, allergyName).Scan(&allergyType.ID, &allergyType.Name, &allergyType.Description)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return allergyType, nil
}

//GetUserAllergy gets the allergy list of the user and returns it without the names --just id
func (ps *PGStore) GetUserAllergy(user *User) ([]*UserAllergyType, error) {
	var allergies []*UserAllergyType
	rows, err := ps.DB.Query(`SELECT * FROM user_allergy_type WHERE UserID = $1`, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var allergy = &UserAllergyType{}
		if err := rows.Scan(&allergy.ID, &allergy.UserID, &allergy.AllergyTypeID); err != nil {
			return nil, err
		}
		allergies = append(allergies, allergy)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allergies, nil
}

//GetUserDiet gets and returns all of the diets that the user has without the names --just id
func (ps *PGStore) GetUserDiet(user *User) ([]*Diet, error) {
	var diets []*Diet
	rows, err := ps.DB.Query(`SELECT * FROM user_diet_type WHERE UserID = $1`, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var diet = &Diet{}
		if err := rows.Scan(&diet.ID, &diet.UserID, &diet.DietTypeID, &diet.BeginDate, &diet.EndDate); err != nil {
			return nil, err
		}
		diets = append(diets, diet)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return diets, nil
}

// need insert for both allergy and diet

//InsertDiet inserts an array of diets for the user
func (ps *PGStore) InsertDiet(user *User, dietNames []string) ([]*Diet, error) {
	var diets []*Diet
	//multiple insert statements for the length of the amount of diets (should always be at least 1)
	for _, v := range dietNames {
		var diet = &Diet{}
		diet.BeginDate = time.Now()
		diet.UserID = user.ID
		//start a transaction
		tx, err := ps.DB.Begin()
		//err if transaction could not start
		if err != nil {
			return nil, err
		}
		dName, err := ps.GetDietByName(v)
		if err != nil {
			return nil, err
		}
		diet.DietTypeID = dName.ID
		sql := `INSERT INTO user_diet_type (UserID, DietTypeID, BeginDate) VALUES ($1, $2, $3) RETURNING id`
		//Receives ONE row from the database
		row := tx.QueryRow(sql, diet.UserID, diet.DietTypeID, diet.BeginDate)
		//scans the value of ID returned from query INTO the user
		err = row.Scan(&diet.ID)
		//err if cant scan -- rollback transaction
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		//commits the transaction-- connection no longer reserved
		diets = append(diets, diet)
		tx.Commit()
	}
	return diets, nil
}

//InsertAllergies inserts the users allergies in the relational table
func (ps *PGStore) InsertAllergies(user *User, allergyNames []string) ([]*UserAllergyType, error) {
	var allergies []*UserAllergyType
	//multiple insert statements for the length of the amount of diets (should always be at least 1)
	for _, v := range allergyNames {
		var allergy = &UserAllergyType{}
		allergy.UserID = user.ID
		//start a transaction
		tx, err := ps.DB.Begin()
		//err if transaction could not start
		if err != nil {
			return nil, err
		}
		aName, err := ps.GetAllergyByName(v)
		if err != nil {
			return nil, err
		}
		allergy.AllergyTypeID = aName.ID
		sql := `INSERT INTO user_allergy_type (UserID, AllergyTypeID) VALUES ($1, $2) RETURNING id`
		//Receives ONE row from the database
		row := tx.QueryRow(sql, allergy.UserID, allergy.AllergyTypeID)
		//scans the value of ID returned from query INTO the user
		err = row.Scan(&allergy.ID)
		//err if cant scan -- rollback transaction
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		//commits the transaction-- connection no longer reserved
		allergies = append(allergies, allergy)
		tx.Commit()
	}
	return allergies, nil
}

//AddDiet adds to the user's Diet
func (ps *PGStore) AddDiet(user *User, diet string) (*Diet, error) {
	var d = &Diet{}
	d.BeginDate = time.Now()
	d.UserID = user.ID
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}
	dName, err := ps.GetDietByName(diet)
	if err != nil {
		return nil, err
	}
	d.DietTypeID = dName.ID
	sql := `INSERT INTO user_diet_type (UserID, DietTypeID, BeginDate) VALUES ($1, $2, $3) RETURNING id`
	//Receives ONE row from the database
	row := tx.QueryRow(sql, d.UserID, d.DietTypeID, d.BeginDate)
	//scans the value of ID returned from query INTO the user
	err = row.Scan(&d.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()

	return d, nil
}

//AddAllergy adds an allergy
func (ps *PGStore) AddAllergy(user *User, allergyName string) (*UserAllergyType, error) {
	var allergy = &UserAllergyType{}
	allergy.UserID = user.ID
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}
	aName, err := ps.GetAllergyByName(allergyName)
	if err != nil {
		return nil, err
	}
	allergy.AllergyTypeID = aName.ID
	sql := `INSERT INTO user_allergy_type (UserID, AllergyTypeID) VALUES ($1, $2) RETURNING id`
	//Receives ONE row from the database
	row := tx.QueryRow(sql, allergy.UserID, allergy.AllergyTypeID)
	//scans the value of ID returned from query INTO the user
	err = row.Scan(&allergy.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	return allergy, nil
}

//RemoveAllergy removes a users allergy
func (ps *PGStore) RemoveAllergy(user *User, allergyName string) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	aName, err := ps.GetAllergyByName(allergyName)
	sql := `DELETE FROM user_allergy_type WHERE userid = $1 AND allergytypeid = $2`
	//executes the sql query
	_, err = tx.Exec(sql, user.ID, aName.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//RemoveDiet removes a diet from a user
func (ps *PGStore) RemoveDiet(user *User, dietName string) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	dName, err := ps.GetDietByName(dietName)
	sql := `DELETE FROM user_diet_type WHERE userid = $1 AND diettypeid = $2`
	//executes the sql query
	_, err = tx.Exec(sql, user.ID, dName.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//GetUserGroceries returns the list of items in the user's grocery list
func (ps *PGStore) GetUserGroceries(user *User) (*GroceryList, error) {
	var gList = &GroceryList{}
	var ingredients string
	var ingredientsarr []string
	err := ps.DB.QueryRow(`SELECT id, userid, array_to_json(ingredients) FROM grocery_list WHERE UserID = $1`, user.ID).Scan(&gList.ID, &gList.UserID, &ingredients)
	json.Unmarshal([]byte(ingredients), &ingredientsarr)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	gList.Ingredients = ingredientsarr
	return gList, nil
}

//GetUserFavorite returns all of the user's favorites
func (ps *PGStore) GetUserFavorite(user *User) (*UserLikesList, error) {
	var ulList = &UserLikesList{}
	var recipes string
	var recipesarr []string
	err := ps.DB.QueryRow(`SELECT id, userid, array_to_json(recipes) FROM user_like_list WHERE UserID = $1`, user.ID).Scan(&ulList.ID, &ulList.UserID, &recipes)
	json.Unmarshal([]byte(recipes), &recipesarr)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	ulList.Recipes = recipesarr
	return ulList, nil
}

//InsertGroceryList inserts a list of the users specified ingredients in the database
func (ps *PGStore) InsertGroceryList(user *User, list []string) (*GroceryList, error) {
	var gList = &GroceryList{}
	gList.UserID = user.ID
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}
	formattedList := formatArray(list)
	gList.Ingredients = list

	sql := `INSERT INTO grocery_list (UserID, Ingredients) VALUES ($1, $2) RETURNING id`
	//Receives ONE row from the database
	row := tx.QueryRow(sql, user.ID, formattedList)
	//scans the value of ID returned from query INTO the user
	err = row.Scan(&gList.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	return gList, nil
}

//AddToBook adds a book to the users like list
func (ps *PGStore) AddToBook(user *User, fav string) error {
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return err
	}

	newFav := "{" + fav + "}"

	sql := `UPDATE user_like_list SET recipes = array_append_distinct(recipes, $1) WHERE userid = $2`
	//executes the sql query
	_, err = tx.Exec(sql, newFav, user.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//DeleteFromBook deletes a recipe from the user's list
func (ps *PGStore) DeleteFromBook(user *User, recipe string) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}

	sql := `UPDATE user_like_list SET recipes = array_remove(recipes, $1) WHERE userid = $2`
	//executes the sql query
	_, err = tx.Exec(sql, recipe, user.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//AddToGrocery adds an item to the grocery list
func (ps *PGStore) AddToGrocery(user *User, ingredient string) error {
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return err
	}

	newIng := "{" + ingredient + "}"

	sql := `UPDATE grocery_list SET ingredients = array_append_distinct(ingredients, $1) WHERE userid = $2`
	//executes the sql query
	_, err = tx.Exec(sql, newIng, user.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//DeleteFromGrocery deletes an item from the grocery list
func (ps *PGStore) DeleteFromGrocery(user *User, ingredient string) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}

	sql := `UPDATE grocery_list SET ingredients = array_remove(ingredients, $1) WHERE userid = $2`
	//executes the sql query
	_, err = tx.Exec(sql, ingredient, user.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//CreateLikesList initiates the users favorites list
func (ps *PGStore) createLikesList(user *User) (*UserLikesList, error) {
	var ulList = &UserLikesList{}
	ulList.UserID = user.ID
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}
	sql := `INSERT INTO user_like_list (UserID) VALUES ($1) RETURNING id`
	//Receives ONE row from the database
	row := tx.QueryRow(sql, user.ID)
	//scans the value of ID returned from query INTO the user
	err = row.Scan(&ulList.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	return ulList, nil
}

//CreateGroceryList initiates the users grocery list
func (ps *PGStore) createGroceryList(user *User) (*GroceryList, error) {
	var gList = &GroceryList{}
	gList.UserID = user.ID
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}
	sql := `INSERT INTO grocery_list (UserID) VALUES ($1) RETURNING id`
	//Receives ONE row from the database
	row := tx.QueryRow(sql, user.ID)
	//scans the value of ID returned from query INTO the user
	err = row.Scan(&gList.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	return gList, nil
}

//formats an array into an acceptable insert for a string array
func formatArray(arr []string) string {
	var fmted = "{"
	for i, v := range arr {
		fmted += v
		if i+1 != len(arr) {
			fmted += ","
		}
	}
	fmted += "}"
	return fmted
}

//AddFriend adds a single friend
func (ps *PGStore) AddFriend(user *User, friend *User) (*FriendsList, error) {
	var friendRow = &FriendsList{}
	friendRow.UserID = user.ID
	friendRow.FriendID = friend.ID
	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}
	today := time.Now()
	sql := `INSERT INTO friends_list (UserID, FriendID, FriendsSince) VALUES ($1, $2, $3) RETURNING id, FriendsSince`
	//Receives ONE row from the database
	row := tx.QueryRow(sql, user.ID, friend.ID, today)
	//scans the value of ID returned from query INTO the user
	err = row.Scan(&friendRow.ID, &friendRow.FriendsSince)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	return friendRow, nil
}

//AddFavFriend adds a friend to your favorites
func (ps *PGStore) AddFavFriend(user *User, friend *User) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}

	sql := `UPDATE friends_list SET RelationshipID = (SELECT id FROM relationship_type WHERE description = $1) WHERE userid = $2 AND friendid = $3`
	//executes the sql query
	_, err = tx.Exec(sql, "Favorite", user.ID, friend.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//GetUserFriendsList returns all of the user's friends
func (ps *PGStore) GetUserFriendsList(user *User) ([]*FriendsList, error) {
	var friendsList = []*FriendsList{}
	rows, err := ps.DB.Query(`SELECT UserID, FriendID, FriendsSince, RelationshipID FROM users A INNER JOIN friends_list B ON A.ID = B.UserID WHERE A.ID = $1`, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var friendRow = &FriendsList{}
		friendRow.UserID = user.ID
		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&friendRow.ID, &friendRow.FriendID, &friendRow.FriendsSince, &friendRow.RelationshipID); err != nil {
			return nil, err
		}
		//adds to array
		friendsList = append(friendsList, friendRow)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return friendsList, nil
}

//GetUserFavFriends returns all of the user's favorite friends
func (ps *PGStore) GetUserFavFriends(user *User) ([]*FriendsList, error) {
	var friendsList = []*FriendsList{}
	rows, err := ps.DB.Query(`SELECT UserID, FriendID, FriendsSince, RelationshipID FROM users A INNER JOIN friends_list B ON A.ID = B.UserID INNER JOIN relationship_type C ON B.RelationshipID = C.ID WHERE A.ID = $1 AND C.Description = $2`, user.ID, "Favorite")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var friendRow = &FriendsList{}
		friendRow.UserID = user.ID
		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&friendRow.ID, &friendRow.FriendID, &friendRow.FriendsSince, &friendRow.RelationshipID); err != nil {
			return nil, err
		}
		//adds to array
		friendsList = append(friendsList, friendRow)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return friendsList, nil
}

//DeleteFriend deletes a friend
func (ps *PGStore) DeleteFriend(user *User, friend *User) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	sql := `DELETE FROM friends_list WHERE userid = $1 AND friendid = $2`
	//executes the sql query
	_, err = tx.Exec(sql, user.ID, friend.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

//RemoveFavFriend removes a friend from the favorites list
func (ps *PGStore) RemoveFavFriend(user *User, friend *User) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	sql := `UPDATE friends_list SET RelationshipID = NULL WHERE userid = $1 AND friendid = $2`
	//executes the sql query
	_, err = tx.Exec(sql, user.ID, friend.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}

// GetAll returns all users
func (ps *PGStore) GetAll() ([]*User, error) {
	var users []*User
	//Query the database to return multiple rows
	rows, err := ps.DB.Query(`SELECT ID, Email, FirstName, LastName, PassHash, PhotoURL, DOB FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//Next refers to the first row initially
	//returns false once EOF
	for rows.Next() {
		var user = &User{}
		//scans values into User struct; error returned if scan unsuccessful
		if err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.PassHash, &user.PhotoURL, &user.DOB); err != nil {
			return nil, err
		}
		//adds to array
		users = append(users, user)
	}
	//error is returned if encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

//GetByID returns the User with the given ID
func (ps *PGStore) GetByID(id UserID) (*User, error) {
	var user = &User{}
	//Queries and then scans; error returned if the scan unsuccessful
	err := ps.DB.QueryRow(`SELECT ID, Email, FirstName, LastName, PassHash, PhotoURL, DOB FROM users WHERE ID = $1`, id).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.PassHash, &user.PhotoURL, &user.DOB)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return user, nil
}

//GetByEmail returns the User with the given email
func (ps *PGStore) GetByEmail(email string) (*User, error) {
	var user = &User{}
	err := ps.DB.QueryRow(`SELECT ID, Email, FirstName, LastName, PassHash, PhotoURL, DOB FROM users WHERE Email = $1`, email).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.PassHash, &user.PhotoURL, &user.DOB)
	if err == sql.ErrNoRows || err != nil {
		return nil, err
	}
	return user, nil
}

//Insert inserts a new NewUser into the store
//and returns a User with a newly-assigned ID
func (ps *PGStore) Insert(newUser *NewUser) (*User, error) {
	u, err := newUser.ToUser()
	//Could not turn new user to user
	if err != nil {
		return nil, err
	}
	if nil == u {
		return nil, fmt.Errorf(".ToUser() returned nil")
	}

	//start a transaction
	tx, err := ps.DB.Begin()
	//err if transaction could not start
	if err != nil {
		return nil, err
	}
	sql := `INSERT INTO users (email, passhash, dob, firstname, lastname, photourl) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	//Receives ONE row from the database
	row := tx.QueryRow(sql, u.Email, u.PassHash, u.DOB, u.FirstName, u.LastName, u.PhotoURL)
	//scans the value of ID returned from query INTO the user
	err = row.Scan(&u.ID)
	//err if cant scan -- rollback transaction
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commits the transaction-- connection no longer reserved
	tx.Commit()
	ps.createLikesList(u)
	ps.createGroceryList(u)
	return u, nil
}

//Update applies UserUpdates to the currentUser
func (ps *PGStore) Update(updates *UserUpdates, currentuser *User) error {
	//start transaction
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}
	sql := `UPDATE users SET FirstName = $1, LastName = $2 WHERE id = $3`
	//executes the sql query
	_, err = tx.Exec(sql, updates.FirstName, updates.LastName, currentuser.ID)
	//err if could not exec, rollback transaction
	if err != nil {
		tx.Rollback()
		return err
	}
	//commits-- connection no longer reserved
	tx.Commit()
	return nil
}
