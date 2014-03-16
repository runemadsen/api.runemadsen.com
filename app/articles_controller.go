package app

import (
  "github.com/codegangsta/martini-contrib/render"
)

func ArticlesIndex(r render.Render) {
  r.JSON(200, "hello")
}