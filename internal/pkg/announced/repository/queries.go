package annrepository

const (
	queryGetMovies = `
	SELECT
		id, poster, title,
		releasedate, (EXTRACT(epoch FROM (SELECT (releasedate - CURRENT_TIMESTAMP)))/86400)::int::varchar(255), info, description
	FROM announced
	ORDER BY releasedate;
	`
)