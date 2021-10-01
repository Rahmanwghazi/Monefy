package responses

import "github.com/Rahmanwghazi/Monefy/business/income"

type Income struct {
	Total       int    `json:"total"`
	Description string `json:"description"`
}

func FromDomain(domain income.IncomeDomain) Income {
	return Income{
		Description: domain.Description,
		Total:       domain.Total,
	}
}

func FromArrayDomain(domain []income.IncomeDomain) []Income {
	var income []Income
	for _, value := range domain {
		income = append(income, Income{
			Description: value.Description,
			Total:       value.Total,
		})
	}
	return income
}
