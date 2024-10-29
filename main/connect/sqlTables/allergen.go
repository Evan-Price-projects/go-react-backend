package connect

import (
	"log"

	"github.com/Evan-Price-projects/go-react-backend/main/connect"
)

func Create_Allergen_Table() {
	db, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Db.Query(`
	Drop table if exists allergen;
	CREATE TABLE IF NOT EXISTS allergen (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
	level INT,
    deleted BOOLEAN,
    date_deleted TIMESTAMP);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Db.Exec(
		"INSERT INTO allergen (id, name, deleted, level, date_deleted) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO NOTHING",
		0, "Peanut", false, nil, nil,
	)

	if err != nil {
		log.Fatalf("Failed to insert allergen: %v", err)
	}
}
