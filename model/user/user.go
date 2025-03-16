package user

import (
	"fmt"
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

// 参数绑定
type User struct {
	// json:"username
	Username string `form:"username"`
	Password int    `form:"password"`
	Addr     string `form:"addr"`
}

func UserAdd(context *gin.Context) {
	context.HTML(http.StatusOK, "user/user_add.html", "")
}

func UserToAdd(context *gin.Context) {
	var user User
	// 将参数绑定到结构体
	err := context.ShouldBind(&user)
	fmt.Println(err)
	fmt.Println(user)
	context.String(http.StatusOK, "user")
}
