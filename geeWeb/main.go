package main

import (
	"geeWeb/gee"
	"net/http"
)


func main() {
	r := gee.NewEngine()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/assert/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath":c.Param("filepath")})
	})
	//r.POST("/", func(c *gee.Context) {
	//	c.JSON(http.StatusOK, gee.H{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})
	r.Run(":9999")
}