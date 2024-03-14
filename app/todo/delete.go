package todo

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

func (s service) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}

		err = s.toDoService.Delete(r.Context(), id)
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}
		utils.Respond(w, nil, http.StatusOK)
	}
}
