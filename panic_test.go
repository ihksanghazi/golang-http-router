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

func TestPanic(t *testing.T){
	router:= httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		fmt.Fprint(w,"Panic : ",err)
	}

	router.GET("/",func(writer http.ResponseWriter,request *http.Request,params httprouter.Params){
		panic("Hayooluh")
	})

	request := httptest.NewRequest(http.MethodGet,"http://localhost:3000/",nil)
	recorder:= httptest.NewRecorder()

	router.ServeHTTP(recorder,request)
	response:= recorder.Result()

	bytes,_:=io.ReadAll(response.Body)
	assert.Equal(t,"Panic : Hayooluh",string(bytes))
}