package setter

import (
	"codex/internal/pkg/database"

	"codex/internal/pkg/user/delivery"
	"codex/internal/pkg/user/repository"
	"codex/internal/pkg/user/usecase"

	"codex/internal/pkg/collections/delivery"
	"codex/internal/pkg/collections/repository"
	"codex/internal/pkg/collections/usecase"

	"codex/internal/pkg/movie/delivery"
	"codex/internal/pkg/movie/repository"
	"codex/internal/pkg/movie/usecase"

	"codex/internal/pkg/actor/delivery"
	"codex/internal/pkg/actor/repository"
	"codex/internal/pkg/actor/usecase"

	"codex/internal/pkg/genres/delivery"
	"codex/internal/pkg/genres/repository"
	"codex/internal/pkg/genres/usecase"

	"codex/internal/pkg/announced/delivery"
	"codex/internal/pkg/announced/repository"
	"codex/internal/pkg/announced/usecase"

	"codex/internal/pkg/comment/delivery"
	"codex/internal/pkg/comment/repository"
	"codex/internal/pkg/comment/usecase"

	"codex/internal/pkg/rating/delivery"
	"codex/internal/pkg/rating/repository"
	"codex/internal/pkg/rating/usecase"

	"codex/internal/pkg/authorization/delivery"
	"codex/internal/pkg/authorization/repository"
	"codex/internal/pkg/authorization/usecase"

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

	actUsc := actusecase.InitActUsc(actRep)
	movUsc := movusecase.InitMovUsc(movRep)
	usrUsc := usrusecase.InitUsrUsc(usrRep)
	colUsc := colusecase.InitColUsc(colRep)
	genUsc := genusecase.InitGenUsc(genRep)
	annUsc := annusecase.InitAnnUsc(annRep)
	comUsc := comusecase.InitComUsc(comRep)
	ratUsc := ratusecase.InitRatUsc(ratRep)
	autUsc := autusecase.InitAutUsc(autRep)

	actdelivery.SetActHandlers(svs.Act.Api, actUsc)
	movdelivery.SetMovHandlers(svs.Mov.Api, movUsc)
	usrdelivery.SetUsrHandlers(svs.Usr.Api, usrUsc)
	coldelivery.SetColHandlers(svs.Col.Api, colUsc)
	gendelivery.SetGenHandlers(svs.Gen.Api, genUsc)
	anndelivery.SetAnnHandlers(svs.Ann.Api, annUsc)
	comdelivery.SetComHandlers(svs.Ann.Api, comUsc)
	ratdelivery.SetRatHandlers(svs.Ann.Api, ratUsc)
	autdelivery.SetAutHandlers(svs.Aut.Api, autUsc)
}
