package connect

import (
	"log"

	"github.com/Evan-Price-projects/go-react-backend/main/connect"
)

func Create_Food_Type_Table() {
	db, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Db.Query(`
CREATE TABLE IF NOT EXISTS food_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    deleted BOOLEAN DEFAULT false,
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
        WHERE conname = 'food_type_name_unique' 
        AND conrelid = 'food_type'::regclass
    ) THEN
        ALTER TABLE food_type ADD CONSTRAINT food_type_name_unique UNIQUE (name);
    END IF;
END $$;
`,
	)

	if err != nil {
		log.Fatalf("Failed to alter table: %v", err)
	}
}
