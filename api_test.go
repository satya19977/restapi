package main

import (
	"bytes"
	"os"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/satya19977/restapi/db"
	"github.com/satya19977/restapi/routes" 
)

// Initialize the DB
func TestMain(m *testing.M) {
	// Initialize the database connection
	db.InitDB()
	// Run the tests
	exitCode := m.Run()

	// Clean up resources, if any

	// Exit with the test exit code
	os.Exit(exitCode)
}

// Helper function to perform HTTP request
func performRequest(method, path string, payload []byte, router *gin.Engine) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func TestGetEvents(t *testing.T) {
	router := gin.New()
	routes.RegisterRoutes(router)

	// Make GET request to "/v1/events"
	rr := performRequest("GET", "/v1/events", nil, router)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPostEvents(t *testing.T) {
	router := gin.New()
	routes.RegisterRoutes(router)

	// Define the payload data
	requestBody := []byte(`{
		"name": "for testing",
		"description": "for testing event ",
		"location": "NY/NJ 33 ",
		"dateTime": "2025-01-01T15:30:00.000Z"
	}`)

	// Make POST request to "/v1/events" with payload data
	rr := performRequest("POST", "/v1/events", requestBody, router)

	// Check the response status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestPutEvents(t *testing.T) {
	router := gin.New()
	routes.RegisterRoutes(router)

	// Define the payload data
	requestBody := []byte(`{
		"name": "put test",
		"description": "Updated put testdescription",
		"location": "Test put location",
		"dateTime": "2025-01-01T15:30:00.000Z"
	}`)

	// Make PUT request to "/v1/events/:id" with payload data
	rr := performRequest("PUT", "/v1/events/6", requestBody, router)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteEvents(t *testing.T) {
	router := gin.New()
	routes.RegisterRoutes(router)

	// Make DELETE request to "/v1/events/:id"
	rr := performRequest("DELETE", "/v1/events/6", nil, router)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
