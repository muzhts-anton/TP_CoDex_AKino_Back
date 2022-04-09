package domain

type FilmType struct {
	Description string `json:"description"`
	ImgSrc      string `json:"imgsrc"`
	Page        string `json:"page"`
	Number      string `json:"number"`
}

type FilmSelection struct {
	Coll []FilmType `json:"collectionlist"`
}

type MovieType struct {
	Id          string `json:"ID"`
	ImgHref     string `json:"poster"`
	Title       string `json:"title"`
	Rating      string `json:"rating"`
	Info        string `json:"info"`
	Description string `json:"description"`
}

type CollType struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	MovieList   []MovieType `json:"movielist"`
}

type CollectionsRepository interface {
	GetCollection(id uint64) (CollType, error)
	GetFeed() (FilmSelection, error)
}

type CollectionsUsecase interface {
	GetCollection(id uint64) (CollType, error)
	GetFeed() (FilmSelection, error)
}
