package instant

import (
  "fmt"
  "net/http"
)


func Listen(port int) {
  http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func Get(route string, handler func(http.ResponseWriter, *http.Request)) {
  http.HandleFunc(route, handler)
}
