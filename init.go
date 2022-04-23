package cdb

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type CDB struct {
// 	db *gorm.DB
// }
type CDB struct {
	DB *gorm.DB
}

// var db *gorm.DB
// const dsn = "root:password@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

func Init(dsn string) (*CDB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to connect database error: %s", err.Error()))
	}

	// if table exist - do nothink, if not - create init structure with test data
	if !db.Migrator().HasTable("roles") {
		if err := db.AutoMigrate(&Role{}); err != nil {
			return nil, err
		}
	}
	if !db.Migrator().HasTable("contacts") {
		if err := db.AutoMigrate(&Contact{}); err != nil {
			return nil, err
		}
	}

	if !db.Migrator().HasTable("leads") {
		if err := db.AutoMigrate(&Lead{}); err != nil {
			return nil, err
		}
	}

	if !db.Migrator().HasTable("tasks") {
		if err := db.AutoMigrate(&Task{}); err != nil {
			return nil, err
		}
	}

	// var lead Lead
	// log.Println(db.Model(&lead).Association("Contacts"))
	// // `user` is the source model, it must contains primary key
	// // `Languages` is a relationship's field name
	// // If the above two requirements matched, the AssociationMode should be started successfully, or it should return error
	// log.Println(db.Model(&lead).Association("Contacts").Error)

	// for _, b := range banks_data {
	// 	db.Create(&b)
	// }
	return &CDB{DB: db}, nil
}
