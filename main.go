package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"project256/actions/essay"
	"project256/actions/wish"
	"project256/actions"
	"project256/actions/feed"
	"project256/util"
)

func main() {
	r := gin.Default()
	r.Use(util.CORSMiddleware())

	sessStore, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "")
	r.Use(sessions.Sessions("project256", sessStore))

	essayRouter := r.Group("/essay", actions.ActionBase())
	essayRouter.GET("/list/:user_id",  essay.EssayList())
	essayRouter.POST("/write", essay.WriteEssay())

	wishRouter := r.Group("/wish", actions.ActionBase())
	wishRouter.GET("/list/:user_id", wish.WishList())
	wishRouter.POST("/make", wish.MakeAWish())

	feedRouter := r.Group("/feed", actions.ActionBase())
	feedRouter.GET("/", feed.FeedList())

	r.Run(":3001")
}
