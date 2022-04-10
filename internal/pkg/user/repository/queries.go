package usrrepository

const (
	queryGetByEmail = `
	SELECT id, username, email, imgsrc
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
)
