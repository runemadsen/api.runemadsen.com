package controllers

import (
  "net/http"
)

func HomeIndex(res http.ResponseWriter, req *http.Request) string {
  return "Hello world!"
}