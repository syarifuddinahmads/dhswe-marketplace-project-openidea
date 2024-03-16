package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils/response"
)

func (s service) IndexProduct() http.HandlerFunc {
	return nil
}

func (s service) StoreProduct(w http.ResponseWriter, r *http.Request) {

	var payload dto.CreateProductRequest
	err := utils.Decode(r, &payload)
	if err != nil {
		response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(w)
		return
	}

	data, err := s.productService.StoreProduct(r.Context(), &payload)
	if err != nil {
		response.ErrorResponse(err).Send(w)
		return
	}

	// Send success response
	response.SuccessResponse(data).Send(w)
}

func (s service) ShowProduct() http.HandlerFunc {
	return nil
}

func (s service) EditProduct() http.HandlerFunc {
	return nil
}

func (s service) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the request URL
	vars := mux.Vars(r)
	productID := vars["id"]

	// Decode the payload
	var payload dto.UpdateProductRequest
	err := utils.Decode(r, &payload)
	if err != nil {
		response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(w)
		return
	}

	// Call the update service method
	err = s.productService.UpdateProduct(r.Context(), productID, &payload)
	if err != nil {
		response.ErrorResponse(err).Send(w)
		return
	}

	// Send success response
	response.SuccessResponse(nil).Send(w)
}

func (s service) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	// Call the update service method
	err := s.productService.DeleteProduct(r.Context(), productID)
	if err != nil {
		response.ErrorResponse(err).Send(w)
		return
	}

	// Send success response
	response.SuccessResponse(nil).Send(w)
}
