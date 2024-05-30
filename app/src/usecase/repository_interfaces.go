package usecase

import (
	"github.com/k-shimizu04/ddd_by_example/domain"
)

type CompanyRepository interface {
	Find() (response []domain.Companies, err error)
}
