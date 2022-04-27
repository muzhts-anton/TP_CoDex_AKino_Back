package setter

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/utils/log"

	"codex/internal/pkg/user/delivery/rest"
	"codex/internal/pkg/user/repository"
	"codex/internal/pkg/user/usecase"

	"codex/internal/pkg/collections/delivery/rest"
	"codex/internal/pkg/collections/repository"
	"codex/internal/pkg/collections/usecase"

	"codex/internal/pkg/movie/delivery/rest"
	"codex/internal/pkg/movie/repository"
	"codex/internal/pkg/movie/usecase"

	"codex/internal/pkg/actor/delivery/rest"
	"codex/internal/pkg/actor/repository"
	"codex/internal/pkg/actor/usecase"

	"codex/internal/pkg/genres/delivery/rest"
	"codex/internal/pkg/genres/repository"
	"codex/internal/pkg/genres/usecase"

	"codex/internal/pkg/announced/delivery/rest"
	"codex/internal/pkg/announced/repository"
	"codex/internal/pkg/announced/usecase"

	"codex/internal/pkg/comment/delivery/rest"
	"codex/internal/pkg/comment/repository"
	"codex/internal/pkg/comment/usecase"

	"codex/internal/pkg/rating/delivery/rest"
	"codex/internal/pkg/rating/repository"
	"codex/internal/pkg/rating/usecase"

	autmcs "codex/internal/pkg/authorization/delivery/grpc"
	"codex/internal/pkg/authorization/delivery/rest"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
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

func setAutMcs() autmcs.AutherClient {
	autconn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Warn("{setAutMcs} mcs Dial")
	}

	return autmcs.NewAutherClient(autconn)
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

	actUsc := actusecase.InitActUsc(actRep)
	movUsc := movusecase.InitMovUsc(movRep)
	usrUsc := usrusecase.InitUsrUsc(usrRep)
	colUsc := colusecase.InitColUsc(colRep)
	genUsc := genusecase.InitGenUsc(genRep)
	annUsc := annusecase.InitAnnUsc(annRep)
	comUsc := comusecase.InitComUsc(comRep)
	ratUsc := ratusecase.InitRatUsc(ratRep)

	actdelivery.SetActHandlers(svs.Act.Api, actUsc)
	movdelivery.SetMovHandlers(svs.Mov.Api, movUsc)
	usrdelivery.SetUsrHandlers(svs.Usr.Api, usrUsc)
	coldelivery.SetColHandlers(svs.Col.Api, colUsc)
	gendelivery.SetGenHandlers(svs.Gen.Api, genUsc)
	anndelivery.SetAnnHandlers(svs.Ann.Api, annUsc)
	comdelivery.SetComHandlers(svs.Com.Api, comUsc)
	ratdelivery.SetRatHandlers(svs.Rat.Api, ratUsc)

	autdelivery.SetAutHandlers(svs.Aut.Api, setAutMcs())
}
