package prices

import (
	"fmt"

	filemanager "example.com/m/fileManager"
)

type TaxIncPriceJob struct {
	TaxRate      float64            `json:"tax_rate"`
	InputPrices  []float64          `json:"input_prices"`
	TaxIncPrices map[string]float64 `json:"tax_included_prices"`
}

func (job *TaxIncPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%0.3f", price)] = price * (1 + job.TaxRate)
	}
	job.TaxIncPrices = result
	filemanager.WriteJSON(fmt.Sprintf("data_%.0f.json", job.TaxRate*100), job)
}

func New(taxRate float64) *TaxIncPriceJob {
	return &TaxIncPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
