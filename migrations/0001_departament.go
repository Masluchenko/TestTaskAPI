package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upDepartament, downDepartament)
}

func upDepartament(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS department (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				parent_id INT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				CONSTRAINT fk_department_parent FOREIGN KEY (parent_id) REFERENCES department(id) ON DELETE SET NULL
			);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downDepartament(tx *sql.Tx) error {
	query := `DROP TABLE departament;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
