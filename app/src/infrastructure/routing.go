package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/k-shimizu04/ddd_by_example/config"
	cctrl "github.com/k-shimizu04/ddd_by_example/interfaces/controllers/companies"
	"github.com/k-shimizu04/ddd_by_example/interfaces/controllers/middleware"
)

type Routing struct {
	configAPIServer *config.APIServer
	Gin             *gin.Engine

	// middleware cors
	corsMiddleware *middleware.CorsMiddleware

	// controller companies
	companyController   *cctrl.CompanyController
	companyIDController *cctrl.CompanyIDController
}

func NewRouting(
	configAPIServer *config.APIServer,
	// middleware
	corsMiddleware *middleware.CorsMiddleware,
	// controller companies
	companyController *cctrl.CompanyController,
	companyIDController *cctrl.CompanyIDController,
) *Routing {
	g := gin.New()
	g.Use(gin.Recovery())

	r := &Routing{
		configAPIServer: configAPIServer,
		Gin:             g,
		// middleware
		corsMiddleware: corsMiddleware,
		// controller companies
		companyController:   companyController,
		companyIDController: companyIDController,
	}

	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	r.Gin.Use(r.corsMiddleware.Execute)

	root := r.Gin.Group("/api")
	{
		root.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "!!! Server is processing !!!")
		})
	}

	v1 := r.Gin.Group("/api/v1")
	{
		companies := v1.Group("/")
		{
			companies.GET("/companies", r.companyController.Execute)
			companies.GET("/companies/:id", r.companyIDController.Execute)
		}
	}
}

func (r *Routing) Run() error {
	return r.Gin.Run(":" + r.configAPIServer.APIPort)
}
