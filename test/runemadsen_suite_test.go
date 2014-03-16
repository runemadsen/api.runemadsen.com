package test_test

import (
  r "github.com/dancannon/gorethink"
  "../app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
  "fmt"
  "net/http"
  "net/http/httptest"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
  "github.com/bitly/go-simplejson"
  //"github.com/codegangsta/martini-contrib/binding"
)

var (
  response *httptest.ResponseRecorder
  responseJSON *simplejson.Json
  session *r.Session
)

func TestTest(t *testing.T) {
  session = app.InitDB()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

func Request(method string, route string, handler martini.Handler) {
  m := martini.Classic()
  m.Get(route, handler)
  m.Use(render.Renderer())
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
  responseJSON = fromJSON(response)
}

func fromJSON(response *httptest.ResponseRecorder) *simplejson.Json {
  js, err := simplejson.NewJson([]byte(response.Body.String()))
  if err != nil {
    fmt.Println("JSON Encoding error:")
    fmt.Println("---------------------------------------------")
    fmt.Println(err)
    fmt.Println("---------------------------------------------")
  }
  return js
}

//func PostRequest(method string, route string, handler martini.Handler, body io.Reader) {
//  m := martini.Classic()
//  m.Post(route, binding.Json(Todo{}), handler)
//  m.Use(render.Renderer())
//  request, _ := http.NewRequest(method, route, body)
//  response = httptest.NewRecorder()
//  m.ServeHTTP(response, request)
//}