CREATE USER pgsuser WITH SUPERUSER;
create table users (
	UserID serial primary key,
	Email varchar(255) NOT NULL,
	PassHash varchar(255),
    FirstName varchar(50),
    LastName varchar(50),
    DOB varchar(50),
    PhotoURL varchar(100)
);

create table events (
    EventID serial primary key,
    EventTypeID int4 references event_type(EventTypeID),
    EventName varchar(255) NOT NULL,
    EventDesc varchar(8000),
    EventMoodTypeID int4 references event_mood_type(EventMoodTypeID),
    EventStartTime DATETIME,
    EventEndTime DATETIME
);

create table event_type (
    EventTypeID serial primary key,
    EventTypeName varchar(255), 
    EventTypeDesc varchar(255)
)

create table event_mood_type (
    EventMoodTypeID serial primary key,
    EventMoodName varchar(255), 
    EventMoodDesc varchar(8000)
)

create table event_attendance (
    EventAttendanceID serial primary key,
    EventID int4 references events(EventID),
    UserID int4 references users(UserID), 
    StatusID int4 references event_attendance_status(StatusID)
)

create table event_attendance_status (
    StatusID serial primary key,
    -- this would be accepted, rejected, or host
    AttendanceStatus varchar(255) 
)