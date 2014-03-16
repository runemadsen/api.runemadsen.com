package main

import (
  "net/http"
  "github.com/codegangsta/martini"
  r "github.com/dancannon/gorethink"
  "./app"
  "log"
  "time"
)

var (
  session *r.Session
)

// Init Everything
// --------------------------------------------------------------------

func main() {

  InitDB()

  // martini
  m := martini.Classic()

  // set up to use contrib renderer
  m.Use(render.Renderer())

  // routes
  m.Get("/", app.HomeIndex)

  // launch
  http.ListenAndServe("0.0.0.0:3000", m)
}

// Setup Database
// --------------------------------------------------------------------

func InitDB() {

  var err error

  session, err = r.Connect(map[string]interface{} {
      "address" : "localhost:8080",//os.Getenv("RETHINKDB_URL"),
      "database": "test",
      "maxIdle" : 10,
      "idleTimeout": time.Second  * 10,
  })

  if err != nil {
      log.Println(err)
  }
  
  err = r.DbCreate("test").Exec(session)
  if err != nil {
      log.Println(err)
  }

  _, err = r.Db("test").TableCreate("articles").RunWrite(session)
  if err != nil {
      log.Println(err)
  }
}