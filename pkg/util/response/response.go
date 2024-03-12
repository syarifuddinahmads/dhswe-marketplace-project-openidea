package response

import (
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
)

type Meta struct {
	Success bool                `json:"success" default:"true"`
	Message string              `json:"message" default:"true"`
	Info    *dto.PaginationInfo `json:"info"`
}
