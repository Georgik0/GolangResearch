package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "Processing /health POST request")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Handling /run GET request")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/run", runHandler)

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Info("test info")
	// logrus.Warn("test warn")
	logrus.WithFields(logrus.Fields{
		"adv_map_len_before_adding": 1,
		"number_of_added_types":     "test",
		"adv_map_len_new":           2,
	}).Warn("test warn")
	logrus.Error("test error")

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
