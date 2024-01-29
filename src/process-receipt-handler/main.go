package main

import (
	processreceipt "github.com/alecprogers/fetch-receipt-challenge/src/process-receipt"
	"github.com/aws/aws-lambda-go/lambda"
)

// This package is used to run the lambda in AWS
func main() {
	lambda.Start(processreceipt.Handler)
}
