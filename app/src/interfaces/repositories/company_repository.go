package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/k-shimizu04/ddd_by_example/derrors"
	"github.com/k-shimizu04/ddd_by_example/domain"
)

type CompanyRepository struct {
	conn *gorm.DB
}

func NewCompanyRepository(
	conn *gorm.DB,
) *CompanyRepository {
	return &CompanyRepository{
		conn: conn,
	}
}

func (c *CompanyRepository) Find() (response []domain.Companies, err error) {
	defer derrors.Wrap(&err, "CompanyRepository.Find()")
	if err := c.conn.Find(&response).Error; err != nil {
		return nil, fmt.Errorf("message: %s, err: %w", err, derrors.OccuredAnRepositoryError)
	}
	return response, nil
}
