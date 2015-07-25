package main

import (
  "fmt"
  "./chi"
  "net/http"
  "log"
  "encoding/json"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
  chi.Send(w, "<h1>Hello Jeah!</h1>", "continue", nil)
}

func main() {

  // Add some routes, handlers
  chi.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
    if json, err := json.Marshal("This is foo"); err == nil {
      chi.Send(w, string(json), "ok", nil)
    } else {
      log.Println(err)
    }
  })
  chi.Get("/foo/?.*", helloWorld)
  chi.Post("/", func(w http.ResponseWriter, r *http.Request) {
    chi.Send(w, "This is a post!", "found", nil)
  })

  // Add some middleware
  chi.AddMiddleware(func(w http.ResponseWriter, r *http.Request) {
    log.Println(fmt.Sprintf("%s %s", r.Method, r.URL))
  })

  alwaysJSON := func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "text/json")
  }
  chi.AddMiddleware(alwaysJSON)

  // Fire up the server on port 8080
  chi.Listen(8080)
}
