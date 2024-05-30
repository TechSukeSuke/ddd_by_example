package companies

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/k-shimizu04/ddd_by_example/interfaces/controllers"
	"github.com/k-shimizu04/ddd_by_example/interfaces/controllers/dto"
	"github.com/k-shimizu04/ddd_by_example/interfaces/controllers/error_handler"
	cuc "github.com/k-shimizu04/ddd_by_example/usecase/companies"
)

type CompanyController struct {
	usecase controllers.CompanyUsecase
}

func NewCompanyController(
	usecase controllers.CompanyUsecase,
) *CompanyController {
	return &CompanyController{
		usecase: usecase,
	}
}

func (c *CompanyController) Execute(ctx *gin.Context) {
	output, err := c.usecase.Execute(&cuc.CompanyInput{})
	if err != nil {
		error_handler.ServerErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.NewCompanyDTO(output))
}
