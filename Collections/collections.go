package collections

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type movieType struct {
	MovieHref   string `json:"movieHref"`
	ImgHref     string `json:"imgHref"`
	Title       string `json:"title"`
	Info        string `json:"info"`
	Rating      string `json:"rating"`
	Description string `json:"description"`
}

type collType struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	MovieList   []movieType `json:"movieList"`
}

var alabd = []collType{
	{
		Title:       "Топ 256",
		Description: "Вот такая вот подборочка :)",
		MovieList: []movieType{
			movieType{
				MovieHref:   "/",
				ImgHref:     "greenMile.png",
				Title:       "Зелёная миля",
				Info:        "1999, США. Драма",
				Rating:      "9.1",
				Description: "Пол Эджкомб — начальник блока смертников в тюрьме «Холодная гора», каждый из узников которого однажды проходит «зеленую милю» по пути к месту казни. Пол повидал много заключённых и надзирателей за время работы. Однако гигант Джон Коффи, обвинённый в страшном преступлении, стал одним из самых необычных обитателей блока.",
			},
			movieType{
				MovieHref:   "/",
				ImgHref:     "showshenkRedemption.png",
				Title:       "Побег из Шоушенка",
				Info:        "1994, США. Драма",
				Rating:      "8.9",
				Description: "Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения.",
			},
		},
	},
}

func GetCol(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	colnum, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "error BadInput", http.StatusBadRequest)
		return
	}
	if colnum != 1 {
		http.Error(w, "Im working on it", http.StatusBadRequest)
		return
	}
	jsonchik := alabd[colnum-1]
	b, err := json.Marshal(jsonchik)
	if err != nil {
		http.Error(w, "lolkek", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
