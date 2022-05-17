package domain

type Feed struct {
	Description string `json:"description"`
	ImgSrc      string `json:"imgsrc"`
	Page        string `json:"page"`
	Num         string `json:"number"`
}

type FeedResponse struct {
	CollectionList []Feed `json:"collectionlist"`
}

type Collection struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Public      bool         `json:"public"`
	UserId      string       `json:"userId"`
	MovieList   []MovieBasic `json:"movielist"`

}

type CollectionsRepository interface {
	GetCollection(id uint64) (Collection, error)
	GetFeed() (FeedResponse, error)
}

// mockgen -destination=../collections/usecase/mock/usecase_mock.go  -package=mock codex/internal/pkg/domain CollectionsUsecase
type CollectionsUsecase interface {
	GetCollection(id uint64) (Collection, error)
	GetFeed() (FeedResponse, error)
}
