package main

import (
	"geeWeb/gee"
	"log"
	"net/http"
	"time"
)

func onlyForv2() gee.HandleFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
func main() {
	r := gee.NewEngine()
	r.Use(gee.Logger())
	//r.GET("/index", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	//})
	//v1 := r.Group("/v1")
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	//v1.GET("/hello", func(c *gee.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//})
	//v1.GET("/hello/:name", func(c *gee.Context) {
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	//})

	v2 := r.Group("/v2")
	v2.Use(onlyForv2())
	//v2.GET("/assert/*filepath", func(c *gee.Context) {
	//	c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	//})
	v2.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	//v2.POST("/login", func(c *gee.Context) {
	//	c.JSON(http.StatusOK, gee.H{
	//		"username": c.PostForm("username"),
	//		"password": c.PostForm("password"),
	//	})
	//})
	r.Run(":9999")
}