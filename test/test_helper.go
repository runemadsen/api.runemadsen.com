package main

import (
  "fmt"
  "encoding/json"
  "github.com/codegangsta/martini"
  "net/http"
  "net/http/httptest"
)

// Test Helpers
// -----------------------------------------------------------------------

func testRequest(method string, route string, handler martini.Handler) *httptest.ResponseRecorder {
  m := martini.Classic()
  m.Get("/", handler)
  request, _ := http.NewRequest(method, route, nil)
  response := httptest.NewRecorder()
  m.ServeHTTP(response, request)
  return response
}

func getJSON(response *httptest.ResponseRecorder) map[string]interface{} {
  var obj interface{}
  dec := json.NewDecoder(response.Body)
  err := dec.Decode(&obj)
  if err != nil {
    fmt.Println("JSON Encoding error:")
    fmt.Println("---------------------------------------------")
    fmt.Println(err)
    fmt.Println("---------------------------------------------")
  }
  return obj.(map[string]interface{})
}