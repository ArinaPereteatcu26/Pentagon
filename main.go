package main

import (
	"github.com/ArinaPereteatcu26/Pentagon/db"
	"github.com/ArinaPereteatcu26/Pentagon/handlers"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/hotspots", handlers.GetHotspots)

	// listen and serve on localhost:8080
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
