package main

import (
  "fmt"
  "./instant"
  "net/http"
  "log"
  "encoding/json"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
  instant.Send(w, "<h1>Hello Jeah!</h1>", "continue", nil)
}

func main() {

  // Add some routes, handlers
  instant.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
    if json, err := json.Marshal("This is foo"); err == nil {
      instant.Send(w, string(json), "ok", nil)
    } else {
      log.Println(err)
    }
  })
  instant.Get("/foo/?.*", helloWorld)
  instant.Post("/", func(w http.ResponseWriter, r *http.Request) {
    instant.Send(w, "This is a post!", "found", nil)
  })

  // Add some middleware
  instant.AddMiddleware(func(w http.ResponseWriter, r *http.Request) {
    log.Println(fmt.Sprintf("%s %s", r.Method, r.URL))
  })

  alwaysJSON := func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "text/json")
  }
  instant.AddMiddleware(alwaysJSON)

  // Fire up the server on port 8080
  instant.Listen(8080)
}
