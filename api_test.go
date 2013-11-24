package main

import (
  "net/http"
  "net/http/httptest"
  "reflect"
  "testing"
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

func TestHelloWorld(t *testing.T) {
  request, _ := http.NewRequest("GET", "/", nil)
  response := httptest.NewRecorder()
  HelloWorld(response, request)
  expect(t, response.Code, http.StatusOK)
  expect(t, response.Body.String(), "Hello world!")

}