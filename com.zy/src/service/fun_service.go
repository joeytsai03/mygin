package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"io"
	"log"
)


// func1: 处理最基本的GET
func Func1 (c *gin.Context)  {
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.String(http.StatusOK, "test1 OK")
}
// func2: 处理最基本的POST
func Func2 (c *gin.Context) {
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	c.String(http.StatusOK, "test2 OK")
}

// func3: 处理带参数的path-GET
// http://0.0.0.0:8888/test3/name=TAO/passwd=123
func Func3(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  test3 OK", name, passwd)
}
// func4: 处理带参数的path-POST
func Func4(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  test4 OK", name, passwd)
}
// func5: 注意':'和'*'的区别
func Func5(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  test5 OK", name, passwd)
}
// 使用Query获取参数
//http://0.0.0.0:8888/test6?name=BBB&passwd=CCC
func Func6(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Query("name")
	passwd := c.Query("passwd")
	c.String(http.StatusOK, "参数:%s %s  test6 OK", name, passwd)
}
// 使用Query获取参数
func Func7(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Query("name")
	passwd := c.Query("passwd")
	c.String(http.StatusOK, "参数:%s %s  test7 OK", name, passwd)
}
// 参数是form中获得,即从Body中获得,忽略URL中的参数
// 使用post_form形式,注意必须要设置Post的type,
// 同时此方法中忽略URL中带的参数,所有的参数需要从Body中获得
// "application/x-www-form-urlencoded"
func Func8(c *gin.Context)  {
	message := c.PostForm("message")
	extra := c.PostForm("extra")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "test8:posted",
		"message": message,
		"nick":    nick,
		"extra": extra,
	})
}
// 接收client上传的文件
// 从FormFile中获取相关的文件data!
// 然后写入本地文件
// 传输文件需要使用multipart/form-data格式的数据, 所有需要设定Post的类型是multipart/form-data.
func Func9(c *gin.Context) {
	// 注意此处的文件名和client处的应该是一样的
	file, header , err := c.Request.FormFile("uploadFile")
	filename := header.Filename
	fmt.Println(header.Filename)
	// 创建临时接收文件
	out, err := os.Create("copy_"+filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// Copy数据
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, "upload file success")
}