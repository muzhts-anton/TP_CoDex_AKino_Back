DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS feed CASCADE;
DROP TABLE IF EXISTS collections CASCADE;
DROP TABLE IF EXISTS movies CASCADE;
DROP TABLE IF EXISTS movies_actors CASCADE;
DROP TABLE IF EXISTS actors CASCADE;
DROP TABLE IF EXISTS movies_movies CASCADE;
DROP TABLE IF EXISTS actors_actors CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS ratings CASCADE;

CREATE TABLE users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(100),
    email VARCHAR(50),
    imgsrc VARCHAR(50) DEFAULT '/profile.svg'
);

CREATE TABLE feed (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    description VARCHAR(50),
    poster VARCHAR(50),
    page VARCHAR(50),
    num BIGINT
);

CREATE TABLE collections (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(50),
    description VARCHAR(200)
);

CREATE TABLE movies (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    poster VARCHAR(50),
    title VARCHAR(50),
    titleoriginal VARCHAR(50),
    rating DOUBLE PRECISION,
    votesnum BIGINT,
    info VARCHAR(100),
    description VARCHAR(1000),
    trailer VARCHAR(100),
    releaseyear VARCHAR(50),
    country VARCHAR(50),
    genre VARCHAR(50),
    motto VARCHAR(200),
    director VARCHAR(50),
    budget VARCHAR(50),
    gross VARCHAR(50),
    duration VARCHAR(50),
    incollection BIGINT REFERENCES collections (id)
);

CREATE TABLE actors (
    id BIGSERIAL NOT NULL PRIMARY KEY, 
    imgsrc VARCHAR(50),
    name VARCHAR(100),
    nameoriginal VARCHAR(100),
    career VARCHAR(100),
    height VARCHAR(50),
    birthday VARCHAR(50),
    birthplace VARCHAR(100),
    genres VARCHAR(100),
    total BIGINT
);

CREATE TABLE movies_actors (
    movie_id BIGINT REFERENCES movies (id),
    actor_id BIGINT REFERENCES actors (id),
    CONSTRAINT movies_actors_id PRIMARY KEY (movie_id, actor_id)
);

CREATE TABLE movies_movies (
    movie_id BIGINT REFERENCES movies (id),
    relation_id BIGINT REFERENCES movies (id),
    CONSTRAINT movies_movies_id PRIMARY KEY (movie_id, relation_id)
);

CREATE TABLE actors_actors (
    actor_id BIGINT REFERENCES actors (id),
    relation_id BIGINT REFERENCES actors (id),
    CONSTRAINT actors_actors_id PRIMARY KEY (actor_id, relation_id)
);

CREATE TABLE comments (
    user_id BIGINT REFERENCES users (id),
    movie_id BIGINT REFERENCES movies (id),
    commentdate VARCHAR(50),
    commenttype VARCHAR(50),
    content VARCHAR(500),
    CONSTRAINT comment_id PRIMARY KEY (movie_id, user_id)
);

CREATE TABLE ratings (
    user_id BIGINT REFERENCES users (id),
    movie_id BIGINT REFERENCES movies (id),
    rating BIGINT,
    CONSTRAINT ratings_id PRIMARY KEY (user_id, movie_id)
);