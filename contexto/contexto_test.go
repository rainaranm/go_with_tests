package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}



func TestServer(t *testing.T) {
	data := "olá, mundo"
	t.Run("retorna dados da store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
	
		svr.ServeHTTP(response, request)
	
		if response.Body.String() != data {
			t.Errorf(`resultado "%s", esperado "%s"`, response.Body.String(), data)
		}
	})

	t.Run("avisa a store para cancelar o trabalhho se a requisição for cancelada", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet,"/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("uma resposta não deveria ter sido escrita")
		}
	})
}