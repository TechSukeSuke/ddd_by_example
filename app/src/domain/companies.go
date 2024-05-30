package domain

type Companies struct {
	Id          int32 `gorm:"primaryKey"`
	Name        string
	Description string
	Area        string
	Genre       string
	ClosingDay  string
	Thumnail    string
}

func (c *Companies) TableName() string {
	return "companies"
}
