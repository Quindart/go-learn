package prices

import (
	"fmt"
	"go-price-calculator/conversion"
	"go-price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64 `json:"tax_rate"`
	InputPrices       []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`// string: float
}

// parsing data from file
func (job *TaxIncludedPriceJob) LoadingData() {
	lines, _ := filemanager.ReadLines("prices.txt")
	prices, _ := conversion.StringToFloat(lines)
	job.InputPrices = prices
}

func (job TaxIncludedPriceJob) Process() {

	job.LoadingData()

	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	job.TaxIncludedPrices = result
	fmt.Println(result)
	filemanager.WriteJSON(fmt.Sprintf("tax_%.0f_result.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}
