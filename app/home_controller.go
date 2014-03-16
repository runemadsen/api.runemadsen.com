package app

import (
  "github.com/codegangsta/martini-contrib/render"
)

func HomeIndex(r render.Render) {
  
  hal := HAL{}
  
  hal.Links = map[string]Link{
    "self"      : Link{"/", false},
    "portfolio" : Link{"/portfolio", false},
  }

  r.JSON(200, hal)
}