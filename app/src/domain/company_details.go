package domain

type CompanyDetails struct {
	Id             int32 `gorm:"primaryKey"`
	CompanyId      int32
	Access         string
	OperatingHours string
	Contact        string
	Thumnails      []string
}

func (cd *CompanyDetails) TableName() string {
	return "company_details"
}
