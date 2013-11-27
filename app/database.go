package app

import (
  r "github.com/dancannon/gorethink"
  "time"
)

var (
  session *r.Session
)

func InitDB() {
  
  var err error

  session, err = r.Connect(map[string]interface{} {
    "address":  "localhost:8080",
    "database": "test",
    "maxIdle": 10,
    "idleTimeout": time.Second * 10,
  })

  if err != nil {
    //log.Println(err.Error())
  }
}
