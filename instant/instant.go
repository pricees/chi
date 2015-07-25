package instant

import (
  "fmt"
  "net/http"
_  "errors"
  "regexp"
)

type RouteHandler struct {
  routeMatcher *regexp.Regexp
  handlerFunc func(http.ResponseWriter, *http.Request)
}

type RouteHandlers []*RouteHandler

func (r RouteHandlers) findRouteHandler(url string) (*RouteHandler, bool){
  for _, routeHandler := range r {
    if routeHandler.routeMatcher.Match([]byte(url)) {
      return routeHandler, true
    }
  }
  return nil, false

}

func (r RouteHandlers) Route(url string, handler func(http.ResponseWriter, *http.Request)) RouteHandlers {
  return append(r, &RouteHandler{ regexp.MustCompile(fmt.Sprintf("^%s%$", url)), handler} )
}

type RouteTable map[string]RouteHandlers

func (r *RouteTable) Route(method string, url string, handler func(http.ResponseWriter, *http.Request)) {
    if _, exists := (*r)[method]; !exists {
      (*r)[method] = make(RouteHandlers, 0, 100)
    }

    if _, exists := (*r)[method].findRouteHandler(url); !exists {
      (*r)[method] = (*r)[method].Route(url, handler)
    }
    fmt.Println(Routes)
}

var Routes *RouteTable

func init() {
  Routes = &RouteTable{}
}

func Send(w http.ResponseWriter, text string) {
  fmt.Fprintf(w, text)
}

func Listen(port int) {
  http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func Get(route string, handler func(http.ResponseWriter, *http.Request)) {
  http.HandleFunc(route, handler)
}
