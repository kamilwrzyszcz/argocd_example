package main

import (
	"encoding/json"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Hello There!")
}

func main() {
	http.HandleFunc("/api/hello", hello)

	http.ListenAndServe(":8000", nil)
}
