package response

import "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"

type Meta struct {
	Message string              `json:"message" default:"true"`
	Info    *dto.PaginationInfo `json:"info"`
}
