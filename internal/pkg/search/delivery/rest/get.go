package serdelivery

import (
	"codex/internal/pkg/domain"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (handler *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	tag := mux.Vars(r)["genre"]
	if tag == "" {
		http.Error(w, domain.Err.ErrObj.ParseId.Error(), http.StatusBadRequest)
		return
	}

	sr, err := handler.SearchUsecase.Search(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(sr)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
