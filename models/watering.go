package models

import (
	"encoding/json"
	"flowers-server/config"
	"fmt"
)

//Watering struct
type Watering struct {
	ID          string `json:"id"`
	Timestamp   string `json:"timestamp"`
	Description string `json:"description"`
}

//GetAllWaterings select all waterings from database
func GetAllWaterings() ([]*Watering, error) {
	rows, err := config.DB.Query("SELECT * FROM flowers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	waterings := make([]*Watering, 0)
	for rows.Next() {
		bk := new(Watering)
		err := rows.Scan(&bk.ID, &bk.Timestamp, &bk.Description)
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

//CreateNewWatering create new watering and return it
func CreateNewWatering(bodyWatering []byte) (*Watering, error) {
	watering := Watering{}
	err := json.Unmarshal(bodyWatering, &watering)
	if err != nil {
		return nil, err
	}

	sqlStatement := `
		INSERT INTO flowers (id, Timestamp, Description)
		VALUES ($1, $2, $3)
		RETURNING id`
	id := 0

	err = config.DB.QueryRow(sqlStatement, watering.ID, watering.Timestamp, watering.Description).Scan(&id)

	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

	return &watering, nil
}
