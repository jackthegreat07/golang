package main

import "testing"

func TestMain(t *testing.T) {
    expected := "Hello, Go!"
    if result := helloGo(); result != expected {
        t.Errorf("Expected %s but got %s", expected, result)
    }
}

func helloGo() string {
    return "Hello, Go!"
}

