package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	dbName string
)

const (
	DataCenter dbName = "datacenter"
)

var _dbMap = make(map[dbName]*gorm.DB, 0)

func Init() {
	db, err := gorm.Open(mysql.Open("root:a335577357@localhost/datacenter"), &gorm.Config{})
	if err != nil {
		return
	}

	_dbMap[DataCenter] = db
}

func GetDB(name dbName) *gorm.DB {
	return _dbMap[name]
}
