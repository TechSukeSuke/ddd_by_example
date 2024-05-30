package parser

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/k-shimizu04/ddd_by_example/derrors"
)

func ParamById(ctx *gin.Context, strId string) (int32, error) {
	param := ctx.Param(strId)
	if param == "" {
		return 0, fmt.Errorf("message: Invalid id due not to indicate. Error: %w", derrors.BadRequest)
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, fmt.Errorf("message: %s. Couldn't cast string to int. Error: %w", err.Error(), derrors.BadRequest)
	}

	return int32(id), nil
}
