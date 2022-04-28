package playlistrepository

const (
	queryGetUserBookmarks = `
	SELECT playlists.id, playlists.title, playlists.poster
	FROM users_playlists
	JOIN users ON users_playlists.user_id = users.id
	JOIN playlists ON users_playlists.playlist_id = playlists.id
	WHERE users_playlists.user_id = $1
	ORDER BY playlists.id;
	`
	
	queryGetCollections = `
	SELECT
		playlists.title, playlists.description, movies.id, movies.poster,
		movies.title, movies.rating, movies.info, movies.description
	FROM playlists
	JOIN playlists_movies ON playlists.id = playlists_movies.playlist_id
	JOIN movies on playlists_movies.movie_id = movies.id
	WHERE playlists.id = $1
	ORDER BY movies.id;
	`

	queryGetFeed = `
	SELECT playlists.title, playlists.poster, playlists.id
	FROM playlists
	JOIN feed ON feed.playlist_id = playlists.id
	ORDER BY feed.id
	LIMIT $1;
	`

	queryCreatePlaylist = `
	INSERT INTO
		playlists (title, public)
	VALUES
		($1, $2)
	RETURNING id;
	`

	queryCreatePlaylistUser = `
	INSERT INTO
		users_playlists (user_id, playlist_id)
	VALUES
		($1, $2);
	`

	queryPlaylistExist = `
	SELECT count(*)
	FROM users_playlists 
	JOIN playlists ON playlists.id = users_playlists.playlist_id
	WHERE users_playlists.user_id = $1 and playlists.title = $2;
	`

	queryAddMovie = `
	INSERT INTO
		playlists_movies (playlist_id, movie_id)
	VALUES
		($1,$2);
	`

	queryDeleteMovie = `
	DELETE FROM playlists_movies
	WHERE playlists_id = $1 and movie_id = $2;
	`

	queryDeletePlaylist = `
	DELETE FROM playlists
	WHERE playlists_id = $1;
	`
)
