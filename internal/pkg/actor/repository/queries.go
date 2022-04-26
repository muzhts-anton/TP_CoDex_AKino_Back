package actrepository

const (
	queryGetActor = `
	SELECT
		id, imgsrc, name, nameoriginal, career, height,
		birthday, birthplace,  string_agg(ag.genre, ', '), total
	FROM actors a JOIN actors_genres ag on a.id = ag.actor_id
	WHERE id = $1
	GROUP BY a.id;
	`

	queryGetMovies = `
	SELECT
		id, poster, title, rating, info, description
	FROM movies
	WHERE id = $1;
	`

	queryGetRelated = `
	SELECT
		actors_actors.relation_id, imgsrc, name
	FROM actors_actors
	JOIN actors on actors_actors.relation_id = actors.id
	WHERE actors_actors.actor_id = $1;
	`
)