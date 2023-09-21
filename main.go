package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

var db *sql.DB

type Hotspot struct {
	HotspotID string `json:"HotspotID"`
	Title     string `json:"Title"`
	Latitude  string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

func hotspotsJson() []Hotspot {
	var hotspots []Hotspot
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
	}
	//var hotspotsJson, _ = json.MarshalIndent(hotspots, "", "  ") // Indent with two spaces
	return hotspots
}

func main() {
	fmt.Println("-----------------------------------------------------------")

	//extracting from DB the values
	hotspots := hotspotsJson()

	//Here is created a map with keys(title of hotspot) and values(coordinates of hotspot)
	//Don't know if we will use it but let it be for the moment
	var hotspotsMap = make(map[string]string)
	j := len(hotspots)
	for i := 0; i < j; i++ {
		hotspotsMap[hotspots[i].Title] = hotspots[i].Latitude + "," + hotspots[i].Longitude
	}
	for key := range hotspotsMap {
		fmt.Printf(" %s,  %s\n",
			key, hotspotsMap[key])
	}

	//http://localhost:8080/hotspots
	//Check if this is OK
	r := gin.Default()
	r.GET("/hotspots", func(c *gin.Context) {
		// Marshal the hotspots data with indentation
		hotspotsJSON, err := json.MarshalIndent(hotspots, "", "  ")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, string(hotspotsJSON))
	})
	//listen and serve on 0.0.0.0:8080
	var err = r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
