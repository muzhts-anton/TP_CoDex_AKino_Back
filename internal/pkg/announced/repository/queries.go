package annrepository

const (
	queryGetMovies = `
	SELECT
		id, poster, title, titleoriginal, (date_part('month', releasedate))::int::varchar(255), (date_part('day', releasedate))::int::varchar(255)
	FROM announced
	ORDER BY releasedate;
	`

	queryGetAnnounced = `
	SELECT
		id, poster, title, titleoriginal, info, description, trailer,
		releasedate, country, director
	FROM announced
	WHERE id = $1;
	`

	queryGetAnnouncedCast = `
	SELECT actors.name, actors.id
	FROM announced_actors
	JOIN actors ON announced_actors.actor_id = actors.id
	WHERE announced_actors.announced_id = $1
	ORDER BY actors.id;
	`
	
	queryGetAnnouncedGenres = `
	SELECT genres.genre, genres.title
	FROM genres
	JOIN announced_genres ON genres.genre = announced_genres.genre
	WHERE announced_genres.announced_id = $1
	ORDER BY genres.genre;
	`
	
	queryGetRelated = `
	SELECT announced.id, announced.poster, announced.title
	FROM announced_announced
	JOIN announced ON announced_announced.relation_id = announced.id
	WHERE announced_announced.announced_id = $1
	ORDER BY announced_announced.relation_id;
	`
	
	queryCountAnnouncedByMonthYear = `
	SELECT COUNT(*) 
	FROM announced 
	WHERE EXTRACT(MONTH FROM releasedate) = $1 AND EXTRACT(YEAR FROM releasedate) = $2
	`
	queryGetAnnouncedsByMonthYear = `
	SELECT id, poster, title, titleoriginal, info, description, trailer, releasedate, country, director
	FROM announced 
	WHERE EXTRACT(MONTH FROM releasedate) = $1 AND EXTRACT(YEAR FROM releasedate) = $2 
	ORDER BY releasedate ASC
	`
	// TO DO EXTRACT will not work like datastamp

)
