package app

import (
  "github.com/codegangsta/martini-contrib/render"
)

func ArticlesIndex(r render.Render) {
  
  hal := HAL{}
  
  hal.Links = map[string]Link{
    "self"      : Link{"/articles", false},
    "show" : Link{"/articles/{id}", true},
  }

  //result := []Todo{}
  //
  //rows, err := rethink.Table("todos").Run(session)
  //if err != nil {
  //    log.Println(err)
  //}
//
  //for rows.Next() {
  //    var t Todo
  //    err := rows.Scan(&t)
  //    if err != nil {
  //        log.Println(err)
  //    }
  //    result = append(result, t)
  //}
  //r.JSON(200, result)

  //r.JSON(200, "hello")

  r.JSON(200, hal)


}