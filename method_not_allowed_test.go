package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowed(t *testing.T){
	router:= httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Boleh")
	})
	router.POST("/",func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w,"Post")
	})

	request := httptest.NewRequest(http.MethodGet,"http://localhost:3000/",nil)
	recorder:= httptest.NewRecorder()

	router.ServeHTTP(recorder,request)
	response:= recorder.Result()

	bytes,_:=io.ReadAll(response.Body)
	assert.Equal(t,"Gak Boleh",string(bytes))
}