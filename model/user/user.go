package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserID   int
	UserName string
}

// string渲染
func RenderStr(context *gin.Context) {
	context.HTML(http.StatusOK, "user/user.html", "user string")
}

/*
stuct渲染：
type H map[string]any
使用gin.H{}是一样的使用方法
*/
func RenderStruct(context *gin.Context) {
	userinfo := UserInfo{UserID: 1, UserName: "ky"}
	context.HTML(http.StatusOK, "user/struct.html", userinfo)
}

// 数组渲染
func RenderArray(context *gin.Context) {
	userarr := []int{1, 2, 3}
	context.HTML(http.StatusOK, "user/array.html", userarr)
}
