package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"net/url"
)

var db *sql.DB

type Hotspot struct {
	HotspotID string `json:"HotspotID"`
	Title     string `json:"Title"`
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

func main() {
	var hotspots []Hotspot
	var hotspotsMap = make(map[string]string)
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("sa", "Skies-Backer-Overlord1-Voucher"),
		Host:   fmt.Sprintf("%s:%d", "49.13.85.200", 1433),
	}

	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("USE PhotoApp SELECT HotspotsID, Title, Latitude, Longitude FROM Hotspots")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var hotspot Hotspot

		err := rows.Scan(
			&hotspot.HotspotID,
			&hotspot.Title,
			&hotspot.Latitude,
			&hotspot.Longitude,
		)
		if err != nil {
			panic(err)
		}
		hotspots = append(hotspots, hotspot)
		hotspotsMap[hotspot.Title] = hotspot.Latitude + "," + hotspot.Longitude
	}
	var hotspotsJson, _ = json.MarshalIndent(hotspots, "", "  ") // Indent with two spaces
	fmt.Println(string(hotspotsJson))
	fmt.Println("-----------------------------------------------------------")
	for key := range hotspotsMap {
		fmt.Printf(" %s,  %s\n",
			key, hotspotsMap[key])
	}

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//// listen and serve on 0.0.0.0:8080
	//err = r.Run(":8080")
	//if err != nil {
	//	panic(err)
	//}
}
