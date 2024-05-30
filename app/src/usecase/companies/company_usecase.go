package companies

import (
	"github.com/k-shimizu04/ddd_by_example/derrors"
	"github.com/k-shimizu04/ddd_by_example/usecase"
	commonType "github.com/k-shimizu04/ddd_by_example/usecase/companies/common"
)

type CompanyUsecase struct {
	CompanyRepository usecase.CompanyRepository
}

func NewCompanyUsecase(
	CompanyRepository usecase.CompanyRepository,
) *CompanyUsecase {
	return &CompanyUsecase{
		CompanyRepository: CompanyRepository,
	}
}

type CompanyInput struct{}

type CompanyOutput struct {
	List []commonType.Company
}

func (c *CompanyUsecase) Execute(input *CompanyInput) (output *CompanyOutput, err error) {
	defer derrors.Wrap(&err, "CompanyUsecase.Execute(%d)", input)

	companys, err := c.CompanyRepository.Find()
	if err != nil {
		return nil, err
	}

	var items []commonType.Company
	for _, v := range companys {
		items = append(items, commonType.Company{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Area:        v.Area,
			Genre:       v.Genre,
			ClosingDay:  v.ClosingDay,
			Thumnail:    v.Thumnail,
		})
	}
	output = &CompanyOutput{
		List: items,
	}

	return output, nil
}
