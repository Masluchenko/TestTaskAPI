package departament

import "TestTaskAPI/pkg/db"

func NewDepartRepository(database *db.Db) *DepartRepository {
	return &DepartRepository{
		Database: database,
	}
}

type DepartRepository struct {
	Database *db.Db
}

func (repo *DepartRepository) CreateDepart(dep *DepartmentRequest) (*DepartmentRequest, error) {
	result := repo.Database.DB.Create(dep)
	if result.Error != nil {
		return nil, result.Error
	}
	return dep, nil
}

func (repo *DepartRepository) CreateEmpl(dep *EmployeeRequest) (*EmployeeRequest, error) {
	result := repo.Database.DB.Create(dep)
	if result.Error != nil {
		return nil, result.Error
	}
	return dep, nil
}
