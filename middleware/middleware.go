package middleware

import "net/http"

var middlewares []func(http.ResponseWriter, *http.Request) 

func init() {
  middlewares = make([]func(http.ResponseWriter, *http.Request), 0, 1)
}

func Add(middleware func(http.ResponseWriter, *http.Request)) {
  middlewares = append(middlewares, middleware)
}

func Run(w http.ResponseWriter, r *http.Request) {
  for _, middleware := range middlewares {
    middleware(w, r)
  }
}
