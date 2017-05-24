package users

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/mail"
	"os"

	"time"

	"database/sql"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

//gravatarBasePhotoURL is the base URL for Gravatar profile photos
const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"
const cost = 10

//UserID defines the type for user IDs
type UserID interface{}

//DietID defines the type of diet IDs
type DietID interface{}

//DietTypeID defines the type of diet type IDs
type DietTypeID interface{}

//AllergyTypeID defines the type of allergy type IDs
type AllergyTypeID interface{}

//UserAllergyTypeID defines the users allergies
type UserAllergyTypeID interface{}

//FriendsListID defines the friends list ID
type FriendsListID interface{}

//RelationshipID defines the type of relationship two friends have
type RelationshipID interface{}

//GroceryListID defines the users grocery list
type GroceryListID interface{}

//RecipeID represents the ID of the recipes from yummly
type RecipeID interface{}

//UserLikesListID defines the users recipe faves
type UserLikesListID interface{}

//GroceryList defines the users list of ingredients
type GroceryList struct {
	ID          GroceryListID `json:"id" bson:"_id"`
	UserID      UserID        `json:"userID"`
	Ingredients []string      `json:"ingredients"`
}

//UserLikesList defines the users favorite recipes
type UserLikesList struct {
	ID      UserLikesListID `json:"id" bson:"_id"`
	UserID  UserID          `json:"userID"`
	Recipes []string        `json:"recipeID"`
}

//UserAllergyType defines the struct that contains user allergies
type UserAllergyType struct {
	ID            UserAllergyTypeID `json:"id" bson:"_id"`
	UserID        UserAllergyTypeID `json:"userID"`
	AllergyTypeID AllergyTypeID     `json:"aTypeID"`
}

//AllergyType defines the name of the allerg
type AllergyType struct {
	ID          AllergyTypeID `json:"id" bson:"_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
}

//Diet represents the users diet
type Diet struct {
	ID         DietID      `json:"id" bson:"_id"`
	UserID     UserID      `json:"userID"`
	DietTypeID DietTypeID  `json:"dTypeID"`
	BeginDate  time.Time   `json:"beginDate"`
	EndDate    pq.NullTime `json:"endDate,omitempty"`
}

//DietType represents the mood for an event -- casual, formal etc
type DietType struct {
	ID          DietTypeID `json:"id" bson:"_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}

//FriendsList defines the struct that contains all friend info
type FriendsList struct {
	ID             FriendsListID `json:"id" bson:"_id"`
	UserID         UserID        `json:"userID"`
	FriendID       UserID        `json:"friendID"`
	FriendsSince   time.Time     `json:"friendsSince"`
	RelationshipID sql.NullInt64 `json:"relationship"`
}

//User represents a user account in the database
type User struct {
	ID        UserID    `json:"id" bson:"_id"`
	Email     string    `json:"email"`
	PassHash  []byte    `json:"-" bson:"passHash"` //stored in mongo, but never encoded to clients
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	PhotoURL  string    `json:"photoURL"`
	DOB       time.Time `json:"dob"`
}

//Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	DOB          string `json:"dob"`
}

//UserUpdates represents updates one can make to a user
type UserUpdates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//Validate validates the new user
func (nu *NewUser) Validate() error {
	//ensure Email field is a valid Email
	//HINT: use mail.ParseAddress()
	//https://golang.org/pkg/net/mail/#ParseAddress

	_, err := mail.ParseAddress(nu.Email)
	if err != nil {
		return err
	}

	//ensure Password is at least 6 chars

	if len(nu.Password) < 6 {
		return fmt.Errorf("password should be at least 6 characters")
	}

	//ensure Password and PasswordConf match

	if nu.Password != nu.PasswordConf {
		return fmt.Errorf("passwords do not match")
	}

	//if you made here, it's valid, so return nil
	return nil
}

//ToUser converts the NewUser to a User
func (nu *NewUser) ToUser() (*User, error) {
	//build the Gravatar photo URL by creating an MD5
	//hash of the new user's email address, converting
	//that to a hex string, and appending it to their base URL:
	//https://www.gravatar.com/avatar/ + hex-encoded md5 has of email
	hash := md5.New()
	emailByte := []byte(nu.Email)
	hash.Write(emailByte)
	md5Email := hex.EncodeToString(hash.Sum(nil))

	gravURL := gravatarBasePhotoURL + md5Email

	//construct a new User setting the various fields
	//but don't assign a new ID here--do that in your
	//concrete Store.Insert() method
	dob, err := time.Parse("01/02/2006", nu.DOB)
	if err != nil {
		return nil, err
	}
	usr := &User{}
	usr.DOB = dob
	usr.PhotoURL = gravURL
	userSetting(usr, nu)
	//call the User's SetPassword() method to set the password,
	//which will hash the plaintext password
	usr.SetPassword(nu.Password)
	//return the User and nil
	return usr, nil
}

//sets the various user fields to equal new user fields
//does not export
func userSetting(u *User, nu *NewUser) {
	u.Email = nu.Email
	u.FirstName = nu.FirstName
	u.LastName = nu.LastName
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	//hash the plaintext password using an adaptive
	//crytographic hashing algorithm like bcrypt
	//https://godoc.org/golang.org/x/crypto/bcrypt

	//converting password to byte
	bytePass := []byte(password)
	passHash, err := bcrypt.GenerateFromPassword(bytePass, cost)
	if err != nil {
		fmt.Printf("error hashing password: %v", err)
		os.Exit(1)
	}
	//set the User's PassHash field to the resulting hash
	u.PassHash = passHash
	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	//compare the plaintext password with the PassHash field
	//using the same hashing algorithm you used in SetPassword
	bytePass := []byte(password)
	return bcrypt.CompareHashAndPassword(u.PassHash, bytePass)
}
