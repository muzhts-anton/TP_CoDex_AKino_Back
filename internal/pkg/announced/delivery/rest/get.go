package anndelivery

import (
	"codex/internal/pkg/domain"

	"net/http"
	"encoding/json"
)

func (handler *AnnouncedHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movs, err := handler.AnnouncedUsecase.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(movs)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
