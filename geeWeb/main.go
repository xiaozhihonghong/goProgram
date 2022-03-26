package main

import (
	"fmt"
	"geeWeb/gee"
	"log"
	"net/http"
	"text/template"
	"time"
)

func onlyForv2() gee.HandleFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.NewEngine()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})  //自定义一个渲染函数
	r.LoadHTMLGlob("templates/*") //加载模板
	r.Static("/assets", "./static") //加载静态资源，将文件路径进行映射,handler进行渲染
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})             //执行渲染
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title": "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})
	//v1 := r.Group("/v1")
	//r.GET("/", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//})
	//v1.GET("/hello", func(c *gee.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//})
	//v1.GET("/hello/:name", func(c *gee.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})

	//v2 := r.Group("/v2")
	//v2.Use(onlyForv2())
	//v2.GET("/assert/*filepath", func(c *gee.Context) {
	//	c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	//})
	//v2.GET("/hello/:name", func(c *gee.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})
	//v2.POST("/login", func(c *gee.Context) {
	//	c.JSON(http.StatusOK, gee.H{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})
	r.Run(":9999")
}