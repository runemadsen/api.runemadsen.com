package app

import (
  "net/http"
)

func HomeIndex(res http.ResponseWriter, req *http.Request) string {
  
  hal := HAL{}
  
  hal.Links = map[string]Link{
    "self"      : Link{"/", false},
    "portfolio" : Link{"/portfolio", false},
  }

  return toJSON(hal)
}