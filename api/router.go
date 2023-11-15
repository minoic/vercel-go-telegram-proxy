package api

import "net/http"
import "github.com/gin-gonic/gin"

var router *gin.Engine

func init() {
	router = gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})
}

func Listen(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
