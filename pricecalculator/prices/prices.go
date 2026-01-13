package prices

import (
	"errors"
	"fmt"
	"pricecalculator/conversion"
	"pricecalculator/iomanager"
)

type TaxIncludedPriceJob struct {
	FileManager       iomanager.IoManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.FileManager.ReadLines()

	if err != nil {
		return errors.New("failed to load data")
	}

	job.InputPrices, err = conversion.StringsToFloats(lines)

	if err != nil {
		return errors.New("failed to convert data")
	}

	return nil
}

func (job *TaxIncludedPriceJob) ProcessData(doneChan chan bool, errChan chan error) {
	err := job.LoadData()

	if err != nil {
		errChan <- err
		return
	}

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.FileManager.WriteResult(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(ioManager iomanager.IoManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		FileManager:       ioManager,
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]string, 4),
	}
}
