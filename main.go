package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"project256/actions/essay"
	"project256/actions"
)

func main() {
	r := gin.Default()
	sessStore, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "")
	r.Use(sessions.Sessions("project256", sessStore))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	essayRouter := r.Group("/essay", actions.ActionBase())
	essayRouter.GET("/list",  essay.EssayList())
	essayRouter.POST("/write", essay.WriteEssay())
	r.Run(":3001")
}
