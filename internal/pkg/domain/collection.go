package domain

type FeedRow struct {
	Description string `json:"description"`
	ImgSrc      string `json:"imgsrc"`
	Page        string `json:"page"`
	Num         string `json:"number"`
}

type Feed struct {
	Coll []FeedRow `json:"collectionlist"`
}

type MovieRow struct {
	Id          string `json:"ID"`
	ImgHref     string `json:"poster"`
	Title       string `json:"title"`
	Rating      string `json:"rating"`
	Info        string `json:"info"`
	Description string `json:"description"`
}

type Collection struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	MovieList   []MovieRow `json:"movielist"`
}

type CollectionsRepository interface {
	GetCollection(id uint64) (Collection, error)
	GetFeed() (Feed, error)
}

type CollectionsUsecase interface {
	GetCollection(id uint64) (Collection, error)
	GetFeed() (Feed, error)
}
