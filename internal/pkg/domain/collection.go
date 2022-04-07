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
	MovieHref   string `json:"moviehref"`
	ImgHref     string `json:"imghref"`
	Title       string `json:"title"`
	Info        string `json:"info"`
	Rating      string `json:"rating"`
	Description string `json:"description"`
}

type CollType struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	MovieList   []MovieType `json:"movieList"`
}

type CollectionsRepository interface {
	GetCollection(id uint64) (CollType, error)
	GetFeed() (FilmSelection, error)
}

type CollectionsUsecase interface {
	GetCollection(id uint64) (CollType, error)
	GetFeed() (FilmSelection, error)
}
