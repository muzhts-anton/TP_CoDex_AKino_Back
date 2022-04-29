package setter

import (
	"codex/internal/pkg/database"

	usrdelivery "codex/internal/pkg/user/delivery"
	usrrepository "codex/internal/pkg/user/repository"
	usrusecase "codex/internal/pkg/user/usecase"

	coldelivery "codex/internal/pkg/collections/delivery"
	colrepository "codex/internal/pkg/collections/repository"
	colusecase "codex/internal/pkg/collections/usecase"

	movdelivery "codex/internal/pkg/movie/delivery"
	movrepository "codex/internal/pkg/movie/repository"
	movusecase "codex/internal/pkg/movie/usecase"

	actdelivery "codex/internal/pkg/actor/delivery"
	actrepository "codex/internal/pkg/actor/repository"
	actusecase "codex/internal/pkg/actor/usecase"

	gendelivery "codex/internal/pkg/genres/delivery"
	genrepository "codex/internal/pkg/genres/repository"
	genusecase "codex/internal/pkg/genres/usecase"

	anndelivery "codex/internal/pkg/announced/delivery"
	annrepository "codex/internal/pkg/announced/repository"
	annusecase "codex/internal/pkg/announced/usecase"

	comdelivery "codex/internal/pkg/comment/delivery"
	comrepository "codex/internal/pkg/comment/repository"
	comusecase "codex/internal/pkg/comment/usecase"

	ratdelivery "codex/internal/pkg/rating/delivery"
	ratrepository "codex/internal/pkg/rating/repository"
	ratusecase "codex/internal/pkg/rating/usecase"

	autdelivery "codex/internal/pkg/authorization/delivery"
	autrepository "codex/internal/pkg/authorization/repository"
	autusecase "codex/internal/pkg/authorization/usecase"

	pladelivery "codex/internal/pkg/playlist/delivery"
	plarepository "codex/internal/pkg/playlist/repository"
	plausecase "codex/internal/pkg/playlist/usecase"

	"github.com/gorilla/mux"
)

type Data struct {
	Db  *database.DBManager
	Api *mux.Router
}

type Services struct {
	Act Data
	Mov Data
	Usr Data
	Col Data
	Gen Data
	Ann Data
	Com Data
	Rat Data
	Aut Data
	Pla Data
}

func SetHandlers(svs Services) {
	actRep := actrepository.InitActRep(svs.Act.Db)
	movRep := movrepository.InitMovRep(svs.Mov.Db)
	usrRep := usrrepository.InitUsrRep(svs.Usr.Db)
	colRep := colrepository.InitColRep(svs.Col.Db)
	genRep := genrepository.InitGenRep(svs.Gen.Db)
	annRep := annrepository.InitAnnRep(svs.Ann.Db)
	comRep := comrepository.InitComRep(svs.Com.Db)
	ratRep := ratrepository.InitRatRep(svs.Rat.Db)
	autRep := autrepository.InitAutRep(svs.Aut.Db)
	plaRep := plarepository.InitPlaRep(svs.Pla.Db)

	actUsc := actusecase.InitActUsc(actRep)
	movUsc := movusecase.InitMovUsc(movRep)
	usrUsc := usrusecase.InitUsrUsc(usrRep)
	colUsc := colusecase.InitColUsc(colRep)
	genUsc := genusecase.InitGenUsc(genRep)
	annUsc := annusecase.InitAnnUsc(annRep)
	comUsc := comusecase.InitComUsc(comRep)
	ratUsc := ratusecase.InitRatUsc(ratRep)
	autUsc := autusecase.InitAutUsc(autRep)
	plaUsc := plausecase.InitPlaUsc(plaRep)

	actdelivery.SetActHandlers(svs.Act.Api, actUsc)
	movdelivery.SetMovHandlers(svs.Mov.Api, movUsc)
	usrdelivery.SetUsrHandlers(svs.Usr.Api, usrUsc)
	coldelivery.SetColHandlers(svs.Col.Api, colUsc)
	gendelivery.SetGenHandlers(svs.Gen.Api, genUsc)
	anndelivery.SetAnnHandlers(svs.Ann.Api, annUsc)
	comdelivery.SetComHandlers(svs.Ann.Api, comUsc)
	ratdelivery.SetRatHandlers(svs.Ann.Api, ratUsc)
	autdelivery.SetAutHandlers(svs.Aut.Api, autUsc)
	pladelivery.SetPlaHandlers(svs.Pla.Api, plaUsc)
}
