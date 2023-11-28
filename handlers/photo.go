package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArinaPereteatcu26/Pentagon/db"
	"github.com/gin-gonic/gin"
)

func Photos(c *gin.Context) {
	id := c.Params.ByName("id")

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := fmt.Sprintf("./static/%s/%s", id, file.Filename)
	// Upload the file to specific dst.
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to store photo")
		log.Println(err)
		return
	}

	err = db.AddPhoto(id, file.Filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to store photo in DB")
		log.Println(err)
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
