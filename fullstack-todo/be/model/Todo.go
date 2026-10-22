package model

import "example/fullstack-be/db"

type Todo struct {
	ID          int64  `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
}

func (e *Todo) Save() error {
	query := `
	insert into todos (
		title,
		description
	) values 
	 (?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(
		e.Title,
		e.Description,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}
	e.ID = id
	return nil
}

func GetAll() ([]Todo, error) {
	query := ` 
	select id, title, description from todos
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {

		var todo Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
