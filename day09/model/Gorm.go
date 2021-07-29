package model

import (
	"gorm.io/gorm"
)

// Students [...]
type Students struct {
	gorm.Model
	Name string `gorm:"column:name;type:longtext" json:"name"`
	Age  int64  `gorm:"column:age;type:bigint" json:"age"`
}

// TableName get sql table name.获取数据库表名
func (m *Students) TableName() string {
	return "students"
}

// StudentsColumns get sql column name.获取数据库列名
var StudentsColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
	Age       string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Name:      "name",
	Age:       "age",
}
