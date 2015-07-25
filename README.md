![Chicago Flag](http://www.billslater.com/mychicago/Chicago_Flag_1.png)

chi
===

Minimal web framework 

```go
package main

import (
  "fmt"
  "./chi"
  "net/http"
  "log"
  "encoding/json"
)

// Basic "Hello World" function
func helloWorld(w http.ResponseWriter, r *http.Request) {
  chi.Send(w, "<h1>Hello World!</h1>", "continue", nil)
}

func main() {

  // GET /foo
  // Response is JSON
  chi.Get("/foo", func(w http.ResponseWriter, r *http.Request) {
    if json, err := json.Marshal("This is foo"); err == nil {
      chi.Send(w, string(json), "ok", nil)
    } else {
      log.Println(err)
    }
  })

  // GET /foo
  // Response is the 'Hello World!'
  chi.Get("/foo/?.*", helloWorld)

  // POST /
  // Response is the 'This is a post!'
  chi.Post("/", func(w http.ResponseWriter, r *http.Request) {
    chi.Send(w, "This is a post!", "found", nil)
  })

  // Add logger middleware
  chi.AddMiddleware(func(w http.ResponseWriter, r *http.Request) {
    log.Println(fmt.Sprintf("%s %s", r.Method, r.URL))
  })

  // Add middleware that will change the content-type for 
  // all requests
  alwaysJSON := func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "text/json")
  }
  chi.AddMiddleware(alwaysJSON)

  // Fire up the server on port 8080
  chi.Listen(8080)
}
```

### Demo ###

__Open one tab. Start the app.__
```bash
$ go run driver.go
```

__Open another tab, use from curl commands to test routes__
```bash
$ curl "http://localhost:8080/foo"
# => "This is foo"

$ curl "http://localhost:8080/foo123"
# => "<h1>Hello World!</h1>"

$ curl "http://localhost:8080/foo/bar/baz" # Optional slash
# => "<h1>Hello World!</h1>"

$ curl -X POST "http://localhost:8080/"
# => "This is a post!"

$ curl -I -X PUT "http://localhost:8080/foo/bar/baz" 
# HTTP/1.1 404 Not Found
# ...
```
_Note: The "Not Found" error is improper. We should respond with a *405 Method Not Allowed*. That is something you should handle._
