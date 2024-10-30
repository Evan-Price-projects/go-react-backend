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
CREATE TABLE IF NOT EXISTS allergen (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    level INT,
    deleted BOOLEAN DEFAULT FALSE,
    date_deleted TIMESTAMP
);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Db.Exec(
		`
DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint 
        WHERE conname = 'allergen_name_unique' 
        AND conrelid = 'allergen'::regclass
    ) THEN
        ALTER TABLE allergen ADD CONSTRAINT allergen_name_unique UNIQUE (name);
    END IF;
END $$;
`,
	)

	if err != nil {
		log.Fatalf("Failed to alter table: %v", err)
	}
	_, err = db.Db.Exec(
		`
		INSERT INTO allergen (name, deleted, level, date_deleted) VALUES ($1, $2, $3, $4) ON CONFLICT (name) DO NOTHING`,
		"Peanut", false, nil, nil,
	)

	if err != nil {
		log.Fatalf("Failed to insert allergen: %v", err)
	}
}
