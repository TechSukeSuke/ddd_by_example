package common_type

type (
	Company struct {
		Id          int32
		Name        string
		Description string
		Area        string
		Genre       string
		ClosingDay  string
		Thumnail    string
	}

	CompanyDetail struct {
		Access         string
		OperatingHours string
		Contact        string
		Thumnails      []string
	}
)
