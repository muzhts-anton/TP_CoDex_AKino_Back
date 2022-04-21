DROP TABLE IF EXISTS users              CASCADE;
DROP TABLE IF EXISTS playlists          CASCADE;
DROP TABLE IF EXISTS movies             CASCADE;
DROP TABLE IF EXISTS actors             CASCADE;
DROP TABLE IF EXISTS comments           CASCADE;
DROP TABLE IF EXISTS ratings            CASCADE;
DROP TABLE IF EXISTS announced          CASCADE;
DROP TABLE IF EXISTS genres             CASCADE;
DROP TABLE IF EXISTS feed               CASCADE;

DROP TABLE IF EXISTS movies_movies      CASCADE;
DROP TABLE IF EXISTS actors_actors      CASCADE;
DROP TABLE IF EXISTS users_playlists    CASCADE;
DROP TABLE IF EXISTS playlists_movies   CASCADE;
DROP TABLE IF EXISTS genres_movies      CASCADE;
DROP TABLE IF EXISTS genres_actors      CASCADE;
DROP TABLE IF EXISTS genres_announced   CASCADE;
DROP TABLE IF EXISTS movies_actors      CASCADE;

CREATE TABLE users (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    username                            VARCHAR(50) NOT NULL,
    password                            VARCHAR(100) NOT NULL,
    email                               VARCHAR(50) NOT NULL,
    imgsrc                              VARCHAR(50) DEFAULT '/static/avatars/profile.svg'
);

CREATE TABLE playlists (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    title                               VARCHAR(50) NOT NULL,
    description                         VARCHAR(200),
    poster                              VARCHAR(50) DEFAULT '/bookmark.webp'
);

CREATE TABLE feed (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    playlist_id                         BIGINT REFERENCES playlists (id)
);

CREATE TABLE movies (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    poster                              VARCHAR(50) NOT NULL,
    title                               VARCHAR(50) NOT NULL,
    titleoriginal                       VARCHAR(50) NOT NULL,
    rating                              DOUBLE PRECISION NOT NULL,
    votesnum                            BIGINT NOT NULL,
    info                                VARCHAR(100) NOT NULL,
    description                         VARCHAR(1000) NOT NULL,
    trailer                             VARCHAR(100) NOT NULL,
    releaseyear                         VARCHAR(50) NOT NULL,
    country                             VARCHAR(50) NOT NULL,
    genre                               VARCHAR(50) NOT NULL, -- TODO: delete
    motto                               VARCHAR(200) NOT NULL,
    director                            VARCHAR(50) NOT NULL,
    budget                              VARCHAR(50) NOT NULL,
    gross                               VARCHAR(50) NOT NULL,
    duration                            VARCHAR(50) NOT NULL
);

CREATE TABLE announced (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    poster                              VARCHAR(50) NOT NULL,
    title                               VARCHAR(50) NOT NULL,
    titleoriginal                       VARCHAR(50) NOT NULL,
    info                                VARCHAR(100) NOT NULL,
    description                         VARCHAR(1000) NOT NULL,
    trailer                             VARCHAR(100) NOT NULL,
    releasedate                         TIMESTAMP NOT NULL,
    country                             VARCHAR(50) NOT NULL,
    genre                               VARCHAR(50) NOT NULL,
    director                            VARCHAR(50) NOT NULL
);

CREATE TABLE actors (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    imgsrc                              VARCHAR(50) NOT NULL,
    name                                VARCHAR(100) NOT NULL,
    nameoriginal                        VARCHAR(100) NOT NULL,
    career                              VARCHAR(100) NOT NULL,
    height                              VARCHAR(50) NOT NULL,
    birthday                            VARCHAR(50) NOT NULL,
    birthplace                          VARCHAR(100) NOT NULL,
    genres                              VARCHAR(100), -- TODO: delete
    total                               BIGINT NOT NULL
);

CREATE TABLE comments (
    user_id                             BIGINT REFERENCES users (id),
    movie_id                            BIGINT REFERENCES movies (id),
    commentdate                         TIMESTAMP NOT NULL,
    commenttype                         VARCHAR(50) NOT NULL,
    content                             VARCHAR(500) NOT NULL,
    CONSTRAINT comment_id               PRIMARY KEY (movie_id, user_id)
);

CREATE TABLE ratings (
    user_id                             BIGINT REFERENCES users (id),
    movie_id                            BIGINT REFERENCES movies (id),
    rating                              BIGINT NOT NULL,
    CONSTRAINT ratings_id               PRIMARY KEY (user_id, movie_id)
);

CREATE TABLE genres (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    genre                               VARCHAR(50) NOT NULL
);


CREATE TABLE movies_actors (
    movie_id                            BIGINT REFERENCES movies (id),
    actor_id                            BIGINT REFERENCES actors (id),
    CONSTRAINT movies_actors_id         PRIMARY KEY (movie_id, actor_id)
);

CREATE TABLE movies_movies (
    movie_id                            BIGINT REFERENCES movies (id),
    relation_id                         BIGINT REFERENCES movies (id),
    CONSTRAINT movies_movies_id         PRIMARY KEY (movie_id, relation_id)
);

CREATE TABLE actors_actors (
    actor_id                            BIGINT REFERENCES actors (id),
    relation_id                         BIGINT REFERENCES actors (id),
    CONSTRAINT actors_actors_id         PRIMARY KEY (actor_id, relation_id)
);

CREATE TABLE users_playlists (
    user_id                             BIGINT REFERENCES users (id),
    playlist_id                         BIGINT REFERENCES playlists (id),
    CONSTRAINT users_playlists_id       PRIMARY KEY (user_id, playlist_id)
);

CREATE TABLE playlists_movies (
    playlist_id                         BIGINT REFERENCES playlists (id),
    movie_id                            BIGINT REFERENCES movies (id),
    CONSTRAINT playlists_movies_id      PRIMARY KEY (playlist_id, movie_id)
);

CREATE TABLE genres_movies (
    genre_id                            BIGINT REFERENCES genres (id),
    movie_id                            BIGINT REFERENCES movies (id),
    CONSTRAINT genres_movies_id         PRIMARY KEY (genre_id, movie_id)
);

CREATE TABLE genres_actors (
    genre_id                            BIGINT REFERENCES genres (id),
    actor_id                            BIGINT REFERENCES actors (id),
    CONSTRAINT genres_actors_id         PRIMARY KEY (genre_id, actor_id)
);

CREATE TABLE genres_announced (
    genre_id                            BIGINT REFERENCES genres (id),
    announced_id                        BIGINT REFERENCES announced (id),
    CONSTRAINT genres_announced_id      PRIMARY KEY (genre_id, announced_id)
);
