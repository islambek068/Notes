package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestFetchTodosFormDB(t *testing.T) {
	// Initialize test database
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("testdb").Collection("todos")
	if err = db.Drop(context.Background()); err != nil {
		t.Fatalf("Failed to drop collection: %v", err)
	}

	// Insert test data
	documents := []interface{}{
		bson.D{
			{Key: "title", Value: "Todo 1"},
			{Key: "completed", Value: false},
			{Key: "created_at", Value: time.Now()},
		},
		bson.D{
			{Key: "title", Value: "Todo 2"},
			{Key: "completed", Value: true},
			{Key: "created_at", Value: time.Now()},
		},
	}
	if _, err = db.InsertMany(context.Background(), documents); err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Call FetchTodosFormDB function
	todos, err := FetchTodosFormDB(db)
	if err != nil {
		t.Fatalf("FetchTodosFormDB returned an error: %v", err)
	}

	// Check the result
	if len(todos) != 2 {
		t.Fatalf("FetchTodosFormDB returned wrong number of todos: expected %d, got %d", 2, len(todos))
	}
	if todos[0].Title != "Todo 1" {
		t.Fatalf("FetchTodosFormDB returned wrong todo title: expected %s, got %s", "Todo 1", todos[0].Title)
	}
	if todos[0].Completed != false {
		t.Fatalf("FetchTodosFormDB returned wrong todo completion status: expected %t, got %t", false, todos[0].Completed)
	}
}
