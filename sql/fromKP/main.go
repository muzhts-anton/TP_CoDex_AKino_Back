package main

import (
	"log"

	"dbfill/database"
)

var dbm *database.DBManager = database.InitDatabase()

func main() {
	dbm.Connect()
	defer dbm.Disconnect()

	for _, filmId := range filmIds {
		// get info from kinopoisk
		jsnF, err := getFilm(filmId)
		if err != nil {
			log.Fatalln(err)
		}

		jsnBO, err := getBoxOffice(filmId)
		if err != nil {
			log.Fatalln(err)
		}

		jsnST, err := getStaff(filmId)
		if err != nil {
			log.Fatalln(err)
		}

		jsnV, err := getVideos(filmId)
		if err != nil {
			log.Fatalln(err)
		}

		jsnSI, err := getSimilars(filmId)
		if err != nil {
			log.Fatalln(err)
		}

		// fill local database
		movId, actorIds, err := fillMovie(jsnF, jsnBO, jsnV, jsnST)
		if err != nil {
			log.Fatalln(err)
		}

		err = fillGenres(jsnF)
		if err != nil {
			log.Fatalln(err)
		}

		err = fillMoviesGenres(jsnF, movId)
		if err != nil {
			log.Fatalln(err)
		}

		err = fillMoviesMovies(movId, jsnSI)
		if err != nil {
			log.Fatalln(err)
		}

		for _, actorId := range actorIds {
			jsnAC, err := getActor(actorId)
			if err != nil {
				log.Fatalln(err)
			}

			actId, err := fillActors(jsnAC)
			if err != nil {
				log.Fatalln(err)
			}

			err = fillMoviesActors(movId, actId)
			if err != nil {
				log.Fatalln(err)
			}

			err = fillActorsGenres(actId, jsnF)
			if err != nil {
				log.Fatalln(err)
			}

			err = fillActorsActors(actId)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	fillPlaylists()
}
