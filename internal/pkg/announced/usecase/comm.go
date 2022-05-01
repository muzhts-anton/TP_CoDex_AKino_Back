package annusecase

import (
	"codex/internal/pkg/domain"
)

type announcedUsecase struct {
	announcedRepo domain.AnnouncedRepository
}

func InitAnnUsc(ar domain.AnnouncedRepository) domain.AnnouncedUsecase {
	return &announcedUsecase{
		announcedRepo: ar,
	}
}

func (gu announcedUsecase) GetMovies() ([]domain.AnnouncedBasic, error) {
	movs, err := gu.announcedRepo.GetMovies()
	if err != nil {
		return []domain.AnnouncedBasic{}, err
	}

	return movs, err
}

func (gu announcedUsecase) GetMovie(id uint64) (domain.Announced, error) {
	announced, err := gu.announcedRepo.GetMovie(id)
	if err != nil {
		return domain.Announced{}, err
	}

	return announced, err
}

func (gu announcedUsecase) GetRelated(id uint64) ([]domain.AnnouncedSummary, error) {
	related, err := gu.announcedRepo.GetRelated(id)
	if err != nil {
		return nil, err
	}

	return related, err
}
