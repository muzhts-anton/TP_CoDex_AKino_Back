package collections

import (
	constants "codex/Constants"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type testRow struct {
	inQuery string
	out     string
	status  int
	name    string
}

var testTableSuccess = [...]testRow{
	{
		inQuery: "1",
		out:     `{"title":"Топ 256","description":"Вот такая вот подборочка :)","movieList":[{"movieHref":"/","imgHref":"greenMile.png","title":"Зелёная миля","info":"1999, США. Драма","rating":"9.1","description":"Пол Эджкомб — начальник блока смертников в тюрьме «Холодная гора», каждый из узников которого однажды проходит «зеленую милю» по пути к месту казни. Пол повидал много заключённых и надзирателей за время работы. Однако гигант Джон Коффи, обвинённый в страшном преступлении, стал одним из самых необычных обитателей блока."},{"movieHref":"/","imgHref":"showshenkRedemption.png","title":"Побег из Шоушенка","info":"1994, США. Драма","rating":"8.9","description":"Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения."}]}`,
		status:  http.StatusOK,
		name:    `limit works`,
	},
}

var testTableFailure = [...]testRow{
	{
		inQuery: "-1",
		out:     constants.ErrParseID + "\n",
		status:  http.StatusBadRequest,
		name:    `negative skip`,
	},
}

func TestGetCollectionsSuccess(t *testing.T) {
	apiPath := "/api/v1/collections/collection/"
	for _, test := range testTableSuccess {
		fmt.Fprintf(os.Stdout, "Test:"+test.name)
		bodyReader := strings.NewReader("")
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", apiPath+test.inQuery, bodyReader)
		vars := map[string]string{
			"id": test.inQuery,
		}
		r = mux.SetURLVars(r, vars)
		GetCol(w, r)
		assert.Equal(t, test.out, w.Body.String(), "Test: "+test.name)
		assert.Equal(t, test.status, w.Code, "Test: "+test.name)
		fmt.Fprintf(os.Stdout, " done\n")
	}
}

func TestGetCollectionsFailure(t *testing.T) {
	apiPath := "/api/collections/getCollections?"
	for _, test := range testTableFailure {
		fmt.Fprintf(os.Stdout, "Test:"+test.name)
		bodyReader := strings.NewReader("")
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", apiPath+test.inQuery, bodyReader)
		GetCol(w, r)
		assert.Equal(t, test.out, w.Body.String(), "Test: "+test.name)
		assert.Equal(t, test.status, w.Code, "Test: "+test.name)
		fmt.Fprintf(os.Stdout, " done\n")
	}
}
