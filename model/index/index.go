package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(context *gin.Context) {
	context.HTML(http.StatusOK, "index/index.html", gin.H{
		"msg": "Hello Gin",
	})
}
