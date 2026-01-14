package pgms

import (
	"fmt"
	"net/http"
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
func Server() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

/*
	`http.ResponseWriter` - Used to send response to the client
	`*http.Request` - Contains request data (URL, method(GET, POST), Headers, Body)
*/
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}