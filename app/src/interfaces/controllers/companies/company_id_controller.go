package companies

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/k-shimizu04/ddd_by_example/interfaces/controllers"
	"github.com/k-shimizu04/ddd_by_example/interfaces/controllers/dto"
	"github.com/k-shimizu04/ddd_by_example/interfaces/controllers/error_handler"
	"github.com/k-shimizu04/ddd_by_example/interfaces/controllers/parser"
	cuc "github.com/k-shimizu04/ddd_by_example/usecase/companies"
)

type CompanyIDController struct {
	usecase controllers.CompanyIDUsecase
}

func NewCompanyIDController(
	usecase controllers.CompanyIDUsecase,
) *CompanyIDController {
	return &CompanyIDController{
		usecase: usecase,
	}
}

func (c *CompanyIDController) Execute(ctx *gin.Context) {
	id, err := parser.ParamById(ctx, "id")
	if err != nil {
		error_handler.ValidationErrorHandler(ctx, err)
		return
	}

	output, err := c.usecase.Execute(&cuc.CompanyIDInput{
		Id: id,
	})
	if err != nil {
		error_handler.ServerErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewCompanyIDDTO(output))
}
