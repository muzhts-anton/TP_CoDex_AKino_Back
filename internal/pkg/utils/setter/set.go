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
}

func SetHandlers(svs Services) {
	actRep := actrepository.InitActRep(svs.Act.Db)
	movRep := movrepository.InitMovRep(svs.Mov.Db)
	usrRep := usrrepository.InitUsrRep(svs.Usr.Db)
	colRep := colrepository.InitColRep(svs.Col.Db)
	genRep := genrepository.InitGenRep(svs.Gen.Db)
	annRep := annrepository.InitAnnRep(svs.Ann.Db)

	actUsc := actusecase.InitActUsc(actRep)
	movUsc := movusecase.InitMovUsc(movRep)
	usrUsc := usrusecase.InitUsrUsc(usrRep)
	colUsc := colusecase.InitColUsc(colRep)
	genUsc := genusecase.InitGenUsc(genRep)
	annUsc := annusecase.InitAnnUsc(annRep)

	actdelivery.SetActHandlers(svs.Act.Api, actUsc)
	movdelivery.SetMovHandlers(svs.Mov.Api, movUsc)
	usrdelivery.SetUsrHandlers(svs.Usr.Api, usrUsc)
	coldelivery.SetColHandlers(svs.Col.Api, colUsc)
	gendelivery.SetGenHandlers(svs.Gen.Api, genUsc)
	anndelivery.SetAnnHandlers(svs.Ann.Api, annUsc)
}
