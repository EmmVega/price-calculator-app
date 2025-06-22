package main

import (
	"calculator-app/filemanager"
	"calculator-app/prices"
	"fmt"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRate))
	for idx, taxRate := range taxRate {
		doneChans[idx] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[idx])
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
