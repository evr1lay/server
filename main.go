package main

import (
	"fmt"
	"net/http"
)

var server_port = 8080

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Method)
	fmt.Println(r.Method)
}

func main() {
	http.HandleFunc("/", defaultHandler)

	fmt.Printf("server is started on http://localhost:%d\n", server_port)

	addr := fmt.Sprintf(":%d", server_port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Please, make sure that this port is free:%d\n", server_port)
	}
}
