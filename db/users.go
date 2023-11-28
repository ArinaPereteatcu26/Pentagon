package db

import "errors"

type User struct {
	PersonId int    `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Hash     string `json:"-"`
}

const getUserQuery = "SELECT PersonId, Login, Hash FROM PhotoApp.dbo.Users WHERE Login = @p1"
const addUserQuery = "INSERT INTO PhotoApp.dbo.Users(Login, Hash) OUTPUT Inserted.PersonId VALUES (@p1, @p2)"

func GetUser(login string) (*User, error) {
	rows, err := db.Query(getUserQuery, login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("no such user")
	}

	var user User
	err = rows.Scan(
		&user.PersonId,
		&user.Login,
		&user.Hash,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func AddUser(user User) (int, error) {
	var id int

	err := db.QueryRow(addUserQuery, user.Login, user.Hash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
