package movrepository

const (
	queryGetMovie = `
	SELECT
		id, poster, title, titleoriginal, rating, info, description, trailer,
		releaseyear, country, genre, motto, director, budget, gross, duration
	FROM
		movies
	WHERE
		id = $1;
	`

	queryGetRelated = `
	SELECT movies.id, movies.poster, movies.title
	FROM movies_movies
	JOIN movies ON movies_movies.relation_id = movies.id
	WHERE movies_movies.movie_id = $1;
	`

	queryGetComment = `
	SELECT
		users.imgsrc, users.username, users.id, comments.commentdate,
		comments.content, comments.commenttype, ratings.rating
	FROM comments
	JOIN movies ON comments.movie_id = movies.id
	JOIN users on comments.user_id = users.id
	LEFT JOIN ratings ON comments.user_id = ratings.user_id
	WHERE movies.id = $1;
	`

	queryGetRatingCount = `
	SELECT COUNT(*)
	FROM ratings
	WHERE user_id = $1;
	`
)