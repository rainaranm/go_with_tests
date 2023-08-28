package main

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <- ctx.Done():
				s.t.Log("spy store foi cancelado")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	} ()

	select {
	case <- ctx.Done():
		return "", ctx.Err()
	case res := <- data:
		return res, nil
	}
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("não implementado")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
