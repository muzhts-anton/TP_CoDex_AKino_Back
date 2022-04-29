package plarepository

const (
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
	RETURNING id, title;
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
	WHERE playlist_id = $1 and movie_id = $2;
	`

	queryDeletePlaylist = `
	DELETE FROM playlists
	WHERE id = $1;
	`
)
