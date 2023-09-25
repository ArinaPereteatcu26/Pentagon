package db

type Hotspot struct {
	HotspotID   int     `json:"hotspot-id"`
	Title       string  `json:"title"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
}

const getHotspotQuery = "SELECT HotspotsID, Title, Latitude, Longitude, Description FROM PhotoApp.dbo.Hotspots"

func GetHotspots() ([]Hotspot, error) {
	var hotspots []Hotspot

	rows, err := db.Query(getHotspotQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
