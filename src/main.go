package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	processreceipt "github.com/alecprogers/fetch-receipt-challenge/src/process-receipt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gorilla/mux"
)

type Context context.Context
type Event events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

// Set port to listen on
const PORT = 8080

// Create map to store data in memory
var receipts = make(map[string]int)

// Create helper struct to parse result from lambda handler
type ProcessReceiptRes struct {
	ID     string `json:"id"`
	Points int    `json:"points"`
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Set endpoint handlers using the router
	router.HandleFunc("/receipts/process", processReceiptHandler).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", getPointsHandler).Methods("GET")

	// Start http server on port 8080
	fmt.Printf("Server listening on :%d...\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), router); err != nil {
		slog.Error(err.Error())
	}
}

func processReceiptHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse body from http request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Create sample context
	ctx := context.Background()

	// Map body to API Gateway event format
	apiGatewayEvent := events.APIGatewayProxyRequest{
		Body: string(body),
	}

	// Invoke Lambda handler function
	res, err := processreceipt.Handler(ctx, apiGatewayEvent)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, res.Body, http.StatusBadRequest)
		return
	}

	// Parse values from res
	var parsedRes ProcessReceiptRes
	if err := json.Unmarshal([]byte(res.Body), &parsedRes); err != nil {
		slog.Error(err.Error())
		http.Error(w, res.Body, http.StatusInternalServerError)
		return
	}

	// Add results to map
	receipts[parsedRes.ID] = parsedRes.Points

	// Print response to client
	fmt.Fprint(w, "{\"id\": \""+parsedRes.ID+"\"}")
}

func getPointsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get id from path
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := receipts[id]
	if !exists {
		slog.Warn(errors.New("receipt with id " + id + " does not exist in map").Error())
		http.Error(w, "{\"error\": \"receipt with id "+id+" not found\"}", http.StatusNotFound)
		return
	}

	slog.Info("| Retrieved point value " + fmt.Sprint(points) + " for id " + id)

	fmt.Fprintf(w, "{\"points\": %d}", points)
}
