package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func PostRoutes() {

	router.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.html", db.Posts)
	})

	router.GET("/posts/:year/:month/:day/:num", func(c *gin.Context) {
		id := PostID{IDYear: c.Params.ByName("year"),
			IDDay:   c.Params.ByName("day"),
			IDMonth: c.Params.ByName("month"),
			IDNum:   c.Params.ByName("num")}

		post, err := GetPost(id)
		if err != nil {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
			return
		}

		c.HTML(http.StatusOK, "post.html", gin.H{
			"postName":    post.Name,
			"timeUpdated": time.Unix(post.TimeUpdated, 0).Format("01/02/2006 - 3:04PM MST"),
			"content":     template.HTML(post.Content),
		})

	})
}
