package actusecase

import (
	"codex/internal/pkg/domain"
)

type actorUsecase struct {
	actorRepo domain.ActorRepository
}

func InitActUsc(ar domain.ActorRepository) domain.ActorUsecase {
	return &actorUsecase{
		actorRepo: ar,
	}
}

func (au actorUsecase) GetActor(id uint64) (domain.Actor, error) {
	actor, err := au.actorRepo.GetActor(id)
	if err != nil {
		return domain.Actor{}, err
	}

	return actor, err
}

func (au actorUsecase) GetMovies(id uint64) ([]domain.MovieBasic, error) {
	movie, err := au.actorRepo.GetMovies(id)
	if err != nil {
		return nil, err
	}

	return movie, err
}

func (au actorUsecase) GetRelated(id uint64) ([]domain.ActorBasic, error) {
	actors, err := au.actorRepo.GetRelated(id)
	if err != nil {
		return nil, err
	}

	return actors, err
}
