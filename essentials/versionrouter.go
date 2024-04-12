package essentials

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VersionRouter struct {
	DefaultVersion string
	*gin.Engine
}

func NewVersionRouter(defaultVersion string, r *gin.Engine) *VersionRouter {
	return &VersionRouter{
		DefaultVersion: defaultVersion,
		Engine:         r,
	}
}

func (r *VersionRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	verHead := req.Header.Get("x-hous-version")
	if verHead == "" {
		verHead = r.DefaultVersion
	}
	uri := "/" + verHead + req.RequestURI
	req.RequestURI = uri
	req.URL.Path = "/" + verHead + req.URL.Path
	r.Engine.ServeHTTP(w, req)

}
