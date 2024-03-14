package auth

import (
	"fmt"
	"net/http"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

func (s service) Register() http.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type response struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := utils.Decode(r, &req)
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}

		id, err := s.userService.RegisterUser(r.Context(), dto.CreateUserParams{
			Name:     req.Name,
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			fmt.Println("Error Register")
			utils.Respond(w, err, 0)
			return
		}
		utils.Respond(w, response{ID: id}, http.StatusOK)
	}
}

func (s service) Login() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type response struct {
		ID bool `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := utils.Decode(r, &req)
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}

		id, err := s.userService.LoginUser(r.Context(), dto.AuthLoginRequest{
			Username: req.Username,
			Password: req.Password,
		})

		if err != nil {
			fmt.Println("Error Register")
			utils.Respond(w, err, 0)
			return
		}
		utils.Respond(w, response{ID: id}, http.StatusOK)
	}
}
