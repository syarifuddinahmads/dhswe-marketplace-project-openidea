package todo

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/model"
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

		id, err := s.toDoService.Create(r.Context(), CreateParams{
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

func (s service) Get() http.HandlerFunc {
	type response struct {
		ID          int          `json:"id"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
		CreatedOn   time.Time    `json:"created_on"`
		UpdatedOn   *time.Time   `json:"updated_on,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.Respond(w, utils.ErrArgument{
				Wrapped: errors.New("valid id must provide in path"),
			}, 0)
			return
		}

		getResponse, err := s.toDoService.Get(r.Context(), id)
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}
		utils.Respond(w, response{
			ID:          getResponse.ID,
			Name:        getResponse.Name,
			Description: getResponse.Description,
			Status:      getResponse.Status,
			CreatedOn:   getResponse.CreatedOn,
			UpdatedOn:   getResponse.UpdatedOn,
		}, http.StatusOK)
	}
}

func (s service) GetAll() http.HandlerFunc {
	type response struct {
		ID          int          `json:"id"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
		CreatedOn   time.Time    `json:"created_on"`
		UpdatedOn   *time.Time   `json:"updated_on,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := s.toDoService.GetAll(r.Context())
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}

		var todoResponses []response
		for _, todo := range todos {
			todoResponses = append(todoResponses, response{
				ID:          todo.ID,
				Name:        todo.Name,
				Description: todo.Description,
				Status:      todo.Status,
				CreatedOn:   todo.CreatedOn,
				UpdatedOn:   todo.UpdatedOn,
			})
		}

		utils.Respond(w, todoResponses, http.StatusOK)
	}
}

func (s service) Update() http.HandlerFunc {
	type request struct {
		Name        *string       `json:"name"`
		Description *string       `json:"description"`
		Status      *model.Status `json:"status"`
	}

	type response struct {
		ID          int          `json:"id"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
		CreatedOn   time.Time    `json:"created_on"`
		UpdatedOn   *time.Time   `json:"updated_on,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.Respond(w, utils.ErrArgument{
				Wrapped: errors.New("valid id must provide in path"),
			}, 0)
			return
		}

		req := request{}
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err = utils.Decode(r, &req)
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}

		err = s.toDoService.Update(r.Context(), UpdateParams{
			ID:          id,
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
