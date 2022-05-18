package migration

import (
	"gorm.io/gorm"
	"rpl-sixmath/exception"
)

func MyMigration(db *gorm.DB, tables ...interface{}) {
	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			err := db.Debug().AutoMigrate(table)
			exception.PanicIfNeeded(err)
		}
	}
}
