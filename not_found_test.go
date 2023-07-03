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

func TestNotFound(t *testing.T){
	router:= httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "404 Not Found alias Gak Ketemu")
	})

	request := httptest.NewRequest(http.MethodGet,"http://localhost:3000/",nil)
	recorder:= httptest.NewRecorder()

	router.ServeHTTP(recorder,request)
	response:= recorder.Result()

	bytes,_:=io.ReadAll(response.Body)
	assert.Equal(t,"404 Not Found alias Gak Ketemu",string(bytes))
}