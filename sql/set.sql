DROP TABLE IF EXISTS Users CASCADE;

DROP TABLE IF EXISTS Movies CASCADE;

DROP TABLE IF EXISTS Collections CASCADE;

DROP TABLE IF EXISTS Feeds CASCADE;

CREATE TABLE Users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(100),
    email VARCHAR(50),
    imgsrc VARCHAR(50)
);

CREATE TABLE Feeds (
    id BIGSERIAL NOT NULL PRIMARY KEY, 
    description VARCHAR(50),
    imgsrc VARCHAR(50),
    page VARCHAR(50),
    num BIGINT
);

CREATE TABLE Collections (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(50),
    description VARCHAR(200)
);

CREATE TABLE Movies (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    poster VARCHAR(50),
    title VARCHAR(50),
    rating DOUBLE PRECISION,
    info varchar(100),
    description VARCHAR(1000),
    incollection BIGINT REFERENCES Collections (id)
);
