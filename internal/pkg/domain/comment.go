package domain

type Comment struct {
	Imgsrc   string `json:"avatarSrc"`
	Username string `json:"username"`
	UserId   string `json:"userID"`
	Rating   string `json:"rating"`
	Date     string `json:"date"`
	Content  string `json:"content"`
	Type     string `json:"type"`
}

type CommentRepository interface {
	PostComment(movieId uint64, userId uint64, content string, comtype string) (Comment, error)
}

type CommentUsecase interface {
	PostComment(movieId uint64, userId uint64, content string, commenttype int) (Comment, error)
}
