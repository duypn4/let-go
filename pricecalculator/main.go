package main

import (
	"fmt"
	"pricecalculator/filemanager"
	"pricecalculator/prices"
)

func main() {
	rates := []float64{0, 0.07, 0.1, 0.15}
	doneChan := make([]chan bool, len(rates))
	errChan := make([]chan error, len(rates))

	for index, rate := range rates {
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		// cmdManager := cmdmanager.New()
		doneChan[index] = make(chan bool)
		errChan[index] = make(chan error)
		job := prices.NewTaxIncludedPriceJob(fileManager, rate)
		go job.ProcessData(doneChan[index], errChan[index])
	}

	for index := range rates {
		select {
		case err := <-errChan[index]:
			fmt.Printf("job-%d: %v\n", index, err)
		case <-doneChan[index]:
			fmt.Printf("job-%d: done\n", index)
		}
	}
}
