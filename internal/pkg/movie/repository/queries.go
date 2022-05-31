package movrepository

const (
	queryGetMovie = `
	SELECT
		id, poster, title, titleoriginal, rating, info, description, trailer,
		releaseyear, country, motto, director, budget, gross, duration
	FROM movies m 
	WHERE id = $1;
	`

	queryGetMovieCast = `
	SELECT actors.name, actors.id
	FROM movies_actors
	JOIN actors ON movies_actors.actor_id = actors.id
	WHERE movies_actors.movie_id = $1
	ORDER BY actors.id;
	`

	queryGetMovieGenres = `
	SELECT genres.genre, genres.title
	FROM genres
	JOIN movies_genres ON genres.genre = movies_genres.genre
	WHERE movies_genres.movie_id = $1
	ORDER BY genres.genre;
	`

	queryGetRelated = `
	SELECT movies.id, movies.poster, movies.title
	FROM movies_movies
	JOIN movies ON movies_movies.relation_id = movies.id
	WHERE movies_movies.movie_id = $1
	ORDER BY movies_movies.relation_id;
	`
)

const (
	queryGetComment = `
	SELECT
		users.imgsrc, users.username, users.id, comments.commentdate,
		comments.content, comments.commenttype
	FROM comments
	JOIN movies ON comments.movie_id = movies.id
	JOIN users on comments.user_id = users.id
	WHERE movies.id = $1
	ORDER BY comments.commentdate DESC;
	`

	queryGetCommentsCount = `
	SELECT COUNT(*)
	FROM comments
	WHERE movie_id = $1 AND user_id = $2;
	`
)

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
)

const (
	queryGetPlaylists = `
	SELECT title, id
	FROM playlists
	JOIN users_playlists ON playlists.id = users_playlists.playlist_id
	WHERE users_playlists.user_id = $1;
	`

	queryGetFilmAvailability = `
	SELECT COUNT(*)
	FROM playlists_movies
	WHERE playlists_movies.playlist_id = $1 and playlists_movies.movie_id = $2;
	`
)
