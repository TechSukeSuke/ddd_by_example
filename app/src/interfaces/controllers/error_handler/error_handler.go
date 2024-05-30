package error_handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/k-shimizu04/ddd_by_example/derrors"
)

type Error struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func ValidationErrorHandler(ctx *gin.Context, err error) {
	var response Error

	switch true {
	case errors.Is(err, derrors.BadRequest):
		response = Error{
			StatusCode: derrors.ToStatus(derrors.BadRequest),
			Message:    err.Error(),
		}
	case errors.Is(err, derrors.Unauthorized):
		response = Error{
			StatusCode: derrors.ToStatus(derrors.Unauthorized),
			Message:    err.Error(),
		}
	case errors.Is(err, derrors.Forbidden):
		response = Error{
			StatusCode: derrors.ToStatus(derrors.Forbidden),
			Message:    err.Error(),
		}
	case errors.Is(err, derrors.NotFound):
		response = Error{
			StatusCode: derrors.ToStatus(derrors.Unauthorized),
			Message:    err.Error(),
		}
	}

	ctx.JSON(response.StatusCode, response)
}

func ServerErrorHandler(ctx *gin.Context, err error) {
	var response Error

	switch true {
	case errors.Is(err, derrors.InternalServerError):
		response = Error{
			StatusCode: derrors.ToStatus(derrors.InternalServerError),
			Message:    err.Error(),
		}
	default:
		response = customErrorHandler(err)
	}

	ctx.JSON(response.StatusCode, response)
}

func customErrorHandler(err error) Error {
	var response Error

	switch true {
	case errors.Is(err, derrors.DBModuleInsertInvalid):
		response = Error{
			StatusCode: derrors.ToStatus(derrors.DBModuleInsertInvalid),
			Message:    err.Error(),
		}
	case errors.Is(err, derrors.OcurredAnQueryServieError):
		response = Error{
			StatusCode: derrors.ToStatus(derrors.OcurredAnQueryServieError),
			Message:    err.Error(),
		}
	case errors.Is(err, derrors.OccuredAnRepositoryError):
		response = Error{
			StatusCode: derrors.ToStatus(derrors.OccuredAnRepositoryError),
			Message:    err.Error(),
		}
	case errors.Is(err, derrors.JSONUnmarshalError):
		response = Error{
			StatusCode: derrors.ToStatus(derrors.JSONUnmarshalError),
			Message:    err.Error(),
		}
	}

	return response
}
