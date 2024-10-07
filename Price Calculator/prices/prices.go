package prices

import "fmt"

type TaxIncPriceJob struct {
	TaxRate      float64
	InputPrices  []float64
	TaxIncPrices map[string]float64
}

func (job *TaxIncPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%0.3f", price)] = price * (1 + job.TaxRate)
		job.TaxIncPrices[fmt.Sprintf("%0.1f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

func New(taxRate float64) *TaxIncPriceJob {
	return &TaxIncPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
