package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Item represents the request payload matching the FastAPI Pydantic schema.
type Item struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Price       float64  `json:"price"`
	Tax         *float64 `json:"tax,omitempty"`
}

// ItemResponse represents the structured response returned by the backend.
type ItemResponse struct {
	Message    string  `json:"message"`
	Item       Item    `json:"item"`
	TotalPrice float64 `json:"total_price"`
}

const apiURL = "http://127.0.0.1:8000/items/"

// CreateItem sends a POST request to the FastAPI /items/ endpoint.
func CreateItem(item Item) (*ItemResponse, error) {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("error marshaling item JSON: %w", err)
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error sending POST request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var itemResp ItemResponse
	if err := json.Unmarshal(bodyBytes, &itemResp); err != nil {
		return nil, fmt.Errorf("error decoding response JSON: %w", err)
	}

	return &itemResp, nil
}

func main() {
	description := "High performance laptop"
	tax := 120.0

	newItem := Item{
		Name:        "Gaming Laptop",
		Description: &description,
		Price:       1200.0,
		Tax:         &tax,
	}

	fmt.Println("Sending POST request to FastAPI backend...")
	response, err := CreateItem(newItem)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("API Response:")
	fmt.Printf("Message:     %s\n", response.Message)
	fmt.Printf("Item Name:   %s\n", response.Item.Name)
	fmt.Printf("Price:       $%.2f\n", response.Item.Price)
	fmt.Printf("Total Price: $%.2f\n", response.TotalPrice)
}
