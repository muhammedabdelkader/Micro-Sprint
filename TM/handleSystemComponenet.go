package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SystemData struct {
	Nodes []struct{ ID string }
	Edges []struct{ From, To string }
}

type SystemsComponentsMiddleware struct {
	systemsData map[string]SystemData
	next        http.Handler
}

func (m *SystemsComponentsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Extract the systems components data from the request
	var data SystemData
	json.NewDecoder(r.Body).Decode(&data)

	// Store the data in the middleware's map
	systemID := data.Nodes[0].ID
	m.systemsData[systemID] = data

	// Pass the request to the next handler in the
	// chain
	m.next.ServeHTTP(w, r)

}

func NewSystemsComponentsMiddleware(next http.Handler) http.Handler {
	systemsData := make(map[string]SystemData)
	return &SystemsComponentsMiddleware{
		systemsData: systemsData,
		next:        next,
	}

}

func submitGraph(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Graph submitted successfully")

}

func viewGraph(w http.ResponseWriter, r *http.Request) {
	// Retrieve the systems components data from the middleware
	systemsComponentsMiddleware, ok := r.Context().Value(&SystemsComponentsMiddleware{}).(*SystemsComponentsMiddleware)
	if !ok {
		http.Error(w, "SystemsComponentsMiddleware not found in request context", http.StatusInternalServerError)
		return

	}
	data, ok := systemsComponentsMiddleware.systemsData["systemID"]
	if !ok {
		http.Error(w, "System data not found", http.StatusBadRequest)
		return

	}

	// Render the data as HTML or JSON
	json.NewEncoder(w).Encode(data)

}
func main() {
	// Create a new instance of the middleware
	systemsComponentsMiddleware := NewSystemsComponentsMiddleware(http.DefaultServeMux)

	// Register the handler functions
	http.Handle("/submitGraph", systemsComponentsMiddleware.ThenFunc(submitGraph))
	http.Handle("/viewGraph", systemsComponentsMiddleware.ThenFunc(viewGraph))

	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)

	}

}
