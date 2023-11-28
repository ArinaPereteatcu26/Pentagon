package main

import (
	"log"

	"github.com/ArinaPereteatcu26/Pentagon/auth"
	"github.com/ArinaPereteatcu26/Pentagon/db"
	"github.com/ArinaPereteatcu26/Pentagon/handlers"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.POST("/login", handlers.LogIn)
	r.POST("/signup", handlers.SignUp)

	r.GET("/hotspots", handlers.GetHotspots)
	r.Static("/static", "./static")

	r.Use(auth.JWTTokenCheck)

	r.POST("/hotspots", handlers.PostHotspot)
	r.POST("/hotspots/:id", handlers.Photos)

	// listen and serve on localhost:8080
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
