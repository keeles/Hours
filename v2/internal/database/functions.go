package database

import (
	"database/sql"
	"fmt"
	"time"
)

func allocateTimeToTask(db *sql.DB, startTimeStr string, finalTaskID int) error {
	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		return err
	}

	endTime := time.Now().UTC()
	duration := endTime.Sub(startTime)
	minutes := int(duration.Minutes())

	if minutes <= 0 {
		_, err = db.Exec(`DELETE FROM active_timer WHERE id = 1`)
		if err != nil {
			return err
		}

		return fmt.Errorf("Timer running for less than 1 minute, no time allocated")
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
		return err
	}

	_, err = db.Exec(`
		UPDATE tasks
		SET minutes = minutes + ?
		WHERE id = ?
	`, minutes, finalTaskID)
	if err != nil {
		return err
	}

	_, err = db.Exec(`DELETE FROM active_timer WHERE id = 1`)
	if err != nil {
		return err
	}

	return nil
}

func verifyTaskId(db *sql.DB, taskID sql.NullInt64, taskName sql.NullString, clientName string, clientID int) (int, string, error) {
	var finalTaskID int
	var finalTaskName string

	if taskID.Valid {
		finalTaskID = int(taskID.Int64)
		finalTaskName = string(taskName.String)
		return finalTaskID, finalTaskName, nil
	}

	finalTaskName, err := SelectTaskForClient(clientName)
	if err != nil {
		return 0, "", err
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
			return 0, "", err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return 0, "", err
		}

		finalTaskID = int(id)

	} else if err != nil {
		return 0, "", err
	}

	return finalTaskID, finalTaskName, nil
}
