CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    event_time TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS Files (
    id SERIAL PRIMARY KEY,
    file_name VARCHAR(255),
    event_id UUID REFERENCES Events(id)
);

-- type Event struct {
-- 	Id          uuid.UUID `json:"id"`
-- 	UserId      int       `json:"user_id"`
-- 	Title       string    `json:"title"`
-- 	Description string    `json:"description"`
-- 	EventTime   time.Time `json:"event_time"`
-- 	Files       []string  `json:"files"`
-- }

