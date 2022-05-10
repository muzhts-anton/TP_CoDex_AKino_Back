package actrepository

const (
	queryGetActor = `
	SELECT
		id, imgsrc, name, nameoriginal, career, height,
		birthday, birthplace, total
	FROM actors a
	WHERE id = $1;
	`

	queryGetMovies = `
	SELECT
		id, poster, title, rating, info, description
	FROM movies
	JOIN movies_actors on movies_actors.movie_id = movies.id
	WHERE movies_actors.actor_id = $1;
	`

	queryGetRelated = `
	SELECT
		actors_actors.relation_id, imgsrc, name
	FROM actors_actors
	JOIN actors on actors_actors.relation_id = actors.id
	WHERE actors_actors.actor_id = $1;
	`
)