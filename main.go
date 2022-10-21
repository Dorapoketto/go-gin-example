package main

import (
	"fmt"
	"github.com/Dorapoketto/go-gin-example/conf"
	"github.com/Dorapoketto/go-gin-example/routers"
	"net/http"
)

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.HTTPPort),
		Handler:        r,
		ReadTimeout:    conf.ReadTimeout,
		WriteTimeout:   conf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
