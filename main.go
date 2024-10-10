package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"landui.com/ci-example/utils"
)

func main() {
	// Call the function
	fmt.Println("Hello, World!")

	go func() {
		http.HandleFunc("/api/get", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
				return
			}

			response := map[string]interface{}{
				"code": 200,
				"msg":  "Hello, World",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		http.HandleFunc("/api/sum", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
				return
			}

			// Define a struct to parse the JSON body
			var requestBody struct {
				ParamA int `json:"a"`
				ParamB int `json:"b"`
			}

			// Parse the JSON body
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				http.Error(w, "Invalid JSON body", http.StatusBadRequest)
				return
			}

			// Calculate the sum
			sum := utils.Add(requestBody.ParamA, requestBody.ParamB)

			// Create the response
			response := map[string]interface{}{
				"code": 200,
				"msg":  "success",
				"data": map[string]int{
					"sum": sum,
				},
			}

			// Set the Content-Type header and encode the response as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		})

		log.Println("Starting API server on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// block forever
	select {}
}
