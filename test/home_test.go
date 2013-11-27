package main

import (
  "fmt"
  "testing"
  "net/http"
  . "launchpad.net/gocheck"
  "../app"
)

func Test(t *testing.T) { TestingT(t) }
type MySuite struct{}
var _ = Suite(&MySuite{})

// GET /
// -----------------------------------------------------------

func (s *MySuite) TestHomeIndex(c *C) {
  
  response := testRequest("GET", "/", app.HomeIndex)
  parsed := getJSON(response)

  c.Check(response.Code, Equals, http.StatusOK)

  fmt.Println("JSON")
  fmt.Println(parsed["_links"])

  // getJSON should return a map[string]interface object

  // convert the _links to a map[string]Link hash

  // check that we have the correct links
}