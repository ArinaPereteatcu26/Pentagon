package handlers

import (
	"fmt"
	"github.com/ArinaPereteatcu26/Pentagon/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetHotspots(c *gin.Context) {
	hotspots, err := db.GetHotspots()
	for i := range hotspots {
		photos, err := db.GetPhotos(strconv.Itoa(hotspots[i].HotspotID))
		if err != nil {
			continue
		}
		hotspots[i].Photos = photos
	}
	if err != nil {
		log.Println("Error: ", err.Error())
		c.AbortWithStatusJSON(500, "Failed to get data from database")
		return
	}
	c.JSON(200, hotspots)
}

func PostHotspot(c *gin.Context) {
	var hotspot db.Hotspot
	err := c.BindJSON(&hotspot)
	if err != nil {
		log.Println("Error", err.Error())
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	if len(hotspot.Title) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Empty title")
		return
	}

	if len(hotspot.Title) > 50 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Title too long")
		return
	}

	if len(hotspot.Description) > 1000 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Description too long")
		return
	}

	id, err := db.AddHotspot(hotspot)
	if err != nil {
		log.Println("Error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error adding hotspot")
		return
	}

	c.JSON(200, id)
	fmt.Printf("%+v", hotspot)
}
