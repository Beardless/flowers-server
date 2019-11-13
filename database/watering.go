package database

import (
	"flowers-server/config"
)

type Watering struct {
	ID          string `json:"id"`
	Timestamp   string `json:"timestamp"`
	Description string `json:"description"`
}

func GetAllWaterings() ([]*Watering, error) {
	rows, err := config.DB.Query("SELECT * FROM flowers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	waterings := make([]*Watering, 0)
	for rows.Next() {
		bk := new(Watering)
		err := rows.Scan(&bk.ID, &bk.Description, &bk.Description)
		if err != nil {
			return nil, err
		}
		waterings = append(waterings, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return waterings, nil
}
