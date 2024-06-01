package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func jsonRequest(r *http.Request, destination any) error {
	contentType := strings.ToLower(strings.TrimSpace(strings.Split(r.Header.Get("Content-Type"), ";")[0]))
	if contentType != "application/json" {
		return fmt.Errorf("invalid_content_type")
	}

	// Read body
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &destination)
	if err != nil {
		return err
	}

	return nil
}

func jsonResponse(statusCode int, body any, w http.ResponseWriter) {
	// Set Content-Type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	// Serialize body if not nil; panic on failure
	if body != nil {
		json, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		w.Write(json)
	}
}
