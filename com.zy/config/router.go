package config

import (
	"github.com/gin-gonic/gin"
	"../src/service"
	"net/http"
)



func Init_router() *gin.Engine {
	// 注册一个默认的路由器
	router := gin.Default()
	group := router.Group("/test")
 	// 最基本的用法
	group.GET("/test1", service.Func1)
	group.POST("/test2", service.Func2)
	// TODO:注意':'必须要匹配,'*'选择匹配,即存在就匹配,否则可以不考虑
	group.GET("/test3/:name/:passwd", service.Func3)
	group.POST("/test4/:name/:passwd", service.Func4)
	group.GET("/test5/:name/*passwd", service.Func5)
	// 使用gin的Query参数形式,/test6?firstname=Jane&lastname=Doe
	group.GET("/test6", service.Func6)
	group.POST("/test7", service.Func7)
	// 下面测试加载HTML: LoadHTMLTemplates
	// 加载static文件夹下所有的文件
	router.LoadHTMLGlob("com.zy/static/templates/**/*")
	// 或者使用这种方法加载也是OK的: router.LoadHTMLFiles("templates/template1.templates", "templates/template2.templates")
	router.GET("users/login", func(c *gin.Context) {
		// 注意下面将gin.H参数传入index.tmpl中!也就是使用的是index.tmpl模板
		c.HTML(http.StatusOK, "users/login.html", gin.H{
			"title": "GIN: 测试加载HTML模板",
		})
	})
	router.GET("posts/login", func(c *gin.Context) {
		// 注意下面将gin.H参数传入index.tmpl中!也就是使用的是index.tmpl模板
		c.HTML(http.StatusOK, "posts/login.html", gin.H{
			"title": "GIN: 测试加载HTML模板",
		})
	})
	return router
}