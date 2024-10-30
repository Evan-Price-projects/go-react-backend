package connect

import (
	"log"

	"github.com/Evan-Price-projects/go-react-backend/main/connect"
)

func Create_Ingredient_Table() {
	db, err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Db.Query(`
CREATE TABLE IF NOT EXISTS ingredient (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    food_types INTEGER[],
    allergens INTEGER[],
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
        WHERE conname = 'ingredient_name_unique' 
        AND conrelid = 'ingredient'::regclass
    ) THEN
        ALTER TABLE ingredient ADD CONSTRAINT ingredient_name_unique UNIQUE (name);
    END IF;
END $$;
`,
	)

	if err != nil {
		log.Fatalf("Failed to alter table: %v", err)
	}
}
