package app

import (
  "encoding/json"
)

func toJSON(hal interface{}) string {
  parsed, _ := json.Marshal(hal)
  return string(parsed)
}