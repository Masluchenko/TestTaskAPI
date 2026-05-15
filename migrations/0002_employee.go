package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upEmployee, downEmployee)
}

func upEmployee(tx *sql.Tx) error {
	query := `CREATE TABLE Employee (
				id SERIAL PRIMARY KEY,
				department_id INT NOT NULL,
				full_name VARCHAR(255) NOT NULL,
				position VARCHAR(255) NOT NULL,
				hired_at DATE NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				CONSTRAINT fk_employee_department FOREIGN KEY (department_id) REFERENCES Department(id) ON DELETE CASCADE
			);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downEmployee(tx *sql.Tx) error {
	query := `DROP TABLE employee;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
