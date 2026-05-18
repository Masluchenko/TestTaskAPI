package departament

import (
	"time"
)

type DepartmentRequest struct {
	ID        uint                `gorm:"primaryKey" json:"id"`
	Name      string              `gorm:"not null" json:"name" binding:"required"`
	ParentID  *uint               `json:"parent_id,omitempty"`
	Parent    *DepartmentRequest  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children  []DepartmentRequest `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	CreatedAt time.Time           `json:"created_at"`
}

type EmployeeRequest struct {
	ID           uint              `gorm:"primaryKey" json:"id"`
	DepartmentID uint              `gorm:"not null" json:"department_id"`
	Department   DepartmentRequest `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	FullName     string            `gorm:"not null" json:"full_name" binding:"required"`
	Position     string            `gorm:"not null" json:"position" binding:"required"`
	HiredAt      *time.Time        `json:"hired_at,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
}

func (DepartmentRequest) TableName() string {
	return "department"
}

func (EmployeeRequest) TableName() string {
	return "employee"
}
