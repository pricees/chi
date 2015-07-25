package main

import (
  "fmt"
  "./instant"
  "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
  instant.Send(w, "Hello Jeah!", "not found", nil)
}

func main() {

  p := 8080

  fmt.Printf("'instant' listening on %d\n", p)
  instant.Get("/", helloWorld)
  instant.Post("/", func(w http.ResponseWriter, r *http.Request) {
    instant.Send(w, "This is a post!", "found", nil)
  })

  instant.Listen(p)
}
