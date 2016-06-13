package main

import (
    "net/http"
    "fmt"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Listening on port %s\n", port)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello world"))
    })
    http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}