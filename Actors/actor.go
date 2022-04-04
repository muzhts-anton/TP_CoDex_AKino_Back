package actors

import (
	"encoding/json"
	"net/http"
)

type ActorInfo struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	NameOrigin   string   `json:"nameorigin"`
	PicUrl       string   `json:"picurl"`
	Career       []string `json:"career"`
	Height       string   `json:"height"`
	Age          string   `json:"age"`
	Birthday     string   `json:"birthday"`
	BirthPlace   string   `json:"birthplace"`
	Gender       string   `json:"gender"`
	FamilyStatus string   `json:"familystatus"`
	FilmNum      string   `json:"filmnum"`
}

var ActorData = ActorInfo{
	"321",
	"Daniel Radcliffe",
	"Дэниэл Рэдклифф",
	"40778.jpg",
	[]string{"Актер", "Продюсер"},
	"165",
	"32",
	"1989-07-23",
	"Фулхэм, Лондон, Англия, Великобритания",
	"male",
	"-",
	"8",
}

func GetActor(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(ActorData)
	if err != nil {
		http.Error(w, "cant marshal json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
