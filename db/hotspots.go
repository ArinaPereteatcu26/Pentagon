package db

type Hotspot struct {
	HotspotID   int      `json:"hotspot_id"`
	Title       string   `json:"title"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	PersonID    int      `json:"-"`
}

const getHotspotQuery = `SELECT ID, Title, Latitude, Longitude, Description, PersonID
	FROM PhotoApp.dbo.Hotspots`

const addHotspotQuery = `INSERT INTO PhotoApp.dbo.Hotspots(Title, Latitude, Longitude, Description, PersonID)
	OUTPUT Inserted.ID 
	VALUES (@p1, @p2, @p3, @p4, @p5)`

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
			&hotspot.PersonID,
		)
		if err != nil {
			return nil, err
		}
		hotspots = append(hotspots, hotspot)
	}
	return hotspots, nil
}

func AddHotspot(hotspot Hotspot) (int, error) {
	var id int
	err := db.
		QueryRow(addHotspotQuery, hotspot.Title, hotspot.Latitude, hotspot.Longitude, hotspot.Description, hotspot.PersonID).
		Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
