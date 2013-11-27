package test_test

import (
	"../app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  //"fmt"
)

var _ = Describe("Home", func() {

  It("returns a valid _links object", func() {
    
    response := testRequest("GET", "/", app.HomeIndex)
    js := fromJSON(response)

    Expect(response.Code).To(Equal(200))
    Expect(js.Get("_links").Get("self").Get("href").MustString()).To(Equal("/"))
    Expect(js.Get("_links").Get("portfolio").Get("href").MustString()).To(Equal("/portfolio"))
  })

})
