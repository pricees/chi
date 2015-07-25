package instant

import (
  "fmt"
  "log"
  "net/http"
  "../router"
  "../middleware"
)

type Router interface {
  ServeHTTP(http.ResponseWriter, *http.Request)
  Route(string, string, func(http.ResponseWriter, *http.Request))
}

type Middlewarer interface {
  Add(func(http.ResponseWriter, *http.Request))
  Run(http.ResponseWriter, *http.Request)
}

var myRouter Router
var myMiddleware Middlewarer

func init() {
  myRouter = router.Routes
  myMiddleware = middleware.Middleware
}

type MyServeMux struct { }

func (m MyServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  myMiddleware.Run(w, r)
  myRouter.ServeHTTP(w, r)
}

func Listen(port int) {
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), MyServeMux{}), nil)
}

func Send(w http.ResponseWriter, text string, status string, headers map[string]string) {
  fmt.Fprintf(w, text)
}

func Get(route string, handler func(http.ResponseWriter, *http.Request)) {
  myRouter.Route("GET", route, handler)
}

func Post(route string, handler func(http.ResponseWriter, *http.Request)) {
  myRouter.Route("POST", route, handler)
}

func AddMiddleware(middleware func(w http.ResponseWriter, r *http.Request)) {
  myMiddleware.Add(middleware)
}

