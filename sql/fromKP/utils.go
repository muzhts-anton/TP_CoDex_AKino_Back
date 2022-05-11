package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// models

type Movie struct {
	Title         string  `json:"nameRu"`
	TitleOriginal string  `json:"nameOriginal"`
	Poster        string  `json:"posterUrlPreview"`
	Rating        float64 `json:"ratingFilmCritics"`
	Votesnum      uint64  `json:"ratingFilmCriticsVoteCount"`
	RelaeseYear   uint64  `json:"year"`
	Duration      uint64  `json:"filmLength"`
	Motto         string  `json:"slogan"`
	Desc          string  `json:"description"`
	Desc2         string  `json:"shortDescription"`
	Countries     []struct {
		Country string `json:"country"`
	} `json:"countries"`
	Genres []struct {
		Genre string `json:"genre"`
	} `json:"genres"`
}

type BoxOffice struct {
	Items []struct {
		Type   string `json:"type"`
		Amount uint64 `json:"amount"`
		Symbol string `json:"symbol"`
	} `json:"items"`
}

type Videos struct {
	Items []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"items"`
}

type Staff struct {
	StaffId uint64 `json:"staffId"`
	Name    string `json:"nameRu"`
	Key     string `json:"professionKey"`
}

type Actor struct {
	Name         string     `json:"nameRu"`
	NameOriginal string     `json:"nameEn"`
	Poster       string     `json:"posterUrl"`
	Height       uint64     `json:"growth"`
	Birthday     string     `json:"birthday"`
	Birthplace   string     `json:"birthplace"`
	Proffesion   string     `json:"profession"`
	Films        []struct{} `json:"films"`
}

type Similars struct {
	Items []struct {
		NameOriginal string `json:"nameOriginal"`
	} `json:"items"`
}

// getters

func getFilm(filmId uint64) (*Movie, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/%d", filmId), nil)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"X-API-KEY": {token[0]},
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	jsn := new(Movie)
	err = json.NewDecoder(resp.Body).Decode(&jsn)
	if err != nil {
		return nil, err
	}

	return jsn, nil
}

func getBoxOffice(filmId uint64) (*BoxOffice, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/%d/box_office", filmId), nil)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"X-API-KEY": {token[0]},
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	jsn := new(BoxOffice)
	err = json.NewDecoder(resp.Body).Decode(&jsn)
	if err != nil {
		return nil, err
	}

	return jsn, nil
}

func getVideos(filmId uint64) (*Videos, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/%d/videos", filmId), nil)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"X-API-KEY": {token[0]},
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	jsn := new(Videos)
	err = json.NewDecoder(resp.Body).Decode(&jsn)
	if err != nil {
		return nil, err
	}

	return jsn, nil
}

func getStaff(filmId uint64) (*[maxstaffresp]Staff, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v1/staff?filmId=%d", filmId), nil)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"X-API-KEY": {token[0]},
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	jsn := new([15]Staff)
	err = json.NewDecoder(resp.Body).Decode(&jsn)
	if err != nil {
		return nil, err
	}

	return jsn, nil
}

func getActor(actorId uint64) (*Actor, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v1/staff/%d", actorId), nil)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"X-API-KEY": {token[0]},
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	jsn := new(Actor)
	err = json.NewDecoder(resp.Body).Decode(&jsn)
	if err != nil {
		return nil, err
	}

	return jsn, nil
}

func getSimilars(filmId uint64) (*Similars, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://kinopoiskapiunofficial.tech/api/v2.2/films/%d/similars", filmId), nil)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"X-API-KEY": {token[0]},
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	jsn := new(Similars)
	err = json.NewDecoder(resp.Body).Decode(&jsn)
	if err != nil {
		return nil, err
	}

	return jsn, nil
}

// fillers

func fillMovie(jsnF *Movie, jsnBO *BoxOffice, jsnV *Videos, jsnST *[maxstaffresp]Staff) (uint64, [maxacts]uint64, error) {
	var countries string
	for i := range jsnF.Countries {
		if i > 0 {
			countries += ", "
		}
		countries += jsnF.Countries[i].Country
	}
	countries += "."

	var budget, gross string
	for _, item := range jsnBO.Items {
		if item.Type == "WORLD" {
			gross = item.Symbol + fmt.Sprint(item.Amount)
		} else if item.Type == "BUDGET" {
			budget = item.Symbol + fmt.Sprint(item.Amount)
		}
	}

	var trailer string = "#"
	for _, vid := range jsnV.Items {
		if vid.Name == "Трейлер (русский язык)" { // FIXME
			trailer = vid.Url
			break
		}
	}

	var directors string
	var actorIds [maxacts]uint64
	dirnum, actnum := 0, 0
	for _, staff := range *jsnST {
		if staff.Key == "DIRECTOR" && dirnum < maxdirectors {
			if dirnum > 0 {
				directors += ", "
			}
			directors += staff.Name
			dirnum++
		} else if staff.Key == "ACTOR" && actnum < maxacts {
			actorIds[actnum] = staff.StaffId
			actnum++
		}
	}
	directors += "."

	dbresp, err := dbm.Query(`
	INSERT INTO
		movies (
			poster,
			title,
			titleoriginal,
			rating,
			votesnum,
			info,
			description,    
			releaseyear, 
			country,
			motto,
			duration,
			gross,
			budget,
			trailer,
			director
	)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
	)
	RETURNING id;`,
		jsnF.Poster,
		jsnF.Title,
		jsnF.TitleOriginal,
		jsnF.Rating,
		jsnF.Votesnum,
		fmt.Sprint(jsnF.RelaeseYear)+", "+jsnF.Countries[0].Country+". "+jsnF.Desc2,
		jsnF.Desc,
		fmt.Sprint(jsnF.RelaeseYear),
		countries,
		jsnF.Motto,
		fmt.Sprint(jsnF.Duration),
		gross,
		budget,
		trailer,
		directors,
	)
	if err != nil {
		return 0, [maxacts]uint64{}, err
	}

	return binary.BigEndian.Uint64(dbresp[0][0]), actorIds, nil
}

func fillGenres(jsnF *Movie) error {
	for _, genre := range jsnF.Genres {
		dbresp, err := dbm.Query(`SELECT COUNT(*) FROM genres WHERE genre = $1;`, genre.Genre)
		if err != nil {
			return err
		}

		if binary.BigEndian.Uint64(dbresp[0][0]) == 0 {
			_, err = dbm.Query(`INSERT INTO genres (genre) VALUES ($1);`, genre.Genre)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func fillMoviesGenres(jsnF *Movie, movId uint64) error {
	for _, genre := range jsnF.Genres {
		_, err := dbm.Query(`INSERT INTO movies_genres (movie_id, genre) VALUES ($1, $2);`, movId, genre.Genre)
		if err != nil {
			return err
		}
	}

	return nil
}

func fillActors(jsnAC *Actor) (uint64, error) {
	dbresp, err := dbm.Query(`SELECT id FROM actors WHERE nameoriginal = $1 AND birthday = $2;`, jsnAC.NameOriginal, jsnAC.Birthday)
	if len(dbresp) == 0 {
		dbresp, err = dbm.Query(`
		INSERT INTO
			actors (
				imgsrc,
				name,
				nameoriginal,
				career,
				height,
				birthday,
				birthplace,
				total
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
		RETURNING id;`,
			jsnAC.Poster,
			jsnAC.Name,
			jsnAC.NameOriginal,
			jsnAC.Proffesion,
			fmt.Sprint(jsnAC.Height),
			jsnAC.Birthday,
			jsnAC.Birthplace,
			len(jsnAC.Films),
		)
		if err != nil {
			return 0, err
		}
	}

	return binary.BigEndian.Uint64(dbresp[0][0]), nil
}

func fillMoviesActors(movId, actId uint64) error {
	_, err := dbm.Query(`INSERT INTO movies_actors (movie_id, actor_id) VALUES ($1, $2);`, movId, actId)
	if err != nil {
		return err
	}

	return nil
}

func fillActorsGenres(actId uint64, jsnF *Movie) error {
	for _, genre := range jsnF.Genres {
		dbresp, err := dbm.Query(`SELECT COUNT(*) FROM actors_genres WHERE actor_id = $1 AND genre = $2;`, actId, genre.Genre)
		if err != nil {
			return err
		}

		if binary.BigEndian.Uint64(dbresp[0][0]) == 0 {
			_, err = dbm.Query(`INSERT INTO actors_genres (actor_id, genre) VALUES ($1, $2);`, actId, genre.Genre)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func fillMoviesMovies(movId uint64, jsnSI *Similars) error {
	for _, similar := range jsnSI.Items {
		dbresp, err := dbm.Query(`SELECT id FROM movies WHERE titleoriginal = $1;`, similar.NameOriginal)
		if err != nil {
			return err
		}

		if len(dbresp) != 0 {
			_, err = dbm.Query(`INSERT INTO movies_movies (movie_id, relation_id) VALUES ($1, $2), ($2, $1);`, movId, binary.BigEndian.Uint64(dbresp[0][0]))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func fillActorsActors(actId uint64) error {
	if actId < 15 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	var rel1, rel2 int
	for rel1 <= 0 {
		rel1 = rand.Intn(int(.75 * float64(actId)))
	}
	rel2 = rand.Intn(int(.25*float64(actId))) + int(.75*float64(actId))
	fmt.Println(actId, rel1, rel2)

	_, err := dbm.Query(`INSERT INTO actors_actors (actor_id, relation_id) VALUES ($1, $2), ($2, $1), ($1, $3), ($3, $1);`, actId, rel1, rel2)
	if err != nil {
		return err
	}

	return nil
}

func fillPlaylists() error {
	return nil
}
