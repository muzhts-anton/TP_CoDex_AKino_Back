package usrrepository

const (

	queryGetById = `
	SELECT id, username, email, imgsrc
	FROM users
	WHERE id = $1;
	`

	queryGetUserRatings = `
	SELECT movie_id, rating
	FROM ratings
	WHERE user_id = $1;
	`
	
	queryUpdateUser = `
	UPDATE users
	SET username = $1
	WHERE id = $2;
	`

	queryGetUserComments = `
	SELECT movies.id, comments.commentdate, comments.commenttype, movies.title
	FROM comments
	JOIN movies ON comments.movie_id = movies.id
	WHERE comments.user_id = $1;
	`

	queryUpdAvatarByUsID = `
	UPDATE users 
	SET imgsrc = $2
	WHERE id = $1;
	`

	queryUserExist = `
	SELECT COUNT(*)
	FROM users
	WHERE id = $1;
	`

	queryGetUserBookmarks = `
	SELECT playlists.id, playlists.title, playlists.poster, playlists.public
	FROM users_playlists
	JOIN users ON users_playlists.user_id = users.id
	JOIN playlists ON users_playlists.playlist_id = playlists.id
	WHERE users_playlists.user_id = $1
	ORDER BY playlists.id;
	`
)
