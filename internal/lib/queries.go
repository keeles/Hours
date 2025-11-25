package lib

import "fmt"

func AddNewClient(name string) error {
	db, err := InitDb()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO clients (name) VALUES (?)", name)
	return err
}

func AddNewTask(clientName, name string) error {
	db, err := InitDb()
	if err != nil {
		return err
	}
	defer db.Close()

	var clientId int
	err = db.QueryRow("SELECT id FROM clients WHERE name = ?", clientName).Scan(&clientId)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO tasks (name, client_id) VALUES (?, ?)", name, clientId)
	return err
}

func GetClientTasks(clientName string) (Tasks, error) {
	db, err := InitDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT t.name, t.hours 
	FROM tasks as t 
	JOIN clients c 
	ON t.client_id = c.id 
	WHERE c.name = ?;
	`, clientName)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make(Tasks)
	for rows.Next() {
		var name string
		var hours int
		if err := rows.Scan(&name, &hours); err != nil {
			return nil, err
		}
		tasks[name] = hours
	}

	return tasks, err
}

func GetAll() (map[string]Tasks, error) {
	db, err := InitDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT c.name, t.name, t.hours FROM clients c JOIN tasks t ON t.client_id = c.id;	
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[string]Tasks)
	for rows.Next() {
		var cName string
		var tName string
		var tHours int

		if err := rows.Scan(&cName, &tName, &tHours); err != nil {
			return nil, err
		}

		if clientTaskList, ok := data[cName]; ok {
			clientTaskList[tName] = tHours
		} else {
			task := make(Tasks)
			task[tName] = tHours
			data[cName] = task
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func DeleteTask(clientName, taskName string) error {
	db, err := InitDb()
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec(`
	DELETE FROM tasks
	WHERE name = ?
	AND client_id = (SELECT id FROM clients WHERE name = ?)	
	`, taskName, clientName)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("task '%s' not found for client '%s'", taskName, clientName)
	}

	return nil
}

func DeleteClient(clientName string) error {
	db, err := InitDb()
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM clients WHERE name = ?", clientName)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("client '%s' not found", clientName)
	}

	return nil
}

func UpdateTaskHours(clientName, taskName string, newHours int, subtract bool) error {
	db, err := InitDb()
	if err != nil {
		return err
	}
	defer db.Close()

	var query string
	if subtract {
		query = `
			UPDATE tasks 
			SET hours = hours - ? 
			WHERE name = ? and client_id = (SELECT id FROM clients WHERE name = ?);
		`
	} else {
		query = `
			UPDATE tasks 
			SET hours = hours + ? 
			WHERE name = ? and client_id = (SELECT id FROM clients WHERE name = ?);
		`
	}

	result, err := db.Exec(query, newHours, taskName, clientName)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("client '%s' not found", clientName)
	}

	return nil
}
