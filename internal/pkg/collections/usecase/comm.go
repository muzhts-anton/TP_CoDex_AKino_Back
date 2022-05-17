package colusecase

import (
	"codex/internal/pkg/domain"
	"strings"
)

type collectionsUsecase struct {
	collectionsRepo domain.CollectionsRepository
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func toMiniCopy(imgSrc string) string {
	var sb strings.Builder
	reversed := Reverse(imgSrc)
	reversedAdditional := Reverse(string("_preview"))
	point := rune('.')
	for _, symbol := range reversed {
		sb.WriteString(string(symbol))
		if (symbol == point){
			sb.WriteString(reversedAdditional)
		}
	}
	return Reverse(sb.String())
}

func InitColUsc(cr domain.CollectionsRepository) domain.CollectionsUsecase {
	return &collectionsUsecase{
		collectionsRepo: cr,
	}
}

func (cu collectionsUsecase) GetCollection(id uint64) (domain.Collection, error) {
	coll, err := cu.collectionsRepo.GetCollection(id)
	if err != nil {
		return domain.Collection{}, err
	}

	for index, _ := range coll.MovieList{
		coll.MovieList[index].Poster = toMiniCopy(coll.MovieList[index].Poster)
	}


	return coll, err
}

func (cu collectionsUsecase) GetFeed() (domain.FeedResponse, error) {
	feed, err := cu.collectionsRepo.GetFeed()
	if err != nil {
		return domain.FeedResponse{}, err
	}

	return feed, nil
}
