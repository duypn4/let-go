package main

import (
	"fmt"
	"pricecalculator/filemanager"
	"pricecalculator/prices"
)

func main() {
	rates := []float64{0, 0.07, 0.1, 0.15}
	for _, rate := range rates {
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		// cmdManager := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(fileManager, rate)
		job.ProcessData()
	}
}
