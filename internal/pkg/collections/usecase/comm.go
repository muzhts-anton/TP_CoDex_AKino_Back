package colusecase

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/addPreview"
)

type collectionsUsecase struct {
	collectionsRepo domain.CollectionsRepository
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

	for index := range coll.MovieList{
		coll.MovieList[index].Poster = addPreview.ToMiniCopy(coll.MovieList[index].Poster)
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
