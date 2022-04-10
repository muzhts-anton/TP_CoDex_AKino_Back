package colrepository

const (
	queryCountCollections = `
	SELECT COUNT(*)
	FROM collections;
	`

	queryGetCollections = `
	SELECT
		collections.title, collections.description, movies.id, movies.poster,
		movies.title, movies.rating, movies.info, movies.description
	FROM collections
	JOIN movies on collections.id = movies.incollection
	WHERE collections.id = $1;
	`

	queryGetFeed = `
	SELECT description, poster, page, num
	FROM feed;
	`
)
