package usecase

import (
	"encoding/json"

	"github.com/k-shimizu04/ddd_by_example/domain"
)

type CompanyIdQueryServiceARG struct {
	Id int32
}

type CompanyIdQueryServiceDTO struct {
	domain.Companies

	Access         string
	OperatingHours string
	Contact        string
	Thumnails      json.RawMessage
}

type CompanyIdQueryService interface {
	Execute(arg *CompanyIdQueryServiceARG) (dto *CompanyIdQueryServiceDTO, err error)
}
