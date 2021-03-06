package main

import (
  "net/http"
  "github.com/codegangsta/martini"
  "./app"
)

var (
  session *r.Session
)

// Init Everything
// --------------------------------------------------------------------

func main() {

  session = app.InitDB()

  // martini
  m := martini.Classic()

  // set up to use contrib renderer
  m.Use(render.Renderer())

  // routes
  m.Get("/", app.HomeIndex)

  // launch
  http.ListenAndServe("0.0.0.0:3000", m)
}