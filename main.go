package main

import (
	"net/http"
	"ApiVersionTry/essentials"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	clr := essentials.New()
	r.GET("/v1/employee", clr.Get)
	r.POST("/v1/employee", clr.Add)
	versionRouter := essentials.NewVersionRouter("v1", r)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        versionRouter,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
