package main

import (
	"fmt"
	"net/http"
)

func (s *httpServer) route() {
	// Static file server.
	s.serverMux.Handle("/static/", http.FileServer(http.Dir(s.config.DocumentRoot)))

	// Other handlers.
	s.serverMux.HandleFunc("/", s.indexHandler)
	s.serverMux.HandleFunc("/test", s.testHandler)
	s.serverMux.HandleFunc("/api/v1/echo", s.echoHandler)
}

func (s *httpServer) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world\r\n")
}

func (s *httpServer) testHandler(w http.ResponseWriter, r *http.Request) {
	// Set the "hello" key in redis first: redis-cli set hello world
	// Then call this handler: curl localhost:8080/test

	// The redis connection is fault-tolerant. Try killing redis and
	// calling /test again. Then run redis and call /test again.

	if v, err := s.redis.Get("hello"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintf(w, "hello %s\r\n", v)
	}
}

func (s *httpServer) echoHandler(w http.ResponseWriter, r *http.Request) {
	appname := getURIParameter("/api/v1/echo/", r)
	if appname == "" {
		fmt.Fprintf(w, "no name after /api/v1/echo/")
		http.NotFound(w, r)
	} else {
		fmt.Fprintf(w, appname)
	}
}
