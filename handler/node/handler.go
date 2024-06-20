package node

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
)

func List(c *gin.Context) {

	kubeConfigPath := ".kube/config"

	//读取配置文件内容，构建配置对象
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)

	if err != nil {
		log.Println(err)

		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})

		return
	}

	//1.获取clientSet对象
	clientSet, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Println(err)

		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})

		return
	}

	//2.获取node列表请求
	nodeList, err := clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})

	if err != nil {
		log.Println(err)

		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nodeList,
	})
}
