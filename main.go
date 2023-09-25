package main

import (
	"io/ioutil"
	"os"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	file := os.Args[1]
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}
