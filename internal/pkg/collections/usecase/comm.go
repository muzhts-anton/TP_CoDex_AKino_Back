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

func (cu collectionsUsecase) GetCollection(id uint64) (domain.CollType, error) {
	coll, err := cu.collectionsRepo.GetCollection(id)
	if err != nil {
		return domain.CollType{}, err
	}

	return coll, err
}

func (cu collectionsUsecase) GetFeed() (domain.FilmSelection, error) {
	feed, err := cu.collectionsRepo.GetFeed()
	if err != nil {
		return domain.FilmSelection{}, err
	}
	
	return feed, nil
}
