package processreceipt

import (
	"fmt"
	"strconv"
)

type InputReceiptItem struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ReceiptItem struct {
	ShortDescription string
	Price            float64
}

// Parse InputReceiptItem struct to ReceiptItem struct
func (ri *InputReceiptItem) parseToReceiptItem() (ReceiptItem, error) {
	// Convert Price string to float64
	fPrice, err := strconv.ParseFloat(ri.Price, 64)

	// If an error occurred, return it
	if err != nil {
		return ReceiptItem{}, err
	}

	// Otherwise, return parsed ReceiptItem struct
	return ReceiptItem{ShortDescription: ri.ShortDescription, Price: fPrice}, nil
}

func (ri *ReceiptItem) print() {
	fmt.Printf("    Description: %s\n    Price: %.2f\n", ri.ShortDescription, ri.Price)
}
