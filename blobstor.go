package main

import (
	"log"
	"net/http"

	_ "github.com/zhaohuXing/blobstor/router"
)

func main() {
	log.Println("Server: 127.0.0.1:8080 start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
