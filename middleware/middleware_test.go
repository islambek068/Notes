package middleware

import (
	"github.com/justinas/nosurf"
	"net/http"
	"testing"
)

func TestNosurfMiddleware(t *testing.T) {
	// Create a mock handler to pass to the nosurf middleware
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do nothing
	})

	// Create a new nosurf middleware with the mock handler
	nosurfHandler := Nosurf(mockHandler)

	// Check that the returned handler is of type *nosurf.CSRFHandler
	_, ok := nosurfHandler.(*nosurf.CSRFHandler)
	if !ok {
		t.Errorf("Unexpected type for nosurfHandler. Expected *nosurf.CSRFHandler, got %T", nosurfHandler)
	}
}
