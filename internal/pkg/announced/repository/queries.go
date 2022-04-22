package annrepository

const (
	queryGetMovies = `
	SELECT
		id, poster, title,
		releasedate, info, description
	FROM announced
	ORDER BY releasedate;
	`
)