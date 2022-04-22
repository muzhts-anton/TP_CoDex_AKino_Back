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
