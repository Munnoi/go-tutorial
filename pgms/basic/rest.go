package pgms

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

// In-memory data store.
var users = []User{
	{ID: 1, Name: "Bob"},
	{ID: 2, Name: "Alice"},
}

// Rest initializes the server and defines the routes.
func Rest() {
	// Register the userHandler function to handle requests to "/users".
	http.HandleFunc("/users", userHandler)
	fmt.Println("Server listening on port 8080")
	// Start the server on port 8080.
	http.ListenAndServe(":8080", nil)
}

// userHandler handles HTTP requests for the "/users" endpoint.
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to application/json for all responses.
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
		case "GET":
			// Check if an "id" query parameter is present (e.g., /users?id=1).
			idStr := r.URL.Query().Get("id")
			if idStr != "" {
				// Convert the id string to an integer.
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				// Find the user with the matching ID.
				for _, user := range users {
					if user.ID == id {
						json.NewEncoder(w).Encode(user)
						return
					}
				}
				// If user is not found, return 404.
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			// If no id is provided, return all users.
			json.NewEncoder(w).Encode(users)
		case "POST":
			var user User
			// Decode the request body into the user struct.
			json.NewDecoder(r.Body).Decode(&user)
			// Assign a new ID based on the current length of the slice.
			user.ID = len(users) + 1
			users = append(users, user)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(user)
		case "PUT":
			idStr := r.URL.Query().Get("id")
			if idStr != "" {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				var updatedUser User
				json.NewDecoder(r.Body).Decode(&updatedUser)

				for i, user := range users {
					if user.ID == id {
						// Update the user in the slice.
						updatedUser.ID = id // Ensure ID remains consistent.
						users[i] = updatedUser
						json.NewEncoder(w).Encode(updatedUser)
						return
					}
				}
				http.Error(w, "User not found", http.StatusNotFound)
			}
		case "DELETE":
			idStr := r.URL.Query().Get("id")
			if idStr != "" {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				for i, user := range users {
					if user.ID == id {
						// Remove the user from the slice using append and slicing.
						users = append(users[:i], users[i+1:]...)
						w.WriteHeader(http.StatusNoContent)
						return
					}
				}
				http.Error(w, "User not found", http.StatusNotFound)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}