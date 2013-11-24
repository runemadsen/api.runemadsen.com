package main

import (
  "testing"
  "net/http"
  . "launchpad.net/gocheck"
  "../controllers"
)

func Test(t *testing.T) { TestingT(t) }
type MySuite struct{}
var _ = Suite(&MySuite{})

// GET /
// -----------------------------------------------------------

func (s *MySuite) TestHomeIndex(c *C) {
  response := testRequest("GET", "/", controllers.HomeIndex)
  c.Check(response.Code, Equals, http.StatusOK)
  c.Check(response.Body.String(), Equals, "Hello world!")
}