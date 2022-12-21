package models

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Age        int    `json:"age"`
}

func NewUser() User {
	return User{}
}

func PostUser(user User) error {
	sql := "INSERT INTO users(username, first_name, last_name, age) VALUES (?,?,?,?)"
	_, err := db.Exec(sql, user.Username, user.First_name, user.Last_name, user.Age)

	return err
}

func PutUser(user User, id int) error {
	sql := "UPDATE users SET username=?, first_name=?, last_name=?, age=? WHERE ID=?"
	_, err := db.Exec(sql, user.Username, user.First_name, user.Last_name, user.Age, id)

	return err
}

func DeleteUser(id int) error {
	sql := "DELETE FROM users WHERE ID=?"
	_, err := db.Exec(sql, id)

	return err
}

func GetUsers() []User {
	users := []User{}

	sql := "SELECT * FROM users"
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := NewUser()
		rows.Scan(&user.ID, &user.Username, &user.First_name, &user.Last_name, &user.Age)
		users = append(users, user)
	}

	return users
}

func GetUser(id int) User {
	user := NewUser()
	sql := "SELECT * FROM users WHERE ID=?"
	rows, _ := db.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.First_name, &user.Last_name, &user.Age)
	}

	return user
}
