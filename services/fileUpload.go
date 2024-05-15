package services

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func StoreFile(userId int, db *sql.DB, filePath string, data []byte) error {
	dir := filepath.Dir(filePath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	filename := fmt.Sprintf("%s_%d.%s", filepath.Base(filePath), timestamp, filepath.Ext(filePath))
	fullPath := filepath.Join(dir, filename)

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO Files (
			file_name,
			user_id
		)
		VALUES (
			$1, $2
		)
		RETURNING id;
	`

	_, err = db.Query(query, filename, userId)
	if err != nil {
		return err
	}

	return nil
}
