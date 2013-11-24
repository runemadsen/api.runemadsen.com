package main

import (
  "net/http"
  "github.com/codegangsta/martini"
  "./lib"
  "./controllers"
)

func main() {

  // database
  lib.InitDB()

  // martini
  m := martini.Classic()

  // routes
  m.Get("/", controllers.HomeIndex)

  // launch
  http.ListenAndServe("0.0.0.0:3000", m)
}