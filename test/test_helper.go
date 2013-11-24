package main

import (
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