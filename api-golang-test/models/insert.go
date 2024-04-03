package models

import "api-golang-test/db"

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	sql := `INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Completed).Scan(&id)

	defer conn.Close()
	return
}
