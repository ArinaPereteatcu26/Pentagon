package main

import (
	"github.com/ArinaPereteatcu26/Pentagon/db"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	// http://localhost:8080/hotspots
	// extracting from DB the values
	// Check if this is OK
	r := gin.Default()
	r.GET("/hotspots", func(c *gin.Context) {
		hotspots, err := db.GetHotspots()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, hotspots)
	})
	// listen and serve on localhost:8080
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
