package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"bitizen.com/ci-example/utils"
)

const (
	MockURLEndPoint = "http://127.0.0.1:8081"
)

func getOp(url string) (string, error) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", url+"/api/dep/op", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Deserialize the JSON response
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return "", err
	}

	// Check the response body
	if int(responseBody["code"].(float64)) != 201 ||
		responseBody["msg"].(string) != "success" {
		return "", fmt.Errorf("handler returned unexpected body: got %v", responseBody)
	}

	return responseBody["data"].(map[string]interface{})["op"].(string), nil
}

func main() {
	// Call the function
	fmt.Println("Starting API Server")

	go func() {
		http.HandleFunc("/api/get", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
				return
			}

			response := map[string]interface{}{
				"code": 201,
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
				"code": 201,
				"msg":  "success",
				"data": map[string]int{
					"sum": sum,
				},
			}

			// Set the Content-Type header and encode the response as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		})

		http.HandleFunc("/api/calc", func(w http.ResponseWriter, r *http.Request) {
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

			op, err := getOp(MockURLEndPoint)
			if err != nil {
				http.Error(w, "Failed to get operation", http.StatusInternalServerError)
				return
			}

			if op == "add" {
				// Calculate the sum
				sum := utils.Add(requestBody.ParamA, requestBody.ParamB)

				// Create the response
				response := map[string]interface{}{
					"code": 201,
					"msg":  "success",
					"data": map[string]int{
						"result": sum,
					},
				}

				// Set the Content-Type header and encode the response as JSON
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
				return
			} else if op == "sub" {
				// Calculate the sum
				sub := utils.Sub(requestBody.ParamA, requestBody.ParamB)

				sub = 1

				// Create the response
				response := map[string]interface{}{
					"code": 201,
					"msg":  "success",
					"data": map[string]int{
						"result": sub,
					},
				}

				// Set the Content-Type header and encode the response as JSON
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
				return
			} else {
				http.Error(w, "Invalid operation", http.StatusInternalServerError)
				return
			}

		})

		log.Println("Starting API server on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// block forever
	select {}
}
