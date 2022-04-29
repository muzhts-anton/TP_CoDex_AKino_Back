package usrrepository

const (
	queryGetByEmail = `
	SELECT id, username, email, imgsrc, password
	FROM users
	WHERE email = $1;
	`

	queryGetById = `
	SELECT id, username, email, imgsrc
	FROM users
	WHERE id = $1;
	`

	queryAddUser = `
	INSERT INTO
		users (username, email, password)
	VALUES
		($1, $2, $3)
	RETURNING id;
	`

	queryUpdateUser = `
	UPDATE users
	SET username = $1
	WHERE id = $2;
	`

	queryGetUserRatings = `
	SELECT movie_id, rating
	FROM ratings
	WHERE user_id = $1;
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
	SELECT playlists.id, playlists.title, playlists.poster
	FROM users_playlists
	JOIN users ON users_playlists.user_id = users.id
	JOIN playlists ON users_playlists.playlist_id = playlists.id
	WHERE users_playlists.user_id = $1
	ORDER BY playlists.id;
	`
)
