package main

import (
	"github.com/gin-gonic/gin"
	"k8s-manager/handler/node"
	"k8s-manager/handler/sys"
)

func main() {
	r := gin.Default()
	r.GET("/ping", sys.Ping)

	r.GET("/node/get", node.List)

	r.Run(":8080") // listen and serve on  0.0.0.0:8080
}
