package dto

import (
	cuc "github.com/k-shimizu04/ddd_by_example/usecase/companies"
)

type Response200Company struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Area        string `json:"area"`
	Genre       string `json:"genre"`
	ClosingDay  string `json:"closing_day"`
	Thumnail    string `json:"thumnail"`
}

func NewCompanyDTO(output *cuc.CompanyOutput) []Response200Company {
	response := []Response200Company{}
	for _, v := range output.List {
		response = append(response, Response200Company(v))
	}
	return response
}

type Response200CompanyIdDetails struct {
	Access         string   `json:"access"`
	OperatingHours string   `json:"operating_hours"`
	Contact        string   `json:"contact"`
	Thumnails      []string `json:"thumnails"`
}

type Response200CompanyId struct {
	Id          int32                        `json:"id"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Area        string                       `json:"area"`
	Genre       string                       `json:"genre"`
	ClosingDay  string                       `json:"closing_day"`
	Thumnail    string                       `json:"thumnail"`
	Details     *Response200CompanyIdDetails `json:"details"`
}

func NewCompanyIDDTO(output *cuc.CompanyIDOutput) *Response200CompanyId {
	if output == nil {
		return nil
	}
	response := Response200CompanyId{
		Id:          output.Id,
		Name:        output.Name,
		Description: output.Description,
		Area:        output.Area,
		Genre:       output.Genre,
		ClosingDay:  output.ClosingDay,
		Thumnail:    output.Thumnail,
		Details: &Response200CompanyIdDetails{
			Access:         output.Details.Access,
			OperatingHours: output.Details.OperatingHours,
			Contact:        output.Details.Contact,
			Thumnails:      output.Details.Thumnails,
		},
	}
	return &response
}
