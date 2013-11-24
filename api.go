package main

import (
  "github.com/codegangsta/martini"
  "./lib"
  "./controllers"
)

var(
  m *martini.ClassicMartini
)

func main() {

  // database
  lib.InitDB()

  // martini
  m = martini.Classic()

  // routes
  m.Get("/", controllers.HelloWorld)

  // launch
  m.Run()
}