package colusecase

import (
	"codex/internal/pkg/domain"
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

	return coll, err
}

func (cu collectionsUsecase) GetFeed() (domain.FeedResponse, error) {
	feed, err := cu.collectionsRepo.GetFeed()
	if err != nil {
		return domain.FeedResponse{}, err
	}

	return feed, nil
}
