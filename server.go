package main

import (
  "fmt"
  "./instant"
  "net/http"
  "log"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
  instant.Send(w, "<h1>Hello Jeah!</h1>", "not found", nil)
}

func main() {

  // Add some routes, handlers
  instant.Get("/", helloWorld)
  instant.Post("/", func(w http.ResponseWriter, r *http.Request) {
    instant.Send(w, "This is a post!", "found", nil)
  })

  // Add some middleware
  instant.AddMiddleware(func(w http.ResponseWriter, r *http.Request) {
    log.Println(fmt.Sprintf("%s %s", r.Method, r.URL))
  })

  // Fire up the server on port 8080
  instant.Listen(8080)
}
