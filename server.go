package main

import (
  "fmt"
  "./instant"
  "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
}

func main() {

  p := 8080
  fmt.Printf("'instant' listening on %d\n", p)
  instant.Get("/", helloWorld)
  instant.Listen(p)
}
