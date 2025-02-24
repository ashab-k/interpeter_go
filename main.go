package main

import (
	"fmt"
	"net/http"

	"interpreter/endpoint"
)

func main() {
 mux := http.NewServeMux()
 
 mux.HandleFunc("/run" , endpoint.RunCodeHandler )
 fmt.Println("Listening on port 8080")
 http.ListenAndServe(":8080" , mux)
}