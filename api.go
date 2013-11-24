package main

import (
  "github.com/codegangsta/martini"
  "net/http"
  "./lib"
)

func main() {

  // database
  lib.InitDB()

  // martini
  m := martini.Classic()

  // routes
  m.Get("/", HelloWorld)

  // launch
  m.Run()
}

func HelloWorld(res http.ResponseWriter, req *http.Request) string {
  return "Hello world!"
}