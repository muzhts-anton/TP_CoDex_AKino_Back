package genrepository

const (
	queryGetMovies = `
	SELECT
		movies.id, movies.poster, movies.title,
		movies.rating, movies.info, movies.description
	FROM movies_genres
	JOIN movies ON movies_genres.movie_id = movies.id
	WHERE movies_genres.genre = $1
	ORDER BY movies.rating
	LIMIT $2;
	`
)