package main

import (
	"fmt"
	"net/http"
)

var server_port = 8080

func main() {
	fmt.Printf("server is started on :%d\n", server_port)

	addr := fmt.Sprintf(":%d", server_port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Please, make sure that this port is free:%d\n", server_port)
	}
}
