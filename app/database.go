package app

import(
  rethink "github.com/dancannon/gorethink"
  "log"
  "time"
)

func InitDB() *rethink.Session {

  session, err := rethink.Connect(map[string]interface{} {
    "address" : "localhost:8080",//os.Getenv("RETHINKDB_URL"),
    "database": "test",
    "maxIdle" : 10,
    "idleTimeout": time.Second  * 10,
  })

  if err != nil {
      log.Println(err)
  }
  
  err = rethink.DbCreate("test").Exec(session)
  if err != nil {
    log.Println(err)
  }

  _, err = rethink.Db("test").TableCreate("articles").RunWrite(session)
  if err != nil {
    log.Println(err)
  }

  return session
}