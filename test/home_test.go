package main

import (
  "github.com/codegangsta/martini"
  "net/http"
  "net/http/httptest"
  "reflect"
  "testing"
  "../controllers"
)

// use this for initing stuff in the test
//func init() {
//   initDb()
//}

// Test Helpers
// -----------------------------------------------------------------------

func expect(t *testing.T, a interface{}, b interface{}) {
  if a != b {
    t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
  }
}

func refute(t *testing.T, a interface{}, b interface{}) {
  if a == b {
    t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
  }
}

func testRequest(method string, route string, handler martini.Handler) *httptest.ResponseRecorder {
  m := martini.Classic()
  m.Get("/", handler)
  request, _ := http.NewRequest(method, route, nil)
  response := httptest.NewRecorder()
  m.ServeHTTP(response, request)
  return response
}

func TestHomeIndex(t *testing.T) {
  response := testRequest("GET", "/", controllers.HomeIndex)
  expect(t, response.Code, http.StatusOK)
  expect(t, response.Body.String(), "Hello world!")
}