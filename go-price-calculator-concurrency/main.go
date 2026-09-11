package main

import (
	"go-price-calculator-concurrency/prices"
)

func main() {
	taxRate := []float64{0, 0.2, 0.3, 0.4}
	myChannels := make([]chan bool, len(taxRate))
	for i, tax := range taxRate {
		myChannels[i] = make(chan bool)
		priceJob := prices.NewTaxIncludedPriceJob(tax)
		go priceJob.Process(myChannels[i])
	}
	for _, channel := range myChannels {
		<-channel
	}

}
