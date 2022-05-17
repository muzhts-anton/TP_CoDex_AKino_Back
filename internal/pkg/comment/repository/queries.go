package comrepository

const (
	queryGetComment = `
	SELECT
		users.imgsrc, users.username, users.id, comments.commentdate,
		comments.content, comments.commenttype, ratings.rating
	FROM comments
	JOIN movies ON comments.movie_id = movies.id
	JOIN users on comments.user_id = users.id
	LEFT JOIN ratings ON comments.user_id = ratings.user_id
	WHERE movies.id = $1
	ORDER BY comments.commentdate DESC;
	`

	queryGetUserComment = `
	SELECT
		users.imgsrc, users.username, users.id, comments.commentdate,
		comments.content, comments.commenttype, ratings.rating
	FROM comments
	JOIN movies ON comments.movie_id = movies.id
	JOIN users on comments.user_id = users.id
	LEFT JOIN ratings ON comments.user_id = ratings.user_id
	WHERE movies.id = $1 AND users.id = $2;
	`

	queryGetCommentsCount = `
	SELECT COUNT(*)
	FROM comments
	WHERE movie_id = $1 AND user_id = $2;
	`

	queryPostComment = `
	INSERT INTO
		comments (user_id, movie_id, commentdate, commenttype, content)
	VALUES
		($1, $2, $3, $4, $5);
	`
)

const (
	queryGetRatingCount = `
	SELECT COUNT(*)
	FROM ratings
	WHERE user_id = $1;
	`
)

const (
	queryUserExist = `
	SELECT COUNT(*)
	FROM users
	WHERE id = $1;
	`

	queryMovieExist = `
	SELECT COUNT(*)
	FROM movies
	WHERE id = $1;
	`
)
