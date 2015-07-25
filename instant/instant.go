package instant

import (
  "fmt"
  "net/http"
  "../router"
)


func Send(w http.ResponseWriter, text string, status string, headers map[string]string) {
  fmt.Fprintf(w, text)
}

func Listen(port int) {
  http.Handle("/", router.Routes)
  http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func Get(route string, handler func(http.ResponseWriter, *http.Request)) {
  router.Routes.Route("GET", route, handler)
}

func Post(route string, handler func(http.ResponseWriter, *http.Request)) {
  router.Routes.Route("POST", route, handler)
}
