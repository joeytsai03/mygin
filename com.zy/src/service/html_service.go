package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func User_login(c *gin.Context) {
	// 注意下面将gin.H参数传入index.tmpl中!也就是使用的是index.tmpl模板
	c.HTML(http.StatusOK, "users/login.html", gin.H{
		"title": "GIN: 测试加载HTML模板",
	})
}

func Posts_logion(c *gin.Context) {
	// 注意下面将gin.H参数传入index.tmpl中!也就是使用的是index.tmpl模板
	c.HTML(http.StatusOK, "posts/login.html", gin.H{
		"title": "GIN: 测试加载HTML模板",
	})
}

func Index(c *gin.Context) {
	// 注意下面将gin.H参数传入index.tmpl中!也就是使用的是index.tmpl模板
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "GIN: 测试加载HTML模板",
	})
}
