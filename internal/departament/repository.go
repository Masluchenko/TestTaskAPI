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

type DepartOrEmpl interface {
	*DepartmentRequest | *EmployeeRequest
}

func (repo *DepartRepository) Create(dep *DepartmentRequest) (*DepartmentRequest, error) {
	result := repo.Database.DB.Create(dep)
	if result.Error != nil {
		return nil, result.Error
	}
	return dep, nil
}

func (repo *DepartRepository) CreateE(dep *DepartmentRequest) (*DepartmentRequest, error) {
	result := repo.Database.DB.Create(dep)
	if result.Error != nil {
		return nil, result.Error
	}
	return dep, nil
}

func CreateT[T DepartOrEmpl](repo *DepartRepository, dep T) (T, error) {
	result := repo.Database.DB.Create(dep)
	if result.Error != nil {
		return nil, result.Error
	}
	return dep, nil
}
