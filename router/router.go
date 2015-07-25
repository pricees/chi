package router

import (
  "regexp"
  "net/http"
  "fmt"
)

func init() {
  Routes = &RouteTable{}
}

var Routes *RouteTable

type RouteTable map[string]RouteHandlers

type RouteHandlers []*RouteHandler

type RouteHandler struct {
  routeMatcher *regexp.Regexp
  handler func(http.ResponseWriter, *http.Request)
}

func (r RouteHandlers) findRouteHandler(url string) (*RouteHandler, bool){
  for _, routeHandler := range r {
    if routeHandler.routeMatcher.Match([]byte(url)) {
      return routeHandler, true
    }
  }
  return nil, false

}

func (rh RouteHandlers) Route(url string, handler func(http.ResponseWriter, *http.Request)) RouteHandlers {
  return append(rh, &RouteHandler{ regexp.MustCompile(fmt.Sprintf("^%s$", url)), handler} )
}

func (rt *RouteTable) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    url := req.URL.String()
    if _, exists := (*rt)[req.Method]; !exists {
      http.NotFound(w, req)
    }

    if rh, _ := (*rt)[req.Method].findRouteHandler(url); rh != nil {
      rh.handler(w, req)
    }
}

func (r *RouteTable) Route(method string, url string, handler func(http.ResponseWriter, *http.Request)) {
    if _, exists := (*r)[method]; !exists {
      (*r)[method] = make(RouteHandlers, 0, 100)
    }

    if _, exists := (*r)[method].findRouteHandler(url); !exists {
      (*r)[method] = (*r)[method].Route(url, handler)
    }
}
