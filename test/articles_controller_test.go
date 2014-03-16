package test_test

import (
  "../app"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  //"fmt"
)

var _ = Describe("Articles", func() {

  Context("#index", func() {
    
    It("returns 200 ok", func() {
      Request("GET", "/articles", app.ArticlesIndex)
      Expect(response.Code).To(Equal(200))
    })

    It("returns a valid _links object", func() {    
      Request("GET", "/articles", app.ArticlesIndex)
      Expect(responseJSON.Get("_links").Get("self").Get("href").MustString()).To(Equal("/articles"))
      //Expect(responseJSON.Get("_links").Get("next").Get("href").MustString()).To(Equal("/articlesSOMEPAGE"))
      Expect(responseJSON.Get("_links").Get("show").Get("href").MustString()).To(Equal("/articles/{id}"))
      //Expect(responseJSON.Get("_links").Get("show").Get("templates").MustString()).To(Equal(true))
    })

    It("has the 5 lastest articles in _embedded", func() {
      Request("GET", "/articles", app.ArticlesIndex)
      
    })

  })

})