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

func TestRouter(t *testing.T){
	router:= httprouter.New()

	router.GET("/",func(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
		fmt.Fprint(writer,"Hello Jigong")
	})

	request := httptest.NewRequest(http.MethodGet,"http://localhost:3000/",nil)
	recorder:= httptest.NewRecorder()

	router.ServeHTTP(recorder,request)
	response:= recorder.Result()

	bytes,_:=io.ReadAll(response.Body)
	assert.Equal(t,"Hello Jigong",string(bytes))
}