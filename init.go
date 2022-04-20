package cdb

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(dsn string) (err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect database error: %s", err.Error()))
	}

	// if table exist - do nothink, if not - create init structure with test data
	if !db.Migrator().HasTable("roles") {
		if err := db.AutoMigrate(&Role{}); err != nil {
			return err
		}
	}
	if !db.Migrator().HasTable("contacts") {
		if err := db.AutoMigrate(&Contact{}); err != nil {
			return err
		}
	}

	// for _, b := range banks_data {
	// 	db.Create(&b)
	// }
	return nil
}
