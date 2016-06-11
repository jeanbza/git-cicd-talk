package main

import (
    "net/http"
    "fmt"
)

func main() {
    fmt.Println("Listening on port 8080")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello world"))
    })
    http.ListenAndServe(":8080", nil)
}