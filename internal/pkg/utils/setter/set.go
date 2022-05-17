package setter

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/utils/config"
	"codex/internal/pkg/utils/log"

	usrdelivery "codex/internal/pkg/user/delivery/rest"
	"codex/internal/pkg/user/repository"
	"codex/internal/pkg/user/usecase"

	coldelivery "codex/internal/pkg/collections/delivery/rest"
	colrepository "codex/internal/pkg/collections/repository"
	colusecase "codex/internal/pkg/collections/usecase"

	movdelivery "codex/internal/pkg/movie/delivery/rest"
	movrepository "codex/internal/pkg/movie/repository"
	"codex/internal/pkg/movie/usecase"

	actdelivery "codex/internal/pkg/actor/delivery/rest"
	actrepository "codex/internal/pkg/actor/repository"
	"codex/internal/pkg/actor/usecase"

	gendelivery "codex/internal/pkg/genres/delivery/rest"
	"codex/internal/pkg/genres/repository"
	"codex/internal/pkg/genres/usecase"

	"codex/internal/pkg/announced/delivery/rest"
	"codex/internal/pkg/announced/repository"
	"codex/internal/pkg/announced/usecase"

	"codex/internal/pkg/search/delivery/rest"
	"codex/internal/pkg/search/repository"
	"codex/internal/pkg/search/usecase"

	"codex/internal/pkg/playlist/delivery"
	"codex/internal/pkg/playlist/repository"
	"codex/internal/pkg/playlist/usecase"

	ratmcs "codex/internal/pkg/rating/delivery/grpc"
	"codex/internal/pkg/rating/delivery/rest"

	autmcs "codex/internal/pkg/authorization/delivery/grpc"
	"codex/internal/pkg/authorization/delivery/rest"

	commcs "codex/internal/pkg/comment/delivery/grpc"
	"codex/internal/pkg/comment/delivery/rest"

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
	Ser Data
	Pla Data

	Rat Data
	Aut Data
	Com Data
}

func setAutMcs() autmcs.AutherClient {
	autconn, err := grpc.Dial(":"+config.DevConfigStore.Mcs.Auth.Port, grpc.WithInsecure())
	if err != nil {
		log.Warn("{setAutMcs} mcs Dial")
	}

	return autmcs.NewAutherClient(autconn)
}

func setComMcs() commcs.PosterClient {
	autconn, err := grpc.Dial(":"+config.DevConfigStore.Mcs.Comment.Port, grpc.WithInsecure())
	if err != nil {
		log.Warn("{setComMcs} mcs Dial")
	}

	return commcs.NewPosterClient(autconn)
}

func setRatMcs() ratmcs.PosterClient {
	autconn, err := grpc.Dial(":"+config.DevConfigStore.Mcs.Rating.Port, grpc.WithInsecure())
	if err != nil {
		log.Warn("{setRatMcs} mcs Dial")
	}

	return ratmcs.NewPosterClient(autconn)
}

func SetHandlers(svs Services) {
	actRep := actrepository.InitActRep(svs.Act.Db)
	movRep := movrepository.InitMovRep(svs.Mov.Db)
	usrRep := usrrepository.InitUsrRep(svs.Usr.Db)
	colRep := colrepository.InitColRep(svs.Col.Db)
	genRep := genrepository.InitGenRep(svs.Gen.Db)
	annRep := annrepository.InitAnnRep(svs.Ann.Db)
	serRep := serrepository.InitSerRep(svs.Ann.Db)
	plaRep := plarepository.InitPlaRep(svs.Pla.Db)

	actUsc := actusecase.InitActUsc(actRep)
	movUsc := movusecase.InitMovUsc(movRep)
	usrUsc := usrusecase.InitUsrUsc(usrRep)
	colUsc := colusecase.InitColUsc(colRep)
	genUsc := genusecase.InitGenUsc(genRep)
	annUsc := annusecase.InitAnnUsc(annRep)
	serUsc := serusecase.InitSerUsc(serRep)
	plaUsc := plausecase.InitPlaUsc(plaRep)

	actdelivery.SetActHandlers(svs.Act.Api, actUsc)
	movdelivery.SetMovHandlers(svs.Mov.Api, movUsc)
	usrdelivery.SetUsrHandlers(svs.Usr.Api, usrUsc)
	coldelivery.SetColHandlers(svs.Col.Api, colUsc)
	gendelivery.SetGenHandlers(svs.Gen.Api, genUsc)
	anndelivery.SetAnnHandlers(svs.Ann.Api, annUsc)
	serdelivery.SetSerHandlers(svs.Ann.Api, serUsc)
	pladelivery.SetPlaHandlers(svs.Pla.Api, plaUsc)

	ratdelivery.SetRatHandlers(svs.Act.Api, setRatMcs())
	autdelivery.SetAutHandlers(svs.Aut.Api, setAutMcs())
	comdelivery.SetComHandlers(svs.Com.Api, setComMcs())
}
