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
