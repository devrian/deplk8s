package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"status":  http.StatusText(http.StatusOK),
			"message": "✅ Healthy Go App!",
		}

		json.NewEncoder(w).Encode(response)
		fmt.Println("✅ /health checked")
	})
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
