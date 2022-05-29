package plarepository

const (
	queryCreatePlaylist = `
	INSERT INTO
		playlists (title, public)
	VALUES
		($1, $2)
	RETURNING id, title, poster, public;
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

	queryAlterPlaylistPublic = `
	UPDATE playlists 
	SET public = $2
	WHERE id = $1;
	`

	queryAlterPlaylistTitle = `
	UPDATE playlists 
	SET title = $2
	WHERE id = $1;
	`
)
