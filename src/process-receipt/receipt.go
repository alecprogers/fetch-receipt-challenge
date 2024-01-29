package processreceipt

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"
)

type InputReceipt struct {
	Retailer     string             `json:"retailer"`
	PurchaseDate string             `json:"purchaseDate"`
	PurchaseTime string             `json:"purchaseTime"`
	Items        []InputReceiptItem `json:"items"`
	Total        string             `json:"total"`
}

type Receipt struct {
	Retailer         string
	PurchaseDatetime time.Time
	Items            []ReceiptItem
	Total            float64
}

// Parse InputReceipt struct to Receipt struct
func (r *InputReceipt) parseToReceipt() (Receipt, error) {
	// Convert PurchaseDate and PurchaseTime strings to time.Time
	datetime := r.PurchaseDate + " " + r.PurchaseTime
	parsedDatetime, err := time.Parse("2006-01-02 15:04", datetime)
	if err != nil {
		return Receipt{}, err
	}

	// Convert Total string to float64
	parsedTotal, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		return Receipt{}, err
	}

	// Convert InputReceiptItem slice to ReceiptItem slice
	parsedReceiptItems := make([]ReceiptItem, len(r.Items))
	for i, item := range r.Items {
		parsedReceiptItems[i], err = item.parseToReceiptItem()

		// If an error occurred, return it
		if err != nil {
			return Receipt{}, err
		}
	}

	// Otherwise, return fully parsed Receipt struct
	return Receipt{
			Retailer:         r.Retailer,
			PurchaseDatetime: parsedDatetime,
			Items:            parsedReceiptItems,
			Total:            parsedTotal,
		},
		nil
}

func (r *Receipt) calculatePoints() int {
	points := 0
	points += getRetailerNamePoints(r.Retailer)
	points += getRoundDollarTotalPoints(r.Total)
	points += getMultipleOf25Points(r.Total)
	points += getPointsPerTwoItems(len(r.Items))
	for _, item := range r.Items {
		points += getItemDescPoints(item.ShortDescription, item.Price)
	}
	points += getOddDayPoints(r.PurchaseDatetime)
	points += getTimeRangePoints(r.PurchaseDatetime)
	slog.Info("| " + fmt.Sprint(points) + " total points")
	return points
}

func (r *Receipt) print() {
	fmt.Printf("\nRetailer: %s\nPurchase datetime: %s\n",
		r.Retailer,
		r.PurchaseDatetime,
	)
	for i, item := range r.Items {
		fmt.Printf("  Item %d:\n", i+1)
		item.print()
	}

	fmt.Printf("Total: %.2f\n\n", r.Total)
}
