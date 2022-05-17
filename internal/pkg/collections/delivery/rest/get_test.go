package coldelivery

// import (
// 	mock2 "codex/internal/pkg/collections/usecase/mock"
// 	"codex/internal/pkg/domain"
// 	"encoding/json"
// 	"fmt"
// 	"github.com/golang/mock/gomock"
// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"testing"
// )

// type testRow struct {
// 	out    string
// 	status int
// 	name   string
// 	id     uint64
// }

// var testTableSuccess = [...]testRow{
// 	{
// 		out:    `{"title":"Топ 256","description":"must see","movielist":[{"ID":"1","poster":"showshenkRedemption.webp","title":"Побег из Шоушенка","rating":"6.8","info":"1994, США. Драма.","description":"Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения."},{"ID":"2","poster":"ironman.webp","title":"Железный Человек","rating":"10.0","info":"2008, США, Канада. Фантастика, Боевик, Приключения.","description":"Миллиардер-изобретатель Тони Старк попадает в плен к Афганским террористам, которые пытаются заставить его создать оружие массового поражения. В тайне от своих захватчиков Старк конструирует высокотехнологичную киберброню, которая помогает ему сбежать. Однако по возвращении в США он узнаёт, что в совете директоров его фирмы плетётся заговор, чреватый страшными последствиями. Используя своё последнее изобретение, Старк пытается решить проблемы своей компании радикально..."}]}` + "\n",
// 		status: http.StatusOK,
// 		name:   `GetCollection work`,
// 		id:     1,
// 	},
// }
// var testTableFailure = [...]testRow{
// 	{
// 		out:    "Bad input\n" + `a`,
// 		status: http.StatusBadRequest,
// 		name:   `GetCollection work`,
// 		id:     1000,
// 	},
// }

// func TestGetCollectionSuccess(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	apiPath := "/api/v1/collections/"
// 	for _, test := range testTableSuccess {
// 		var cl domain.Collection
// 		_ = json.Unmarshal([]byte(test.out[:len(test.out)-1]), &cl)
// 		mock := mock2.NewMockCollectionsUsecase(ctrl)
// 		mock.EXPECT().GetCollection(test.id).Times(1).Return(cl, nil)
// 		handler := CollectionsHandler{CollectionsUsecase: mock}
// 		fmt.Fprintf(os.Stdout, "Test:"+test.name)
// 		bodyReader := strings.NewReader("")
// 		w := httptest.NewRecorder()

// 		colId := strconv.Itoa(int(test.id))
// 		r := httptest.NewRequest("GET", apiPath+colId, bodyReader)
// 		vars := map[string]string{
// 			"id": colId,
// 		}
// 		r = mux.SetURLVars(r, vars)
// 		handler.GetCollection(w, r)
// 		result := test.out[:len(test.out)-1]
// 		assert.Equal(t, result, w.Body.String(), "Test: "+test.name)
// 		assert.Equal(t, test.status, w.Code, "Test: "+test.name)
// 		fmt.Fprintf(os.Stdout, " done\n")
// 	}
// }

// func TestGetCollectionError(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	apiPath := "/api/v1/collections/"
// 	for _, test := range testTableFailure {
// 		var cl domain.Collection
// 		_ = json.Unmarshal([]byte(test.out[:len(test.out)-1]), &cl)
// 		mock := mock2.NewMockCollectionsUsecase(ctrl)
// 		mock.EXPECT().GetCollection(test.id).Times(1).Return(domain.Collection{}, domain.Err.ErrObj.BadInput)

// 		handler := CollectionsHandler{CollectionsUsecase: mock}
// 		bodyReader := strings.NewReader("")
// 		colId := strconv.Itoa(int(test.id))
// 		w := httptest.NewRecorder()
// 		r := httptest.NewRequest("GET", apiPath+colId, bodyReader)
// 		vars := map[string]string{
// 			"id": colId,
// 		}
// 		r = mux.SetURLVars(r, vars)
// 		handler.GetCollection(w, r)
// 		result := test.out[:len(test.out)-1]
// 		assert.Equal(t, result, w.Body.String(), "Test: "+test.name)
// 	}
// }
