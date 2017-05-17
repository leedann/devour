CREATE USER pgstest WITH SUPERUSER;

create table users (
	ID serial primary key,
	Email varchar(255),
	PassHash varchar(255),
    FirstName varchar(50),
    LastName varchar(50),
    DOB date,
    PhotoURL varchar(255)
);
create table friends_list (
	ID serial primary key,
	UserID integer references users(ID),
	FriendID integer references users(ID),
	FriendsSince Date,
	-- favorite friends -- perhaps coworkers/roomates/workout buddies and stuff later
	RelationshipID integer references relationship_type(ID)
);
create table relationship_type (
	ID serial primary key,
	Description varchar(8000)
);
create table grocery_list (
	ID integer primary key,
	UserID integer references users(ID),
	--users would look for ingredients and then get it populated from the api
	Ingredients varchar(8000)[]
);
create table user_like_list (
	ID serial primary key,
	UserID integer references users(ID),
	Recipes varchar(8000)[]
);
-- Supported Allergies
-- Dairy, Egg, Gluten, Peanut, Seafood, Sesame, Soy, Sulfite, Tree Nut, Wheat
create table user_allergy_type (
	ID serial primary key,
	UserID integer references users(ID),
	AllergyTypeID integer references allergy_type(ID)
);
create table allergy_type (
	ID serial primary key,
	Name varchar(255),
	Description varchar(8000)
);
-- Supported Diets
-- Lacto vegetarian, Ovo vegetarian, Pescetarian, Vegan, Vegetarian
create table user_diet_type (
	ID serial primary key,
	UserID integer references users(ID),
	DietTypeID integer references diet_type(ID),
	BeginDate date,
	EndDate date
); 
create table diet_type (
	ID serial primary key,
	Name varchar(255),
	Description varchar(8000)
);
create table event_type (
	ID serial primary key,
	Name varchar(255),
	Description varchar(8000)
);
INSERT INTO event_type (Name, Description)
VALUES ("Business", "This is a business party")

create table event_mood_type (
	ID serial primary key,
	Name varchar(255),
	Description varchar(8000)
);

create table event_attendance (
	ID serial primary key,
	EventID integer references events(ID),
	UserID integer references users(ID),
	StatusID integer references event_attendance_status(ID)
);

create table event_attendance_status (
	ID serial primary key,
	-- this would be accept reject or hosting
	AttendanceStatus varchar(255)
);
create table events (
	ID serial primary key,
	EventTypeID integer references event_type(ID),
	Name var(255),
	Description varchar(8000),
	MoodTypeID integer references event_mood_type(ID),
	StartTime timestamp,
	EndTime timestamp,
	Recipes varchar(8000)[]
);

