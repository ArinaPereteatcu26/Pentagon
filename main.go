package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
	"os"
)

var db *sql.DB

type Hotspot struct {
	HotspotID   int     `json:"hotspot-id"`
	Title       string  `json:"title"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
}

func hotspotsJson() ([]Hotspot, error) {
	var hotspots []Hotspot
	rows, err := db.Query("USE PhotoApp SELECT HotspotsID, Title, Latitude, Longitude,Description FROM Hotspots")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)
	for rows.Next() {
		var hotspot Hotspot

		err := rows.Scan(
			&hotspot.HotspotID,
			&hotspot.Title,
			&hotspot.Latitude,
			&hotspot.Longitude,
			&hotspot.Description,
		)
		if err != nil {
			return nil, err
		}
		hotspots = append(hotspots, hotspot)
	}
	return hotspots, nil
}

func main() {
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(os.Getenv("DB_LOGIN"), os.Getenv("DB_PASS")),
		Host:   fmt.Sprintf("%s:%d", os.Getenv("DB_SERVER"), 1433),
	}
	var err error
	db, err = sql.Open("sqlserver", u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	//http://localhost:8080/hotspots
	//extracting from DB the values
	//Check if this is OK
	r := gin.Default()
	r.GET("/hotspots", func(c *gin.Context) {
		hotspots, err := hotspotsJson()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, hotspots)
	})
	//listen and serve on localhost:8080
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
