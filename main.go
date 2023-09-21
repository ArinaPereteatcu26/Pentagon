package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
)

func main() {

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(os.Getenv("DB_LOGIN"), os.Getenv("DB_PASSWORD")),
		Host:   fmt.Sprintf("%s:%d", os.Getenv("DB_HOST"), 1433),
	}

	_, err := sql.Open("sqlserver", u.String())

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
