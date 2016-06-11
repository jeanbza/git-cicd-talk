package main_test

import (
    "testing"
    "net/http"
    "fmt"
)

// "go run main.go" first
func TestIntegration(t *testing.T) {
    resp, err := http.Get("http://localhost:8080")

    if err != nil || resp.StatusCode != 200 {
        fmt.Printf("Failed with resp %v error %v \n", resp, err)
        t.Fail()
    }
}
