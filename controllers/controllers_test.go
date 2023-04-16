package controllers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	// backup environment variables
	mongoUriOrig := os.Getenv("URI")
	dbOrig := os.Getenv("DB")
	colOrig := os.Getenv("COL")

	// set up mock environment variables
	os.Setenv("URI", "mongodb://127.0.0.1:27017/")
	os.Setenv("DB", "Notes")
	os.Setenv("COL", "Notes")

	// assert that the db variable is set correctly
	if db.Name() != "Notes" {
		t.Errorf("Expected db collection name to be 'mockCOL', got %s", db.Name())
	}

	// restore original environment variables
	os.Setenv("URI", mongoUriOrig)
	os.Setenv("DB", dbOrig)
	os.Setenv("COL", colOrig)
}

func TestCreateTodo(t *testing.T) {
	// Set up a mock HTTP request and response.
	req, err := http.NewRequest("POST", "/todo", strings.NewReader(`{"title":"Test todo"}`))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Call the CreateTodo function with the mock request and response.
	CreateTodo(rr, req)

	// Check the response status code.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body for the expected message.
	expectedBody := `{"message":"created successfully"`
	if !strings.HasPrefix(rr.Body.String(), expectedBody) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
	}
}

func TestDeleteOneTodo(t *testing.T) {
	// Create a new request with a URL that includes a valid object ID
	req, err := http.NewRequest("DELETE", "/api/todos/643c1d4b23c27d0082f9a1e4", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response from the handler
	rr := httptest.NewRecorder()

	// Create a new router and add the handler to it
	r := chi.NewRouter()
	r.Delete("/api/todos/{id}", DeleteOneTodo)

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the body of the response
	expected := `{"message":"successfully deleted"`
	if !strings.HasPrefix(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
