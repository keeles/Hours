package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
)

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

func ClientExists(clientName string) (bool, error) {
	db, err := InitDb()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM clients WHERE name = ?
		)
	`, clientName).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func GetClientTasks(clientName string) (Tasks, error) {
	db, err := InitDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT t.name, t.minutes 
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
		var minutes int
		if err := rows.Scan(&name, &minutes); err != nil {
			return nil, err
		}
		tasks[name] = minutes
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
	SELECT c.name, t.name, t.minutes FROM clients c JOIN tasks t ON t.client_id = c.id;	
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[string]Tasks)
	for rows.Next() {
		var cName string
		var tName string
		var tMinutes int

		if err := rows.Scan(&cName, &tName, &tMinutes); err != nil {
			return nil, err
		}

		if clientTaskList, ok := data[cName]; ok {
			clientTaskList[tName] = tMinutes
		} else {
			task := make(Tasks)
			task[tName] = tMinutes
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

func UpdateTaskMinutes(clientName, taskName string, newMinutes float32, subtract bool) error {
	db, err := InitDb()
	if err != nil {
		return err
	}
	defer db.Close()

	var query string
	if subtract {
		query = `
			UPDATE tasks 
			SET minutes = minutes - ? 
			WHERE name = ? and client_id = (SELECT id FROM clients WHERE name = ?);
		`
	} else {
		query = `
			UPDATE tasks 
			SET minutes = minutes + ? 
			WHERE name = ? and client_id = (SELECT id FROM clients WHERE name = ?);
		`
	}

	result, err := db.Exec(query, newMinutes, taskName, clientName)
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

func GetTimer() (Timer, bool, error) {
	db, err := InitDb()
	if err != nil {
		return Timer{}, false, err
	}
	defer db.Close()

	var startTime string
	var client string
	var task sql.NullString

	err = db.QueryRow(`
		SELECT a.start_time, c.name, t.name FROM active_timer AS a LEFT JOIN clients AS c ON a.client_id = c.id LEFT JOIN tasks as t on a.task_id = t.id
	`).Scan(&startTime, &client, &task)

	if err == sql.ErrNoRows {
		return Timer{}, false, nil
	}

	if err != nil {
		return Timer{}, false, err
	}

	parsedTime, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		return Timer{}, false, err
	}

	var taskPtr *string
	if task.Valid {
		taskPtr = &task.String
	}

	data := Timer{
		ClientName: client,
		TaskName:   taskPtr,
		StartTime:  parsedTime,
	}

	return data, true, nil
}

func StartTimer(clientName, taskName string) error {
	db, err := InitDb()
	if err != nil {
		return err
	}
	defer db.Close()

	startTime := time.Now().UTC().Format(time.RFC3339)
	var query string
	var args []any

	var id int
	_ = db.QueryRow(`
		SELECT id FROM tasks 
		JOIN clients ON tasks.client_id = clients.id 
		WHERE tasks.name = ? 
		AND clients.name = ?
	`, taskName, clientName).Scan(&id)

	if taskName != "" {
		query = `
		INSERT INTO active_timer (id, client_id, task_id, start_time)
		VALUES (
			1,
			(SELECT id FROM clients WHERE name = ?),
			(SELECT tasks.id 
				FROM tasks
				JOIN clients ON tasks.client_id = clients.id
				WHERE tasks.name = ? AND clients.name = ?
			),
			?
		)
		`
		args = []any{clientName, taskName, clientName, startTime}

	} else {
		query = `
		INSERT INTO active_timer (id, client_id, task_id, start_time)
		VALUES (
			1,
			(SELECT clients.id FROM clients WHERE clients.name = ?),
			NULL,
			?
		)
		`
		args = []any{clientName, startTime}
	}

	result, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("failed to start timer: client or task not found")
	}

	return nil
}

// TODO: Cleanup
func StopTimer() (string, string, error) {
	db, err := InitDb()
	if err != nil {
		return "", "", err
	}
	defer db.Close()

	var startTimeStr string
	var taskID sql.NullInt64
	var clientID int
	var clientName string
	var taskName sql.NullString

	err = db.QueryRow(`
		SELECT timer.client_id, timer.task_id, timer.start_time, clients.name, tasks.name
		FROM active_timer as timer
		JOIN clients ON clients.id = timer.client_id
		LEFT JOIN tasks ON tasks.id = timer.task_id
		WHERE timer.id = 1
	`).Scan(&clientID, &taskID, &startTimeStr, &clientName, &taskName)

	if err != nil {
		return "", "", err
	}

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		return "", "", err
	}

	endTime := time.Now().UTC()
	duration := endTime.Sub(startTime)
	minutes := int(duration.Minutes())

	if minutes <= 0 {
		_, err = db.Exec(`DELETE FROM active_timer WHERE id = 1`)
		if err != nil {
			return "", "", err
		}

		return "", "", fmt.Errorf("Timer running for less than 1 minute, no time allocated")
	}

	var finalTaskID int
	var finalTaskName string

	if !taskID.Valid {
		finalTaskName, err = SelectTaskForClient(clientName)
		if err != nil {
			return "", "", err
		}

		err = db.QueryRow(`
			SELECT id FROM tasks
			WHERE name = ?
			AND client_id = ?
		`, finalTaskName, clientID).Scan(&finalTaskID)

		if err == sql.ErrNoRows {
			res, err := db.Exec(`
				INSERT INTO tasks (client_id, name, minutes)
				VALUES (?, ?, 0)
			`, clientID, finalTaskName)
			if err != nil {
				return "", "", err
			}

			id, err := res.LastInsertId()
			if err != nil {
				return "", "", err
			}

			finalTaskID = int(id)

		} else if err != nil {
			return "", "", err
		}
	} else {
		finalTaskID = int(taskID.Int64)
		finalTaskName = string(taskName.String)
	}

	_, err = db.Exec(`
		INSERT INTO time_entries (task_id, start_time, end_time, minutes)
		VALUES (?, ?, ?, ?)
	`,
		finalTaskID,
		startTime.Format(time.RFC3339),
		endTime.Format(time.RFC3339),
		minutes,
	)
	if err != nil {
		return "", "", err
	}

	_, err = db.Exec(`
		UPDATE tasks
		SET minutes = minutes + ?
		WHERE id = ?
	`, minutes, finalTaskID)
	if err != nil {
		return "", "", err
	}

	_, err = db.Exec(`DELETE FROM active_timer WHERE id = 1`)
	if err != nil {
		return "", "", err
	}

	return clientName, finalTaskName, nil
}

func SelectTaskForClient(clientName string) (string, error) {
	tasks, err := GetClientTasks(clientName)
	if err != nil {
		return "", err
	}
	index := -1
	taskNames := []string{}
	for name := range tasks {
		taskNames = append(taskNames, name)
	}
	prompt := promptui.SelectWithAdd{
		Label:    "Select task for time allocation",
		Items:    taskNames,
		AddLabel: "New Task",
	}
	index, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	if index == -1 {
		err = AddNewTask(clientName, result)
		if err != nil {
			return "", err
		}

		fmt.Printf("Created new task: %s \n", result)
		return result, nil
	}

	fmt.Printf("Selected task: %s \n", result)
	return result, nil
}

func SelectClientForTimer() (string, error) {
	clients, err := GetAll()
	if err != nil {
		return "", err
	}
	index := -1
	clientNames := []string{}
	for name := range clients {
		clientNames = append(clientNames, name)
	}
	prompt := promptui.SelectWithAdd{
		Label:    "Select client associated with timer",
		Items:    clientNames,
		AddLabel: "New Client",
	}
	index, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	if index == -1 {
		err = AddNewClient(result)
		if err != nil {
			return "", err
		}

		fmt.Printf("Created new client: %s \n", result)
		return result, nil
	}

	fmt.Printf("Selected Client: %s \n", result)
	return result, nil
}
