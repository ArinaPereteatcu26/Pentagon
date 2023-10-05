package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Photos(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := fmt.Sprintf("./static/%s/%s", c.Params.ByName("id"), file.Filename)
	// Upload the file to specific dst.
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		c.AbortWithStatusJSON(500, "failed to store photo")
		log.Println(err)
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
