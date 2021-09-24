package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:fmt.Printf("success!")
	log.Fatal(http.ListenAndServe(":1999", http.FileServer(http.Dir("./"))))

}
