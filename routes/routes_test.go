package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoHandler(t *testing.T) {
	// Create a new request that simulates a GET request to the "/" endpoint
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new ResponseRecorder (which implements the http.ResponseWriter interface)
	rec := httptest.NewRecorder()

	// Call the todoHandler function and pass in the ResponseRecorder and the GET request
	handler := todoHandler()
	handler.ServeHTTP(rec, req)

	// Check that the response status code is 200 OK
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %v but got %v", http.StatusOK, rec.Code)
	}

	// Check that the response body is not empty
	if len(rec.Body.Bytes()) == 0 {
		t.Error("Expected non-empty response body but got empty body")
	}
}
