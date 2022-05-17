package genrepository

const (
	queryGetGenreWithMovies = `
	SELECT
		movies.id, movies.poster, movies.title,
		movies.rating, movies.info, movies.description,
		genres.description, genres.title
	FROM movies_genres
	JOIN movies ON movies_genres.movie_id = movies.id
	JOIN genres ON movies_genres.genre = genres.genre
	WHERE movies_genres.genre = $1
	ORDER BY movies.rating DESC
	LIMIT $2;
	`

	queryGetGenres = `
	SELECT
		genre, imgsrc
	FROM genres
	ORDER BY genre;
	`
)
