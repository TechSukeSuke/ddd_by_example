package companies

import (
	"encoding/json"
	"fmt"

	"github.com/k-shimizu04/ddd_by_example/derrors"
	"github.com/k-shimizu04/ddd_by_example/usecase"
	commonType "github.com/k-shimizu04/ddd_by_example/usecase/companies/common"
)

type CompanyIDUsecase struct {
	CompanyIdQueryService usecase.CompanyIdQueryService
}

func NewCompanyIDUsecase(
	CompanyIdQueryService usecase.CompanyIdQueryService,
) *CompanyIDUsecase {
	return &CompanyIDUsecase{
		CompanyIdQueryService: CompanyIdQueryService,
	}
}

type CompanyIDInput struct {
	Id int32
}

type CompanyIDOutput struct {
	commonType.Company
	Details commonType.CompanyDetail
}

func (c *CompanyIDUsecase) Execute(input *CompanyIDInput) (output *CompanyIDOutput, err error) {
	defer derrors.Wrap(&err, "CompanyIDUsecase.Execute(%d)", input)

	Company, err := c.CompanyIdQueryService.Execute(&usecase.CompanyIdQueryServiceARG{
		Id: input.Id,
	})
	if err != nil {
		return nil, err
	}

	var unmarshaled struct {
		Thumnails []string
	}

	if err := json.Unmarshal(Company.Thumnails, &unmarshaled.Thumnails); err != nil {
		return nil, fmt.Errorf("message: %s, err: %w", err, derrors.JSONUnmarshalError)
	}

	output = &CompanyIDOutput{
		Company: commonType.Company{
			Id:          Company.Id,
			Name:        Company.Name,
			Description: Company.Description,
			Area:        Company.Area,
			Genre:       Company.Genre,
			ClosingDay:  Company.ClosingDay,
			Thumnail:    Company.Thumnail,
		},
		Details: commonType.CompanyDetail{
			Access:         Company.Area,
			OperatingHours: Company.OperatingHours,
			Contact:        Company.Contact,
			Thumnails:      unmarshaled.Thumnails,
		},
	}

	return output, nil
}
