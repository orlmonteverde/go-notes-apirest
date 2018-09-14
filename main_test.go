package main

import (
    "net/http"
    "testing"
)

const domain = "http://localhost:8080"

func TestRoutes(t *testing.T) {
    r, err := http.Get(domain + "/api/users")
    if err != nil {
        t.Fatal("Not found")
    }

    if r.StatusCode != http.StatusOK {
        t.Error("Failed")
    }

}
