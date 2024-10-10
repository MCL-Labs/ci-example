package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Call the function
	fmt.Println("Starting API Mock Server")

	go func() {
		http.HandleFunc("/api/dep/op", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
				return
			}

			// Seed the random number generator
			//rand.Seed(time.Now().UnixNano())

			// Randomly select "sub" or "add"
			op := "sub"
			// if rand.Intn(2) == 0 {
			// 	op = "add"
			// }

			response := map[string]interface{}{
				"code": 201,
				"msg":  "success",
				"data": map[string]interface{}{
					"op": op,
				},
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	select {}
}
