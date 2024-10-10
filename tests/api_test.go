package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

const (
	URLEndPoint = "http://127.0.01:8080"
)

func TestAPIGet(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", URLEndPoint+"/api/get", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the status code
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Deserialize the JSON response
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Fatal(err)
	}

	// Check the response body
	expectedCode := 201
	expectedMsg := "Hello, World"
	if int(responseBody["code"].(float64)) != expectedCode ||
		responseBody["msg"].(string) != expectedMsg {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, map[string]interface{}{"code": expectedCode, "msg": expectedMsg})
	}

}

func TestAPISum(t *testing.T) {
	// Create a request body
	requestBody, err := json.Marshal(map[string]int{
		"a": 1,
		"b": 2,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create a request to pass to our handler
	req, err := http.NewRequest("POST", URLEndPoint+"/api/sum", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Deserialize the JSON response
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Fatal(err)
	}

	// Check the response body
	expectedCode := 201
	expectedMsg := "success"
	expectedData := map[string]int{"sum": 3}
	if int(responseBody["code"].(float64)) != expectedCode ||
		responseBody["msg"].(string) != expectedMsg ||
		!compareMaps(responseBody["data"].(map[string]interface{}), expectedData) {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, map[string]interface{}{"code": expectedCode, "msg": expectedMsg, "data": expectedData})
	}
}

func compareMaps(actual map[string]interface{}, expected map[string]int) bool {
	if len(actual) != len(expected) {
		return false
	}
	for key, expectedValue := range expected {
		actualValue, ok := actual[key]
		if !ok {
			return false
		}

		if int(actualValue.(float64)) != expectedValue {
			return false
		}
	}
	return true
}
