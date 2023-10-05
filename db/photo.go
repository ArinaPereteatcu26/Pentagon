package db

const getPhotosQuery = "SELECT Location FROM PhotoApp.dbo.Photos WHERE ID = @p1"
const addPhotoQuery = "INSERT INTO PhotoApp.dbo.Photos(ID,Location) VALUES (@p1, @p2)"

func GetPhotos(hotspotID string) ([]string, error) {
	var photos []string

	rows, err := db.Query(getPhotosQuery, hotspotID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo string

		err := rows.Scan(
			&photo,
		)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func AddPhoto(hotspotId string, filename string) error {
	_, err := db.Exec(addPhotoQuery, hotspotId, filename)
	return err
}
