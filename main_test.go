package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// n=1を指定したときのフィボナッチ数が合っているか
func TestGettingRoute1(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fib?n=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"result":1}`, w.Body.String())

}

// n=2を指定したときのフィボナッチ数が合っているか
func TestGettingRoute2(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fib?n=2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"result":1}`, w.Body.String())

}

// n=99を指定したときのフィボナッチ数が合っているか
func TestGettingRoute99(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fib?n=99", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"result":218922995834555169026}`, w.Body.String())

}

// ルーティングにないパスへのアクセス
func TestNoRoute(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hoge", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"message":"no route","status":404}`, w.Body.String())
}

// 	リクエストパラメータnがマイナス
func TestGettingRouteMinus(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fib?n=-98", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"message":"Please set positive number in request parameter","status":404}`, w.Body.String())
}

// 	リクエストパラメータnが文字列
func TestGettingRouteString(t *testing.T) {
	router := router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fib?n=hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"message":"Please set positive number in request parameter","status":404}`, w.Body.String())
}