package pgms

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"strings"
)

/*
	`http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))` ->  is a convenience function used to register a function as an HTTP handler for a specific URL path pattern on the default server multiplexer (DefaultServeMux)
		- pattern -> The URL path pattern (e.g., "/", "/hello", "GET /posts/{id}") to match incoming requests against
		- handler -> A function with the signature func(http.ResponseWriter, *http.Request) that will be called when a request matches the pattern

	`http.ListenAndServe(addr string, handler Handler)`
		- addr -> the address to listen on
		- handler -> is any type that implements the ServeHTTP(http.ResponseWriter, *http.Request) method 
			- If nil is passed, it uses Go's default multiplexer (http.DefaultServeMux), to which you can register handler functions using http.HandleFunc. Alternatively, you can create a custom multiplexer using http.NewServeMux() and pass that as the handler
*/
// Server initializes the HTTP server and defines the route handlers.
func Server() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/fetch", fetchHandler)
	http.HandleFunc("/createUser", createUserHandler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

/*
	`http.ResponseWriter` - Used to send response to the client
	`*http.Request` - Contains request data (URL, method(GET, POST), Headers, Body)
*/
// rootHandler handles requests to the root path "/".
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Write a simple string response to the client.
	fmt.Fprintf(w, "Hello World!")
}

// helloHandler handles requests to "/hello" and returns a JSON response.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate a JSON response.
	w.Header().Set("Content-Type", "application/json")

	// Check if the request method is GET.
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Encode a map into JSON and write it to the response.
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
}

// fetchHandler makes an external API call and returns the result.
func fetchHandler(w http.ResponseWriter, r *http.Request) {
	// Perform a GET request to an external API.
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Ensure the response body is closed after the function exits to prevent resource leaks.
	defer resp.Body.Close()

	// Copy the external API's response body to the client's response writer.
	io.Copy(w, resp.Body)
}

// createUserHandler simulates a POST request to an external service.
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Create a reader from a JSON string to use as the request body.
	data := strings.NewReader(`{"name":"Bob"}`)

	// Perform a POST request to an external URL.
	resp, err := http.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		data,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Return the status code and body of the external request to the client.
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}