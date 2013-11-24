package controllers

import (
  "net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) string {
  return "Hello world!"
}