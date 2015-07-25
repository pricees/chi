package main

import (
  "fmt"
  "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
}

func main() {
  fmt.Println("Loading server on 8080")
  http.HandleFunc("/", helloWorld)
  http.ListenAndServe(":8080", nil)
}
