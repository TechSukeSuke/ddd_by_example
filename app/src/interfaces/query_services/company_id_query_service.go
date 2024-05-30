package query_services

import (
	"fmt"

	"github.com/k-shimizu04/ddd_by_example/derrors"
	"github.com/k-shimizu04/ddd_by_example/usecase"
	"gorm.io/gorm"
)

type CompanyIdQueryService struct {
	conn *gorm.DB
}

func NewCompanyIdQueryService(
	conn *gorm.DB,
) *CompanyIdQueryService {
	return &CompanyIdQueryService{
		conn: conn,
	}
}

func (c *CompanyIdQueryService) Execute(arg *usecase.CompanyIdQueryServiceARG) (dto *usecase.CompanyIdQueryServiceDTO, err error) {
	defer derrors.Wrap(&err, "CompanyIdQueryService.Execute(%#+v)", arg)

	sql := `
	SELECT
		c.id             AS id,
		c.name           AS name,
		c.description    AS description,
		c.area           AS area,
		c.genre          AS genre,
		c.closing_day    AS closing_day,
		c.thumnail       AS thumnail,
		cd.access        AS access,
		cd.opening_hours AS opening_hours,
		cd.contact       AS contact,
		cd.thumnails     AS thumnails
	FROM
		companies AS c
	INNER JOIN company_details AS cd
		on c.id = cd.company_id
	WHERE c.id = ?
	;`

	if err := c.conn.Raw(sql, arg.Id).Take(&dto).Error; err != nil {
		return nil, fmt.Errorf("message: %s, err: %w", err, derrors.OcurredAnQueryServieError)
	}

	return dto, nil
}
