package v1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"kodix/src/internal/auto/entity"
	"kodix/src/internal/facade"
	"kodix/src/internal/http/api/response"
	"kodix/src/internal/models"
	"net/http"
	"strconv"
)

func GetAutoById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		autoId := mux.Vars(request)["id"]
		id, err := strconv.Atoi(autoId)
		if err != nil {
			response.Send(nil, response.ErrBadParams, writer)
		}

		result, err := facade.Services().Auto().GetAutoById(request.Context(), id)
		if err != nil {
			response.Send(nil, response.ErrServerError, writer)
		}

		response.Send(result, response.ErrEmpty, writer)
	}
}

func GetAllAuto() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		result, err := facade.Services().Auto().GetAllAuto(request.Context())
		if err != nil {
			response.Send(nil, response.ErrServerError, writer)
		}

		response.Send(result, response.ErrEmpty, writer)
	}
}

func CreateUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var attribute models.AutoAttributes

		err := json.NewDecoder(request.Body).Decode(&attribute)
		if err != nil {
			response.Send(nil, response.ErrServerError, writer)
		}

		model := entity.NewAuto(
			attribute.Brand,
			attribute.Model,
			attribute.Mileage,
		)
		if err := model.SetStatus(attribute.Status); err != nil {
			response.Send(nil, response.ErrBadParams, writer)
		}
		if err := model.SetPrice(int(attribute.Price)); err != nil {
			response.Send(nil, response.ErrBadParams, writer)
		}

		if err := facade.Services().Auto().CreateAuto(request.Context(), *model); err != nil {
			response.Send(nil, response.ErrServerError, writer)
		}
		response.Send(nil, response.ErrEmpty, writer)
	}
}

func UpdateAutoById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var attribute models.AutoAttributes

		id, err := strconv.Atoi(mux.Vars(request)["id"])
		if err != nil {
			response.Send(nil, response.ErrBadParams, writer)
		}

		if err := json.NewDecoder(request.Body).Decode(&attribute); err != nil {
			response.Send(nil, response.ErrServerError, writer)
		}

		model := entity.NewAuto(
			attribute.Brand,
			attribute.Model,
			attribute.Mileage,
		)
		if err := model.SetStatus(attribute.Status); err != nil {
			response.Send(nil, response.ErrBadParams, writer)
		}
		if err := model.SetPrice(int(attribute.Price)); err != nil {
			response.Send(nil, response.ErrBadParams, writer)
		}
		if err := facade.Services().Auto().UpdateAuto(request.Context(), *model, id); err != nil {
			response.Send(nil, response.ErrServerError, writer)
		}
		response.Send(nil, response.ErrEmpty, writer)
	}
}

func DeleteAutoById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		autoId := mux.Vars(request)["id"]
		id, err := strconv.Atoi(autoId)
		if err != nil {
			response.Send(nil, response.ErrBadParams, writer)
		}

		if err := facade.Services().Auto().DeleteAuto(request.Context(), id); err != nil {
			response.Send(nil, response.ErrServerError, writer)
		}

		response.Send(nil, response.ErrEmpty, writer)
	}
}
