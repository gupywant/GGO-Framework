package utils

import (
	"context"
	"encoding/json"
	"net/http"
)

func MethodCheckJson(handler http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var body map[string]interface{}
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(r.Context(), jsonBodyKey, body)
			r = r.WithContext(ctx)
		}

		handler(w, r)
	}
}

func MethodCheckAny(handlerFunc http.HandlerFunc, allowedMethod string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	}
}