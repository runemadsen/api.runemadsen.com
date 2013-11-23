package main

import (
  "net/http"
  "net/http/httptest"
  "reflect"
  "testing"
)

// Test Helpers
// -----------------------------------------------------------------------

func expect(t *testing.T, a interface{}, b interface{}) {
  if a != b {
    t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
  }
}

func refute(t *testing.T, a interface{}, b interface{}) {
  if a == b {
    t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
  }
}

func TestSomething(t *testing.T) {
  //expect(t, "Rune", "Rune")

  //recorder := httptest.NewRecorder()
  //req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
  //if err != nil {
  //  t.Error(err)
  //}
  //context := New().createContext(recorder, req)
  //router.Handle(recorder, req, context)
  
  handler := func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "something failed", http.StatusInternalServerError)
  }
  
  req, err := http.NewRequest("GET", "http://localhost:3000", nil)
  if err != nil {
    t.Error(err)
  }
  
  recorder := httptest.NewRecorder()
  handler(recorder, req)
  
  expect(t, recorder.Code, http.StatusOK)
  expect(t, recorder.Body.String(), "Hello world!")

}