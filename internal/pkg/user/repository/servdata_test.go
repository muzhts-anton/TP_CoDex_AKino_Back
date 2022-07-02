package usrrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/log"
	
	"encoding/binary"
	"errors"
	"regexp"
	"testing"
	// "flag"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

func MockDatabase() (*database.DBManager, pgxmock.PgxPoolIface, error) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		log.Error(errors.New("failed to create mock"))
	}
	return &database.DBManager{Pool: mock}, mock, err
}

func TestGetSuccess(t *testing.T) {
	// flag.Parse()
	mdb, pool, err := MockDatabase()
	assert.Equal(t, nil, err, "create a mock")
	repository := InitUsrRep(mdb)
	defer pool.Close()

	mu := domain.User{
		Id:             1,
		Username:       "Instance",
		Password:       "",
		Email:          "example@example.com",
		Imgsrc:         "tmp.webp",
		RepeatPassword: "",
	}
	byteId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteId, mu.Id)

	rows := pgxmock.NewRows([]string{"id", "username", "email", "imgsrc"})
	rows.AddRow(byteId, []uint8(mu.Username), []uint8(mu.Email), []uint8(mu.Imgsrc))

	pool.ExpectBegin()
	pool.ExpectQuery(regexp.QuoteMeta(queryGetById)).WithArgs(mu.Id).WillReturnRows(rows)
	pool.ExpectCommit()

	actual, err := repository.GetById(mu.Id)
	assert.NoError(t, err)
	assert.Equal(t, mu, actual)
}
