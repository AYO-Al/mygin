package main

import (
	"github.com/gin-gonic/gin"
	"mygin/model/index"
	"mygin/model/user"
)

func main() {
	router := gin.Default()

	// router.LoadHTMLFiles(file1,file2)可以指定多个文件，但要写具体的文件名
	// /**/* 指代所有二级子目录下的所有文件
	router.LoadHTMLGlob("templates/**/*")

	// 将./static映射到/static
	router.Static("/static", "./static")

	// index/index.html更方便多级管理
	// 要在对应的html文件使用{{define index/index.html}}{{end}}
	router.GET("/", index.Hello)

	// 可以直接指定user.html但各个目录下html文件不能重名
	router.GET("/user", user.RenderStr)
	router.GET("/user_struct", user.RenderStruct)
	router.GET("/user_arr", user.RenderArray)

	router.Run(":8080")
}
