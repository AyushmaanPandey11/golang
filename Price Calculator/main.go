package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxInclPrices := make([]float64, len(prices))
		for priceIdx, price := range prices {
			taxInclPrices[priceIdx] = price * (1 + taxRate)
		}
		result[taxRate] = taxInclPrices
	}

	fmt.Println(result)
}
