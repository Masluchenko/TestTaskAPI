package departament

import "TestTaskAPI/db"

type DepartRepository struct {
	Database *db.Db
}

func NewDepartRepository(database *db.Db) *DepartRepository {
	return &DepartRepository{
		Database: database,
	}
}

type DepatType interface {
	*DepartmentRequest
	*EmployeeRequest
}

func (repo *DepartRepository) Create[T DepatType](table *T) (*T, error) {
	result := repo.Database.DB.Create(table)
	if result.Error != nil {
		return nil, result.Error
	}
	return table, nil
}
