package db

type User struct {
	PersonId int    `json:"person_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

const getUserQuery = "SELECT PersonId, Login, Password FROM PhotoApp.dbo.Users"
const addUserQuery = "INSERT INTO PhotoApp.dbo.Users(Login, Password) OUTPUT Inserted.PersonId VALUES (@p1, @p2)"

func GetUsers() ([]User, error) {
	var users []User

	rows, err := db.Query(getUserQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.PersonId,
			&user.Login,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func AddUser(user User) (int, error) {
	var id int
	err := db.QueryRow(addUserQuery, user.Login, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
