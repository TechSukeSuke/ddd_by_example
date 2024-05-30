//go:build wireinject
// +build wireinject

package infrastructure

import (
	"github.com/google/wire"
	"github.com/k-shimizu04/ddd_by_example/config"
	ctrl "github.com/k-shimizu04/ddd_by_example/interfaces/controllers"
	cctrl "github.com/k-shimizu04/ddd_by_example/interfaces/controllers/companies"
	cm "github.com/k-shimizu04/ddd_by_example/interfaces/controllers/middleware"
	"github.com/k-shimizu04/ddd_by_example/interfaces/query_services"
	"github.com/k-shimizu04/ddd_by_example/interfaces/repositories"
	uc "github.com/k-shimizu04/ddd_by_example/usecase"
	cuc "github.com/k-shimizu04/ddd_by_example/usecase/companies"
	"gorm.io/gorm"
)

func InitRouting(
	configAPIServer *config.APIServer,
	d *gorm.DB,
) *Routing {
	wire.Build(
		// repository
		repositories.NewCompanyRepository,

		// query_service
		query_services.NewCompanyIdQueryService,

		// usecase companies
		cuc.NewCompanyUsecase,
		cuc.NewCompanyIDUsecase,

		// middleware cors
		cm.NewCorsMiddleware,

		// controller companies
		cctrl.NewCompanyController,
		cctrl.NewCompanyIDController,

		// all routing
		NewRouting,

		// bind infrastructure
		// wire.Bind(new(usecase.Redis), new(*redis.Redis)),
		// wire.Bind(new(usecase.Transaction), new(*repositories.Transaction)), // TODO
		// wire.Bind(new(usecase.Mail), new(*mail.Mail)),
		// wire.Bind(new(usecase.AWS), new(*aws.AWS)),

		// bind repository
		wire.Bind(new(uc.CompanyRepository), new(*repositories.CompanyRepository)),

		// bind query service
		wire.Bind(new(uc.CompanyIdQueryService), new(*query_services.CompanyIdQueryService)),

		// bind usecase
		wire.Bind(new(ctrl.CompanyUsecase), new(*cuc.CompanyUsecase)),
		wire.Bind(new(ctrl.CompanyIDUsecase), new(*cuc.CompanyIDUsecase)),
	)

	return &Routing{}
}
