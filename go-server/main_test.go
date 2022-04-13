package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/camcast3/SimpleCalculatorApi/arithmetic"
	"github.com/stretchr/testify/assert"
)

func TestPostRoute(t *testing.T) {
	router := setupRouter()
	sumSet := arithmetic.Summation{Value1: 42.2, Value2: 51.3}
	res, err := json.Marshal(sumSet)

	if err != nil {
		fmt.Println(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/summation", bytes.NewBuffer(res))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "93.5", w.Body.String())
}
