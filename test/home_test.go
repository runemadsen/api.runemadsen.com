package main

import (
  "net/http"
  "testing"
  "../controllers"
)

func TestHomeIndex(t *testing.T) {
  response := testRequest("GET", "/", controllers.HomeIndex)
  expect(t, response.Code, http.StatusOK)
  expect(t, response.Body.String(), "Hello world!")
}