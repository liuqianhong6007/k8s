package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/liuqianhong6007/k8s/internal"
	"github.com/liuqianhong6007/k8s/util"
)

func init() {
	util.AddRoute(util.Routes{
		{
			Method:  http.MethodGet,
			Path:    "/namespace/list",
			Handler: ListNamespace,
		},
	})
}

func ListNamespace(c *gin.Context) {
	namespaces, err := internal.K8sClientset().CoreV1().Namespaces().List(c, v1.ListOptions{})
	util.CheckValue(c, err, "list namespace error")

	c.JSON(http.StatusOK, namespaces)
}
