package users

import "errors"

//ErrUserNotFound is returned when the requested user is not found in the store
var ErrUserNotFound = errors.New("user not found")

//Store represents an abstract store for model.User objects.
//This interface is used by the HTTP handlers to insert new users,
//get users, and update users. This interface can be implemented
//for any persistent database you want (e.g., MongoDB, PostgreSQL, etc.)
type Store interface {
	//GetAll returns all users
	GetAll() ([]*User, error)

	//GetByID returns the User with the given ID
	GetByID(id UserID) (*User, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*User, error)

	//Insert inserts a new NewUser into the store
	//and returns a User with a newly-assigned ID
	Insert(newUser *NewUser) (*User, error)

	CreateLikesList(user *User) (*UserLikesList, error)

	CreateGroceryList(user *User) (*GroceryList, error)

	//Update applies UserUpdates to the currentUser
	Update(updates *UserUpdates, currentuser *User) error

	//Gets the corresponding diet type by the name
	GetDietByName(dietName string) (*DietType, error)

	//Gets the corresponding Allergy type by the name
	GetAllergyByName(allergyName string) (*AllergyType, error)

	//Gets the corresponding diet type by the id
	GetDietByID(id DietTypeID) (*DietType, error)

	//Gets the corresponding Allergy type by the id
	GetAllergyByID(id AllergyTypeID) (*AllergyType, error)

	//Gets all of the allergies that the user has
	GetUserAllergy(user *User) ([]*UserAllergyType, error)

	//Gets all of the diets the user is undertaking
	GetUserDiet(user *User) ([]*Diet, error)

	//Inserts (a) new Diet(s) for the user
	InsertDiet(user *User, dietNames []string) ([]*Diet, error)

	//Adds a single Diet to the user
	AddDiet(user *User, diet string) (*Diet, error)

	//Adds a single Allergy to the user
	AddAllergy(user *User, allergyName string) (*UserAllergyType, error)

	//Removes an allergy from the user
	RemoveAllergy(user *User, allergyName string) error

	//Removes a diet from the user
	RemoveDiet(user *User, dietName string) error

	//Inserts allergies of the user
	InsertAllergies(user *User, allergyNames []string) ([]*UserAllergyType, error)

	//Gets the user their grocery list
	GetUserGroceries(user *User) (*GroceryList, error)

	//Gets the users all their saved recipes
	GetUserBook(user *User) (*UserLikesList, error)

	//Adds a favorite recipe
	AddToBook(user *User, fav string) error

	//Deletes a recipe from book
	DeleteFromBook(user *User, recipe string) error

	//Inserts multiple ingredients into the list
	InsertGroceryList(user *User, list []string) (*GroceryList, error)

	//Adds an item to the grocery list
	AddToGrocery(user *User, ingredient string) error

	//Deletes an item from the grocery list
	DeleteFromGrocery(user *User, ingredient string) error

	//Adds a friend
	AddFriend(user *User, friend *User) (*FriendsList, error)

	//Adds a friend as a favorite
	AddFavFriend(user *User, friend *User) error

	//Gets the user's friends list
	GetUserFriendsList(user *User) ([]*FriendsList, error)

	//Gets all of the users favorite friends
	GetUserFavFriends(user *User) ([]*FriendsList, error)

	//Deletes a friend from the friends list
	DeleteFriend(user *User, friend *User) error

	//Removes a friend from the friend list
	RemoveFavFriend(user *User, friend *User) error
}
