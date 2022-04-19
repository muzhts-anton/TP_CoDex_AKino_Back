package colrepository

const (
	queryCountCollections = `
	SELECT COUNT(*)
	FROM collections;
	`

	queryGetCollections = `
	SELECT
		playlists.title, collections.description, movies.id, movies.poster,
		movies.title, movies.rating, movies.info, movies.description
	FROM collections
	JOIN collections_movies ON collections.id = collections_movies.collection_id
	JOIN movies on collections_movies.movie_id = movies.id
	JOIN playlists ON collections.id = playlists.id
	WHERE collections.id = $1;
	`

	queryGetFeed = `
	SELECT playlists.title, playlists.poster, playlists.id
	FROM playlists
	JOIN feed ON feed.playlist_id = playlists.id
	ORDER BY feed.id
	LIMIT $1;
	`
)
