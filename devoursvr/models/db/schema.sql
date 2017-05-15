CREATE USER pgstest WITH SUPERUSER;
create table users (
	ID serial primary key,
	Email varchar(255),
	PassHash varchar(255),
	UserName varchar(100),
    FirstName varchar(50),
    LastName varchar(50),
    PhotoURL varchar(100),
    MobilePhone varchar(12)
);
create table channels (
	ID          serial primary key,
	Name        varchar(255),
	Description varchar(8000),
	CreatedAt   timestamp,
	CreatorID   integer REFERENCES users(ID),
	Members     integer[],
	Private     boolean
);
create table messages (
	ID        serial primary key,
	ChannelID integer REFERENCES channels(ID),
	Body      varchar(8000),
	CreatedAt timestamp,
	CreatorID integer REFERENCES users(ID),
	EditedAt  timestamp
);
