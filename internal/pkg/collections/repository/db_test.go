package colrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/config"
	mylog "codex/internal/pkg/utils/log"
	"encoding/binary"
	"errors"
	"regexp"
	"strconv"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

type testRow struct {
	inQuery    string
	bodyString string
	out        string
	status     int
	name       string
}

func MockDatabase() (*database.DBManager, pgxmock.PgxPoolIface, error) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		mylog.Error(errors.New("failed to create mock"))
	}
	return &database.DBManager{Pool: mock}, mock, err
}

func TestGetSuccess(t *testing.T) {
	config.DevConfigStore.FromJson()
	config.ProdConfigStore.FromJson()

	_, pool, err := MockDatabase()
	assert.Equal(t, nil, err, "create a mock")


	defer pool.Close()
	mockCollections := domain.Collection{
		Title:        "Топ 256",
		Description: "must see",
		MovieList:   []domain.MovieBasic{{
			Id: "1",
			Poster: "showshenkRedemption.webp",
			Title: "Побег из Шоушенка",
			Rating: "9",
			Info: "1994, США. Драма.",
			Description: "Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения.",
		},
		{
			Id: "2",
			Poster: "ironman.webp",
			Title: "Железный Человек",
			Rating: "10",
			Info: "2008, США, Канада. Фантастика, Боевик, Приключения.",
			Description: "Миллиардер-изобретатель Тони Старк попадает в плен к Афганским террористам, которые пытаются заставить его создать оружие массового поражения. В тайне от своих захватчиков Старк конструирует высокотехнологичную киберброню, которая помогает ему сбежать. Однако по возвращении �� США он узнаёт, что в совете директоров его фирмы плетётся заговор, чреватый страшными последствиями. Используя своё последнее изобретение, Старк пытается решить проблемы своей компании радикально...",
	
		}},
	}

	countByte := make([]uint8, 8)
	binary.BigEndian.PutUint64(countByte, uint64(1))
	numByte := make([]uint8, 8)

	Num, err := strconv.ParseUint("1", 10, 64)
	if err != nil {
		mylog.Error(errors.New("failed to parse id"))
	}
	
	binary.BigEndian.PutUint64(numByte, Num)
	rowsCount := pgxmock.NewRows([]string{"count"}).AddRow(countByte)
	rowsColl := pgxmock.NewRows([]string{"Title", "Description", "Id", "Poster", "Title", "Rating", "Info", "Description"}).AddRow([]uint8(mockCollections.Title), []uint8(mockCollections.Description), []uint8(mockCollections.MovieList[0].Id), []uint8(mockCollections.MovieList[0].Poster),[]uint8(mockCollections.MovieList[0].Title), []uint8(mockCollections.MovieList[0].Rating), []uint8(mockCollections.MovieList[0].Info), []uint8(mockCollections.MovieList[0].Description))


	pool.ExpectBegin()
	pool.ExpectQuery(regexp.QuoteMeta(queryCountCollectionsMock)).WithArgs().WillReturnRows(rowsCount)
	pool.ExpectCommit()
	pool.ExpectBegin()
	pool.ExpectQuery(regexp.QuoteMeta(queryGetFeedMock)).WithArgs(uint64(10)).WillReturnRows(rowsColl)
	pool.ExpectCommit()
}