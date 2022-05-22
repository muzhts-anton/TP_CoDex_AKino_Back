package serrepository

const (
	queryGetMovies = `
	SELECT
		id, poster, title, rating, info, description
	FROM movies
	WHERE title ILIKE $1
	ORDER BY rating;
	`

	queryGetActors = `
	SELECT
		id, imgsrc, name
	FROM actors
	WHERE name ILIKE $1
	ORDER BY id;
	`

	queryGetGenres = `
	SELECT genre
	FROM genres
	WHERE genre ILIKE $1
	ORDER BY genre;
	`

	queryGetAnnounced = `
	SELECT
		id, poster, title,
		releasedate, info, description
	FROM announced
	WHERE title ILIKE $1
	ORDER BY releasedate;
	`

	queryGetBookmarks = `
	SELECT id, title, poster
	FROM playlists
	WHERE public = TRUE AND (title ILIKE $1 OR description ILIKE $1)
	ORDER BY id;
	`

	queryGetUsers = `
	SELECT id, username, imgsrc
	FROM users
	WHERE username ILIKE $1
	ORDER BY id;
	`
)
