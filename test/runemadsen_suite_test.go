package test_test

// http://onsi.github.io/ginkgo/#getting_ginkgo
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
  "fmt"
  "github.com/codegangsta/martini"
  "net/http"
  "net/http/httptest"
  "github.com/bitly/go-simplejson"
)

var (
  js *simplejson.Json
  response *httptest.ResponseRecorder
)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

func Request(method string, route string, handler martini.Handler) {
  m := martini.Classic()
  m.Get("/", handler)
  request, _ := http.NewRequest(method, route, nil)
  response = httptest.NewRecorder()
  m.ServeHTTP(response, request)
  js = fromJSON(response)
}

func fromJSON(response *httptest.ResponseRecorder) *simplejson.Json {
  js, err := simplejson.NewJson([]byte(response.Body.String()))
  if err != nil {
    fmt.Println("JSON Encoding error:")
    fmt.Println("---------------------------------------------")
    fmt.Println(err)
    fmt.Println("---------------------------------------------")
  }
  return js
}