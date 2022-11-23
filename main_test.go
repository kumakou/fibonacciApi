package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fib?n=99", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"result":218922995834555169026}`, w.Body.String())
}