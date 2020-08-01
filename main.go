package main

import (
	"log"
	"net/http"
	"object_storage/object"
	"os"
)

func main() {
	http.HandleFunc("/objects/", object.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
