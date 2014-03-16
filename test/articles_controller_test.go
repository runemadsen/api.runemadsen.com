package test_test

import (
  "../app"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  //"fmt"
)

var _ = Describe("Articles", func() {

  Context("#index", func() {
    
    It("returns a 200 Status Code", func() {
        Request("GET", "/articles", app.ArticlesIndex)
        Expect(response.Code).To(Equal(200))
    })

  })

})