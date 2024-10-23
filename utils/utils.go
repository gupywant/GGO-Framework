package utils

import (
	"net/http"
	"encoding/json"
	"fmt"
	_"log"
	"os"
)

type contextKey string

const jsonBodyKey contextKey = "jsonBody"

// GetJSONBody retrieves the parsed JSON body from the context
func GetJSONBody(r *http.Request) map[string]interface{} {
	body, _ := r.Context().Value(jsonBodyKey).(map[string]interface{})
	return body
}

func DDserver(v interface{}) {
    fmt.Printf("%+v\n", v)
    os.Exit(1)
}

func ODserver(v interface{}) {
    fmt.Printf("%+v\n", v)
}

func DDclient(w http.ResponseWriter, v interface{}) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Format the variable
	response := map[string]interface{}{
		"data": v,
	}

	// Send the response
	json.NewEncoder(w).Encode(response)

	// Log the output to the console for server-side debugging
	fmt.Printf("%+v\n", v) // Print to console for server-side logs
}