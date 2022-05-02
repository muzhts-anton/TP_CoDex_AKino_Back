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
DROP TABLE IF EXISTS movies_actors      CASCADE;
DROP TABLE IF EXISTS actors_actors      CASCADE;
DROP TABLE IF EXISTS users_playlists    CASCADE;
DROP TABLE IF EXISTS playlists_movies   CASCADE;
DROP TABLE IF EXISTS movies_genres      CASCADE;
DROP TABLE IF EXISTS actors_genres      CASCADE;
DROP TABLE IF EXISTS announced_genres   CASCADE;
DROP TABLE IF EXISTS actors_movies      CASCADE;
DROP TABLE IF EXISTS announced_actors   CASCADE;
DROP TABLE IF EXISTS announced_announced   CASCADE;

CREATE TABLE users (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    username                            VARCHAR(50) NOT NULL,
    password                            VARCHAR(100) NOT NULL,
    email                               VARCHAR(50) NOT NULL,
    imgsrc                              VARCHAR(50) DEFAULT 'static/avatars/profile.svg'
);

-- CREATE TRIGGER t_playlists_user
-- AFTER INSERT OR UPDATE OR DELETE ON users FOR EACH ROW EXECUTE PROCEDURE add_to_playlists ();

-- CREATE OR REPLACE FUNCTION add_to_playlists(user_id BIGINT) RETURNS TRIGGER AS $$
-- BEGIN
--     INSERT INTO
--         playlists (title)
--     VALUES
--         ('Избранное')
--     RETURNING id;

--         -- ('Мне нравится');
-- END;
-- $$ LANGUAGE plpgsql;

CREATE TABLE playlists (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    title                               VARCHAR(50) NOT NULL,
    description                         VARCHAR(200),
    poster                              VARCHAR(50) DEFAULT '/bookmark.webp',
    public                              BOOLEAN DEFAULT TRUE
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
    genre                               VARCHAR(50) PRIMARY KEY,
    imgsrc                              VARCHAR(50),
    description                         VARCHAR(100) NULL,
    title                               VARCHAR(50) NULL
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
    playlist_id                         BIGINT REFERENCES playlists (id) ON DELETE CASCADE,
    CONSTRAINT users_playlists_id       PRIMARY KEY (user_id, playlist_id)
);

CREATE TABLE playlists_movies (
    playlist_id                         BIGINT REFERENCES playlists (id) ON DELETE CASCADE,
    movie_id                            BIGINT REFERENCES movies (id),
    CONSTRAINT playlists_movies_id      PRIMARY KEY (playlist_id, movie_id) 
);

CREATE TABLE movies_genres (
    movie_id                            BIGINT REFERENCES movies (id),
    genre                               VARCHAR(50) REFERENCES genres (genre),
    CONSTRAINT movies_genres_id         PRIMARY KEY (movie_id, genre)
);

CREATE TABLE actors_genres (
    actor_id                            BIGINT REFERENCES actors (id),
    genre                               VARCHAR(50) REFERENCES genres (genre),
    CONSTRAINT actors_genres_id         PRIMARY KEY (actor_id, genre)
);

CREATE TABLE announced_genres (
    announced_id                        BIGINT REFERENCES announced (id),
    genre                               VARCHAR(50) REFERENCES genres (genre),
    CONSTRAINT announced_genres_id      PRIMARY KEY (announced_id, genre)
);

CREATE TABLE announced_actors (
    announced_id                        BIGINT REFERENCES announced (id),
    actor_id                            BIGINT REFERENCES actors (id),
    CONSTRAINT announced_actors_id      PRIMARY KEY (announced_id, actor_id)
);

CREATE TABLE announced_announced (
    announced_id                        BIGINT REFERENCES movies (id),
    relation_id                         BIGINT REFERENCES movies (id),
    CONSTRAINT announced_announced_id   PRIMARY KEY (announced_id, relation_id)
);