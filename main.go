package main

import "fmt"
import "net/http"


func main() {
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to my website!")
})

http.ListenAndServe(":8080", nil)
}