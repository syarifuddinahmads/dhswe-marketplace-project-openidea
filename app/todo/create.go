package todo

import (
	"net/http"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/todo/model"
	toDoService "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/todo/service"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

func (s service) Create() http.HandlerFunc {
	type request struct {
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
	}

	type response struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := utils.Decode(r, &req)
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}

		id, err := s.toDoService.Create(r.Context(), toDoService.CreateParams{
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
		})
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}
		utils.Respond(w, response{ID: id}, http.StatusOK)
	}
}
