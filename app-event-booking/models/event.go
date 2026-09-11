package models

import (
	"fmt"
	"time"

	"example.com/app-event-booking/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:required`
	Description string    `binding:required`
	Location    string    `binding:required`
	Date        time.Time `binding:required`
	UserID      int
}

var events []Event = []Event{}

// | Cách                        | Dùng để                         | Ví dụ                                          |
// | --------------------------- | ------------------------------- | ---------------------------------------------- |
// | `DB.Exec()`                 | Thay đổi database/dữ liệu       | `CREATE`, `INSERT`, `UPDATE`, `DELETE`         |
// | `Prepare()` + `stmt.Exec()` | Chuẩn bị SQL trước rồi thực thi | Thường hữu ích khi chạy **cùng SQL nhiều lần** |
// | `DB.Query()`                | Lấy dữ liệu                     | `SELECT`                                       |

func (e *Event) Save() error {
	query := `
        INSERT INTO events (
            name,
            description,
            location,
            dateTime,
            user_id
        )
        VALUES (?, ?, ?, ?, ?)
    `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		e.Name,
		e.Description,
		e.Location,
		e.Date,
		e.UserID,
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
func GetAllEvents() ([]Event, error) {
	query := `
        SELECT id, name, description, location, dateTime, user_id
        FROM events
    `

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.Date,
			&event.UserID,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	fmt.Println(row)
	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.Date,
		&event.UserID,
	)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func UpdateByID(id int64, event Event) error {
	query := `
		UPDATE events
		SET name = ? , description = ? , location = ?, dateTime = ?
		WHERE id  = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.Date,
	)
	return err
}
