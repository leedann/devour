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

CREATE OR REPLACE FUNCTION array_append_distinct(anyarray, anyarray)
RETURNS anyarray AS $$
  SELECT ARRAY(SELECT unnest($1) union SELECT unnest($2))
$$ LANGUAGE sql;

create table relationship_type (
	ID serial primary key,
	Description varchar(8000)
);


INSERT INTO relationship_type (Description)
VALUES ('Favorite');

create table friends_list (
	ID serial primary key,
	UserID integer references users(ID),
	FriendID integer references users(ID),
	FriendsSince Date,
	-- favorite friends -- perhaps coworkers/roomates/workout buddies and stuff later
	RelationshipID integer references relationship_type(ID)
);

create table grocery_list (
	ID serial primary key,
	UserID integer references users(ID),
	--users would look for ingredients and then get it populated from the api
	Ingredients varchar(8000)[]
);

-- the recipes are the string IDs from yummlys api
create table user_like_list (
	ID serial primary key,
	UserID integer references users(ID),
	Recipes varchar(8000)[]
);

create table allergy_type (
	ID serial primary key,
	Name varchar(255),
	Description varchar(8000)
);

-- Supported Allergies
-- Dairy, Egg, Gluten, Peanut, Seafood, Sesame, Soy, Sulfite, Tree Nut, Wheat
create table user_allergy_type (
	ID serial primary key,
	UserID integer references users(ID),
	AllergyTypeID integer references allergy_type(ID)
);

INSERT INTO allergy_type (Name, Description)
VALUES 
('Dairy', 'People who are lactose intolerant are missing the enzyme lactase, which breaks down lactose, a sugar found in milk and dairy products.'),
('Egg', 'Egg allergy develops when the bodys immune system becomes sensitized and overreacts to proteins in egg whites or yolks.'),
('Gluten', 'This represents three different medical conditions that could explain what’s going on: celiac disease, wheat allergy, or non-celiac gluten sensitivity (NCGS).'),
('Peanut', 'Allergic to nuts'),
('Seafood', 'Allergic to seafood'),
('Sesame', 'Allergic to sesame'),
('Soy', 'Allergic to soy'),
('Sulfite', 'Allergic to sulfite'),
('Tree Nut', 'Allergic to tree nut'),
('Wheat', 'Allergic to Wheat'),
('None', 'No allergies');

create table diet_type (
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

INSERT INTO diet_type (Name, Description)
VALUES 
('Lacto vegetarian', 'Abstaining from meat and eggs, but will eat dairy'),
('Ovo vegetarian', 'Will not eat meat or dairy, but will eat eggs'),
('Pescetarian', 'Will not eat meat, but will eat fish'),
('Vegan', 'Does not eat or use animal products'),
('Vegetarian', 'Does not eat meat, and sometimes other animal products.'),
('Everything', 'No restrictions!');

create table event_type (
	ID serial primary key,
	Name varchar(255),
	Description varchar(8000)
);

INSERT INTO event_type (Name, Description)
VALUES 
('Formal', 'this is a formal event'),
('Semi-Formal', 'This party is for semi-formal event'),
('Casual', 'This is a casual event'),
('Festive', 'This is a festive occasion!'),
('Black Tie', 'This is a black tie event'),
('White Tie', 'This is a white tie event'),
('Other', 'This event type is not specified!');


create table event_mood_type (
	ID serial primary key,
	Name varchar(255),
	Description varchar(8000)
);

INSERT INTO event_mood_type (Name, Description)
VALUES 
('Fun', 'The mood will be fun'),
('Silly', 'The mood is silly'),
('Fancy', 'The mood is fancy'),
('Relaxed', 'The mood is relaxed'),
('Focused', 'The mood is focused');

create table event_attendance_status (
	ID serial primary key,
	-- this would be accept reject or hosting
	AttendanceStatus varchar(255)
);

INSERT INTO event_attendance_status (AttendanceStatus)
VALUES 
('Pending'),
('Attending'),
('Not Attending'),
('Host'),
('Maybe');

create table events (
	ID serial primary key,
	EventTypeID integer references event_type(ID),
	Name varchar(255),
	Description varchar(8000),
	MoodTypeID integer references event_mood_type(ID),
	StartTime timestamp,
	EndTime timestamp
);

create table recipe_suggestions (
	ID serial primary key,
	EventID integer references events(ID),
	UserID integer references users(ID),
	Recipe varchar(255)
);

create table event_attendance (
	ID serial primary key,
	EventID integer references events(ID),
	UserID integer references users(ID),
	StatusID integer references event_attendance_status(ID)
);


