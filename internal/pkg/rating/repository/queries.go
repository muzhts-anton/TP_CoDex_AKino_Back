package ratrepository

const (
	queryGetRatingCount = `
	SELECT COUNT(*)
	FROM ratings
	WHERE user_id = $1;
	`

	queryGetUserRating = `
	SELECT rating
	FROM ratings
	WHERE user_id = $1 AND movie_id = $2;
	`

	queryPostRating = `
	INSERT INTO
		ratings (user_id, movie_id, rating)
	VALUES
		($1, $2, $3);
	`

	queryChangeRating = `
	UPDATE ratings
	SET rating = $1
	WHERE user_id = $2;
	`

	queryGetMovieRating = `
	SELECT movies.rating
	FROM ratings
	JOIN movies ON ratings.movie_id = movies.id
	WHERE ratings.movie_id = $1;	
	`

	queryGetMovieVotesnum = `
	SELECT movies.votesnum
	FROM ratings
	JOIN movies ON ratings.movie_id = movies.id
	WHERE ratings.movie_id = $1;	
	`

	queryGetRatingUserCount = `
	SELECT COUNT(*)
	FROM ratings
	JOIN movies ON ratings.movie_id = movies.id
	JOIN users on ratings.user_id = users.id
	WHERE ratings.user_id = $1 and ratings.movie_id = $2;
	`

	queryGetOldRatingUser = `
	SELECT ratings.rating
	FROM ratings
	JOIN movies ON ratings.movie_id = movies.id
	JOIN users ON ratings.user_id = users.id
	WHERE ratings.user_id = $1 AND ratings.movie_id = $2;
	`

	queryIncrementVotesnum = `
	UPDATE movies
	SET votesnum = votesnum + 1
	WHERE id = $1;
	`

	querySetMovieRating = `
	UPDATE movies
	SET rating = $1
	WHERE id = $2;
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
