package instant

import (
  "fmt"
  "log"
  "net/http"
  "../router"
)


var middlewares []Middleware
type Middleware struct { }

func (m Middleware) Run(w http.ResponseWriter, r *http.Request) {

}

type MyServeMux struct { }

func (m MyServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  for _, m := range middlewares {
    m.Run(w, r)
  }
  router.Routes.ServeHTTP(w, r)
}

func Send(w http.ResponseWriter, text string, status string, headers map[string]string) {
  fmt.Fprintf(w, text)
}

func Listen(port int) {
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), MyServeMux{}), nil)
}

func Get(route string, handler func(http.ResponseWriter, *http.Request)) {
  router.Routes.Route("GET", route, handler)
}

func Post(route string, handler func(http.ResponseWriter, *http.Request)) {
  router.Routes.Route("POST", route, handler)
}
