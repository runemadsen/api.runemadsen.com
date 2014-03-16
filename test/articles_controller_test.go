package test_test

import (
  "../app"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  //"log"
  //"fmt"
)

var _ = Describe("Articles", func() {

  Context("#index", func() {

    BeforeEach(func() {
      //article := Article{"My Heading", "My Body!"}
      // _, err := rethink.Table("articles").Insert(article).RunWrite(session)
      //if err != nil {
      //  log.Println(err)
      //} 
    })
    
    It("lets test", func() {

    })

    It("returns 200 ok", func() {
      Request("GET", "/articles", app.ArticlesIndex)
      Expect(response.Code).To(Equal(200))
    })

    It("returns links", func() {    
      Request("GET", "/articles", app.ArticlesIndex)
      Expect(responseJSON.Get("_links").Get("self").Get("href").MustString()).To(Equal("/articles"))
      //Expect(responseJSON.Get("_links").Get("next").Get("href").MustString()).To(Equal("/articlesSOMEPAGE"))
      Expect(responseJSON.Get("_links").Get("show").Get("href").MustString()).To(Equal("/articles/{id}"))
      //Expect(responseJSON.Get("_links").Get("show").Get("templates").MustString()).To(Equal(true))
    })

    It("returns articles", func() {
      Request("GET", "/articles", app.ArticlesIndex)
      
    })

  })

})