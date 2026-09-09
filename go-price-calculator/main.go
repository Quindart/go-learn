package main

import (
	"go-price-calculator/prices"
)

func main() {
	taxRate := []float64{0, 0.2, 0.3, 0.4}
	for _, tax := range taxRate {
		priceJob := prices.NewTaxIncludedPriceJob(tax)
		priceJob.Process()
	}
}
