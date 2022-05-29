package main

import (
	"github.com/ccchooko/HookoGee/gee"
	"log"
	"net/http"
)
func main() {
	r := gee.New()
	r.GET("/index", indexHandler)
	//r.GET("/hello/:name", helloHandler)
	//r.POST("/login", loginHandler)
	//r.GET("/assets/*filepath", nil)
	v1 := r.Group("/v1")
	v1.GET("/hello/:name", helloHandler)

	v11 := v1.Group("/v11")
	v11.GET("/hello/:name", helloHandler)

	log.Fatal(r.Run(":9999"))
}

func indexHandler(c *gee.Context) {
	c.HTML(http.StatusOK, "<h1>Hello HookoGee</h1>")
}

func helloHandler(c *gee.Context) {
	// expect /hello?name=geektutu
	//c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Params["name"], c.Path)
}

func loginHandler(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"username": c.PostForm("username"),
		"passwd": c.PostForm("passwd"),
	})
}
