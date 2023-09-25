package handlers

import (
	"github.com/ArinaPereteatcu26/Pentagon/db"
	"github.com/gin-gonic/gin"
	"log"
)

func GetHotspots(c *gin.Context) {
	hotspots, err := db.GetHotspots()
	if err != nil {
		log.Println("Error: ", err.Error())
		c.AbortWithStatusJSON(500, "Failed to get data from database")
		return
	}
	c.JSON(200, hotspots)
}
