package main

import (
  "fmt"
  "net/http"
)

func helloRest(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
}

func main() {
  fmt.Printf("Starting server...\n")
  http.HandleFunc("/", helloRest)
  http.ListenAndServe(":3535", nil)
}
