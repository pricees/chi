package middleware

import "net/http"
import "fmt"

type middleware func(http.ResponseWriter, *http.Request)
type Middlewares struct {
  middlewares []middleware
}

func NewMiddlewares() Middlewares {
  return Middlewares{ middlewares: make([]middleware, 0, 1) }
}

var Middleware Middlewares

func init() {
  Middleware = NewMiddlewares()
}

func (m *Middlewares) Add(middleware func(http.ResponseWriter, *http.Request)) {
  fmt.Println("Adding middleware")
  m.middlewares = append(m.middlewares, middleware)
  fmt.Println(len(m.middlewares))
}

func (m Middlewares) Run(w http.ResponseWriter, r *http.Request) {
  for _, middleware := range m.middlewares {
    middleware(w, r)
  }
}
