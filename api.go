package main

import (
  "net/http"
  "github.com/codegangsta/martini"
  "./app"
)

func main() {

  // database
  app.InitDB()

  // martini
  m := martini.Classic()

  // routes
  m.Get("/", app.HomeIndex)

  // launch
  http.ListenAndServe("0.0.0.0:3000", m)
}