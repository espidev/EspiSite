package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func AdminRoutes() {
	admin := router.Group(config.AdminRoute)
	admin.Use(AuthRequired())
	admin.Use(IsAdmin())

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{})
	})

	admin.GET("/new-post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new-post.html", gin.H{
			"title":   "post title",
			"content": template.HTML("<p>this is some good quality content right here!</p>"),
		})
	})

	admin.POST("/new-post", func(c *gin.Context) {

		t := time.Now()

		id := PostID{
			IDYear:  strconv.Itoa(t.Year()),
			IDMonth: t.Month().String(),
			IDDay:   strconv.Itoa(t.Day()),
			IDNum:   strconv.FormatInt(postNum, 10),
		}

		postNum++

		post := IPost{
			Name:        c.PostForm("title"),
			UserID:      "",
			Categories:  []string{},
			ID:          id,
			TimeCreated: t.Unix(),
			TimeUpdated: t.Unix(),
			Icon:        "",
			Content:     c.PostForm("content"),
		}

		db.Posts = append(db.Posts, &post)

		c.Redirect(302, "/posts/"+id.IDYear+"/"+id.IDMonth+"/"+id.IDDay+"/"+id.IDNum)

		go StoreDB()
	})
}
