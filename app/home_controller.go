package app

import (
  "net/http"
  "encoding/json"
)

func HomeIndex(res http.ResponseWriter, req *http.Request) string {
  
  hal := HAL{}
  
  hal.Links = map[string]Link{
    "rune" : Link{"/rune", false},
  }

  parsed, err := json.Marshal(hal)
  // this err parsing should be in a toJSON function, so I can do this:
  // return toJSON(hal)
  
  return string(parsed)
}