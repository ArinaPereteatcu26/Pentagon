package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ArinaPereteatcu26/Pentagon/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	claims := c.Value("claims").(*jwt.StandardClaims)
	if claims == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "No claims")
		return
	}
	fmt.Println("claim", claims.Subject)
	hotspot.PersonID, err = strconv.Atoi(claims.Subject)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
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

	c.JSON(http.StatusOK, id)
	fmt.Printf("%+v", hotspot)
}
