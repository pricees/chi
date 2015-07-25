package middleware

import "net/http"

type middleware func(http.ResponseWriter, *http.Request)
type Middlewares struct {
  middlewares []middleware
}

func NewMiddlewares() Middlewares {
  return Middlewares{ middlewares: make([]middleware, 0, 1) }
}

var Middleware *Middlewares

func init() {
  m := NewMiddlewares()
  Middleware = &m
}

func (m *Middlewares) Add(middleware func(http.ResponseWriter, *http.Request)) {
  m.middlewares = append(m.middlewares, middleware)
}

func (m Middlewares) Run(w http.ResponseWriter, r *http.Request) {
  for _, middleware := range m.middlewares {
    middleware(w, r)
  }
}
