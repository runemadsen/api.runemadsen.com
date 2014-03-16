package app

import (
  "log"
  "time"
  "os"
  r "github.com/dancannon/gorethink"
)

func InitDB() *r.Session {

  session, err := r.Connect(map[string]interface{} {
      "address" : os.Getenv("RETHINKDB_URL"),
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

  return session
}