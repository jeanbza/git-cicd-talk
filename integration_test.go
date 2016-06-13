package main_test

import (
    "testing"
    "net/http"
    "fmt"
    "os"
)

// "go run main.go" first
func TestIntegration(t *testing.T) {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    resp, err := http.Get(fmt.Sprintf("http://localhost:%s", port))

    if err != nil || resp.StatusCode != 200 {
        fmt.Printf("Failed with resp %v error %v \n", resp, err)
        t.Fail()
    }
}
