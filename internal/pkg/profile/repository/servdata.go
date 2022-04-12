package profilerepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"

	"math"
	"time"
)

type dbProfileRepository struct {
	dbm *database.DBManager
}

func InitMovRep(manager *database.DBManager) domain.ProfileRepository {
	return &dbProfileRepository{
		dbm: manager,
	}
}

func (mr *dbProfileRepository) GetProfile(id uint64) (domain.Profile, error) {
}

func (mr *dbProfileRepository) GetBookmarks(id uint64) ([]domain.BookmarksSummary, error) {

}

func (mr *dbProfileRepository) GetReviews(id uint64) ([]domain.ReviewsSummary, error) {
}

