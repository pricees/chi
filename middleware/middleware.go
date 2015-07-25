package middleware

import "net/http"

type Middlewares []func(http.ResponseWriter, *http.Request) 
var Middleware Middlewares

func init() {
  Middleware = make(Middlewares, 0, 1)
}

func (m Middlewares) Add(middleware func(http.ResponseWriter, *http.Request)) {
  m = append(m, middleware)
}

func (m Middlewares) Run(w http.ResponseWriter, r *http.Request) {
  for _, middleware := range m {
    middleware(w, r)
  }
}
