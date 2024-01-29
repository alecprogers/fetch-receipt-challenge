package processreceipt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

type Context context.Context
type Event events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

var InvalidReceiptRes = Response{
	StatusCode:      400,
	IsBase64Encoded: false,
	Body:            "{\"error\": \"invalid receipt\"}",
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

var SuccessRes = Response{
	StatusCode:      200,
	IsBase64Encoded: false,
	Headers: map[string]string{
		"Content-Type": "application/json",
	},
}

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (Response, error) {
	// Unmarshal JSON to InputReceipt struct
	bodyJsonData := []byte(event.Body)
	var inputReceipt InputReceipt

	// Enable DisallowUnknownFields to reject extra/missing properties
	decoder := json.NewDecoder(bytes.NewReader(bodyJsonData))
	decoder.DisallowUnknownFields()

	// Unmarshal JSON and check for errors
	if err := decoder.Decode(&inputReceipt); err != nil {
		return InvalidReceiptRes, err
	}

	// Parse InputReceipt to Receipt
	receipt, err := inputReceipt.parseToReceipt()
	if err != nil {
		return InvalidReceiptRes, err
	}

	// Print receipt
	// receipt.print()

	// Calculate point total for receipt
	points := receipt.calculatePoints()

	// Generate a random uuid
	id := uuid.New()
	slog.Info("| Assigned id " + fmt.Sprint(id))

	// Populate values to response body
	SuccessRes.Body = "{\"id\": \"" + fmt.Sprint(id) + "\", \"points\": " + fmt.Sprint(points) + "}"

	return SuccessRes, nil
}
