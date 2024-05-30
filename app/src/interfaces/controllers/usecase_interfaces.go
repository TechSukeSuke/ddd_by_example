package controllers

import (
	cuc "github.com/k-shimizu04/ddd_by_example/usecase/companies"
)

type (
	CompanyUsecase interface {
		Execute(input *cuc.CompanyInput) (output *cuc.CompanyOutput, err error)
	}

	CompanyIDUsecase interface {
		Execute(input *cuc.CompanyIDInput) (output *cuc.CompanyIDOutput, err error)
	}
)
