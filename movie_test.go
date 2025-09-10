package fanart_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-fanart"
)

func TestGetMovie(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	// Replace with a valid movie ID for testing
	const movieID = "550" // Example: Fight Club

	movie, err := fanart.GetMovie(ctx, client, movieID)
	if err != nil {
		t.Fatalf("Failed to get movie: %v", err)
	}

	if movie == nil {
		t.Fatal("Expected movie, got nil")
	}

	// Add more assertions based on the expected movie data
}