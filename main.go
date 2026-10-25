package main

import (
    "fmt"
    "net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello")
}

func main() {
	// Handlers
    http.HandleFunc("/", getHandler)

    fmt.Println("Starting server at port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting the server:", err)
    }
}
