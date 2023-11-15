package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var router *gin.Engine

var apiUrl = "https://api.telegram.org"

func init() {
	router = gin.Default()
	router.Any("/*path", func(context *gin.Context) {
		uri := context.Param("path")
		if !strings.Contains(uri, "bot") {
			context.String(http.StatusNotFound, "404 Not found")
			return
		}
		url := apiUrl + uri
		req, err := http.NewRequestWithContext(context, context.Request.Method, url, context.Request.Body)
		if err != nil {
			fmt.Println(err)
			context.String(http.StatusBadRequest, err.Error())
			return
		}
		req.Header = context.Request.Header
		req.PostForm = context.Request.PostForm
		req.Form = context.Request.Form
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			context.String(http.StatusBadRequest, err.Error())
			return
		}
		context.DataFromReader(resp.StatusCode, resp.ContentLength, "application/json", resp.Body, nil)
	})
}

func Listen(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
