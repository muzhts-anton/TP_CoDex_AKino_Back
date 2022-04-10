package domain

type Actor struct {
	Id           uint64 `json:"ID"`
	Imgsrc       string `json:"avatar"`
	Name         string `json:"name"`
	NameOriginal string `json:"originalName"`
	Career       string `json:"career"`
	Height       string `json:"height"`
	Birthday     string `json:"birthdate"`
	Birthplace   string `json:"birthplace"`
	Genres       string `json:"genres"`
	Total        uint64 `json:"total"`
}

type ActorBasic struct {
	Href   string `json:"href"`
	Poster string `json:"poster"`
	Name   string `json:"title"`
}

type ActorResponse struct {
	Person  Actor        `json:"actor"`
	Related []ActorBasic `json:"related"`
	Movies  []MovieBasic `json:"movies"`
}
