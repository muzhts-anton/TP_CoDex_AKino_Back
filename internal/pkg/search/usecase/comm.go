package serusecase

import (
	"codex/internal/pkg/domain"
)

type searchUsecase struct {
	searchRepo domain.SearchRepository
}

func InitSerUsc(ar domain.SearchRepository) domain.SearchUsecase {
	return &searchUsecase{
		searchRepo: ar,
	}
}

func (gu searchUsecase) Search(tag string) (domain.SearchResponse, error) {
	movs, err := gu.searchRepo.SearchMovies(tag)
	if err != nil {
		return domain.SearchResponse{}, err
	}

	gens, err := gu.searchRepo.SearchGenres(tag)
	if err != nil {
		return domain.SearchResponse{}, err
	}

	acts, err := gu.searchRepo.SearchActors(tag)
	if err != nil {
		return domain.SearchResponse{}, err
	}

	anns, err := gu.searchRepo.SearchAnnounced(tag)
	if err != nil {
		return domain.SearchResponse{}, err
	}

	bkms, err := gu.searchRepo.SearchBookmarks(tag)
	if err != nil {
		return domain.SearchResponse{}, err
	}

	usrs, err := gu.searchRepo.SearchUsers(tag)
	if err != nil {
		return domain.SearchResponse{}, err
	}

	return domain.SearchResponse{
		Movies:    movs,
		Genres:    gens,
		Actors:    acts,
		Announced: anns,
		Bookmarks: bkms,
		Users:     usrs,
	}, nil
}
