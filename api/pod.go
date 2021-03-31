package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liuqianhong6007/k8s/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/liuqianhong6007/k8s/internal"
)

func init() {
	util.AddRoute(util.Routes{
		{
			Method:  http.MethodGet,
			Path:    "/pod/list",
			Handler: ListPod,
		},
		{
			Method:  http.MethodGet,
			Path:    "/pod",
			Handler: GetPod,
		},
	})
}

func ListPod(c *gin.Context) {
	namespace := c.Query("namespace")
	util.CheckValue(c, namespace != "", "param[namespace] is null")

	list, err := internal.K8sClientset().CoreV1().Pods(namespace).List(c, metav1.ListOptions{})
	util.CheckValue(c, err, "list pod error")

	c.JSON(http.StatusOK, list)
}

func GetPod(c *gin.Context) {
	namespace := c.Query("namespace")
	util.CheckValue(c, namespace != "", "param[namespace] is null")

	podName := c.Query("podName")
	util.CheckValue(c, podName != "", "param[podName] is null")

	pod, err := internal.K8sClientset().CoreV1().Pods(namespace).Get(c, podName, metav1.GetOptions{})
	util.CheckValue(c, err, "get pod error")

	c.JSON(http.StatusOK, pod)
}
