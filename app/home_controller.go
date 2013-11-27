package app

import (
  //"fmt"
  "net/http"
  "encoding/json"
)

func HomeIndex(res http.ResponseWriter, req *http.Request) string {
  
  hal := HAL{}
  
  hal.Links = map[string]Link{
    "rune" : Link{"/rune", false},
  }

  parsed, err := json.Marshal(hal)

  //fmt.Println("JSON:")
  //fmt.Println(string(parsed))

  if err != nil {
    //fmt.Println("JSON ERROR:")
    //fmt.Println(err)
  }
  
  return string(parsed)
}