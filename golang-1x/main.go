package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := Response{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/", Handle)
	http.ListenAndServe(":8080", nil)
}
