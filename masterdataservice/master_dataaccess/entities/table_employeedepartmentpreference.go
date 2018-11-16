package entities

import "github.com/jinzhu/gorm"

type TableEmployeeDepartmentPreference struct {
	gorm.Model
	EmployeeId      uint            `gorm:"column:employeeid;not_null"`
	DepartmentId    uint            `gorm:"column:departmentid;not_null"`
	PreferenceOrder int             `gorm:"column:PreferenceOrder;not_null"`
	TableEmployee   TableEmployee   `gorm:"foreignkey:employeeid"`
	TableDepartment TableDepartment `gorm:"foreignkey:departmentid"`
}

func (c TableEmployeeDepartmentPreference) TableName() string {
	return "table_empdepartmentpreference"
}
