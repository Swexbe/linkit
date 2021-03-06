package main

import (
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swexbe/linkit/backend/controllers"
	"github.com/swexbe/linkit/backend/core"
	"github.com/swexbe/linkit/backend/middlewares"
	"github.com/swexbe/linkit/backend/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	godotenv.Load()
	db = postgres.Init()

	l1 := &core.Link{
		Shortcut: "hej2",
		LinkURL:  "google.se",
		Creator:  "admin",
	}

	s1 := &core.Session{
		Link:      l1,
		IP:        "1.1.1.1",
		Country:   "SE",
		Timestamp: time.Now(),
	}

	db.Create(l1)
	db.Create(s1)

}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	r := gin.Default()
	redisURI := os.Getenv("REDIS_URI")
	cookieSecret := os.Getenv("COOKIE_SECRET")
	// Init session middleware
	store, err := redis.NewStore(10, "tcp", redisURI, "", []byte(cookieSecret))
	panicIfErr(err)
	r.Use(sessions.Sessions("LinkITSession", store))
	r.Use(middlewares.GammaAuth())

	linkGroup := r.Group("/api/links")
	controllers.RouteLinkController(linkGroup, db)

	userGroup := r.Group("/api/user")
	controllers.RouteUserController(userGroup)

	sessionGroup := r.Group("/api/session")
	controllers.RouteSessionController(sessionGroup)

	rPublic := gin.Default()
	publicGroup := rPublic.Group("/")
	controllers.RoutePublicController(publicGroup)

	go rPublic.Run(":4001")

	r.Run(":4000")

}
