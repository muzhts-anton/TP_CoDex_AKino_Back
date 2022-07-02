package autrepository

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
	queryAddBasicPlaylists = `
    INSERT INTO
        playlists (title)
    VALUES
        ('Посмотреть позже'),
		('Мне нравится')
    RETURNING id;
	`

	queryBindBasicPlaylists = `
	INSERT INTO
    	users_playlists (user_id, playlist_id)
	VALUES
    	($1, $2),
    	($1, $3);
	`
)
