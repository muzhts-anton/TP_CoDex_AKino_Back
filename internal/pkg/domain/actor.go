package domain

type Actor struct {
	Id           string `json:"ID"`
	Imgsrc       string `json:"avatar"`
	Name         string `json:"name"`
	NameOriginal string `json:"originalName"`
	Career       string `json:"career"`
	Height       string `json:"height"`
	Birthday     string `json:"birthdate"`
	Birthplace   string `json:"birthplace"`
	Total        string `json:"total"`
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

type ActorRepository interface {
	GetActor(id uint64) (Actor, error)
	GetMovies(id uint64) ([]MovieBasic, error)
	GetRelated(id uint64) ([]ActorBasic, error)
}

type ActorUsecase interface {
	GetActor(id uint64) (Actor, error)
	GetMovies(id uint64) ([]MovieBasic, error)
	GetRelated(id uint64) ([]ActorBasic, error)
}
