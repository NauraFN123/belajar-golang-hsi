package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	ErrEmailEmpty = errors.New("email tidak boleh kosong")
	ErrAgeInvalid = errors.New("umur harus 18 tahun atau lebih")
)

type Response struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: Method=%s, Path=%s, Query=%s", r.Method, r.URL.Path, r.URL.RawQuery)

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(Response{Error: "Method not allowed"})
		log.Printf("WARN: Method not allowed for %s %s", r.Method, r.URL.Path)
		return
	}

	email := r.URL.Query().Get("email")
	ageStr := r.URL.Query().Get("age")

	var validationErr error

	if email == "" {
		validationErr = fmt.Errorf("validation failed: %w", ErrEmailEmpty)
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {

		if validationErr != nil {
			validationErr = fmt.Errorf("%w; age parsing failed: %v", validationErr, err)
		} else {
			validationErr = fmt.Errorf("age parsing failed: %v", err)
		}
	} else if age < 18 {
		// Jika age kurang dari 18
		if validationErr != nil { // Jika email sudah tidak valid
			validationErr = fmt.Errorf("%w; validation failed: %w", validationErr, ErrAgeInvalid)
		} else {
			validationErr = fmt.Errorf("validation failed: %w", ErrAgeInvalid)
		}
	}

	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseMsg := "Invalid data"

		if errors.Is(validationErr, ErrEmailEmpty) {
			responseMsg = "Email tidak boleh kosong"
		} else if errors.Is(validationErr, ErrAgeInvalid) {
			responseMsg = "umur harus 18 tahun atau lebih"
		} else if _, ok := validationErr.(*strconv.NumError); ok {
			responseMsg = "umur harus berupa angka valid"
		}

		json.NewEncoder(w).Encode(Response{Error: responseMsg})
		log.Printf("WARN: Validation failed for email='%s', age='%s'. Error: %v", email, ageStr, validationErr)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Response{Status: "ok"})
	log.Printf("INFO: Validation successful for email='%s', age=%d", email, age)
}

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/validate", validateHandler)
	log.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
