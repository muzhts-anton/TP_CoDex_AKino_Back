package movies

import (
	"encoding/json"
	"net/http"
)

type MovieInfo struct {
	Id              string   `json:"id"`
	Title           string   `json:"title"`
	TitleOrigin     string   `json:"titleorigin"`
	Rating          string   `json:"rating"`
	Description     string   `json:"description"`
	Revenue         string   `json:"revenue"`
	PosterUrl       string   `json:"posturl"`
	TrailerUrl      string   `json:"trailerurl"`
	ContentType     string   `json:"contenttype"`
	ReleaseYear     string   `json:"releaseyear"`
	Duration        string   `json:"duration"`
	Premiere        string   `json:"premiere"`
	OriginCountries []string `json:"origincountries"`
}

var MovieData = MovieInfo{
	"123",
	"Гарри Поттер и узник Азкабана",
	"Harry Potter and the Prisoner of Azkaban",
	"7.8",
	"В третьей части истории о юном волшебнике полюбившиеся всем герои — Гарри Поттер, Рон и Гермиона — возвращаются уже на третий курс школы чародейства и волшебства Хогвартс. На этот раз они должны раскрыть тайну узника, сбежавшего из зловещей тюрьмы Азкабан, чье пребывание на воле создает для Гарри смертельную опасность...",
	"$0.00",
	"322.jpg",
	"https://www.youtube.com/watch?v=WxOulfAfKvk&list=RDWxOulfAfKvk&start_radio=1",
	"film",
	"2004",
	"142",
	"2004-05-14",
	[]string{"США", "Великобритания"},
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(MovieData)
	if err != nil {
		http.Error(w, "cant marshal json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
