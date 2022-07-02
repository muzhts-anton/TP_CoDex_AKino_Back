package colrepository

const (

	queryGetCollectionBasic = `
	SELECT
		playlists.title, playlists.description, playlists.public
	FROM playlists
	WHERE playlists.id = $1;
	`

	queryGetCollectionMovies = `
	SELECT
	movies.id, movies.poster, movies.title, movies.rating, movies.info, movies.description
	FROM playlists
	JOIN playlists_movies ON playlists.id = playlists_movies.playlist_id
	JOIN movies on playlists_movies.movie_id = movies.id
	WHERE playlists.id = $1
	ORDER BY movies.rating DESC;
	`

	queryGetFeed = `
	SELECT playlists.title, playlists.poster, playlists.id
	FROM playlists
	JOIN feed ON feed.playlist_id = playlists.id
	ORDER BY feed.id
	LIMIT $1;
	`

	queryCheckUserAccess = `
	SELECT user_id
	FROM users_playlists
	WHERE playlist_id = $1
	`

	queryGetPlaylistIsPublic = `
	SELECT public
	FROM playlists
	WHERE id = $1;
	`
)
