package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	fmt.Println("listening for requests on :8080")
	http.ListenAndServe(":8080", nil)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	file := os.Args[1]
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sending file to " , r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}
