package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/k8s/util"
)

func init() {
	util.AddRoute(util.Routes{
		{
			Method:  http.MethodGet,
			Path:    "/ping",
			Handler: ping,
		},
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
