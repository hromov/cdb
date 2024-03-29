package cdb

import (
	"errors"
	"fmt"

	"github.com/hromov/cdb/contacts"
	"github.com/hromov/cdb/leads"
	"github.com/hromov/cdb/misc"
	"github.com/hromov/cdb/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsnForTests = "root:password@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

type CDB struct {
	*gorm.DB
}

//test part
func (db *CDB) Contacts() *contacts.Contacts {
	return &contacts.Contacts{DB: db.DB}
}

func (db *CDB) Leads() *leads.Leads {
	return &leads.Leads{DB: db.DB}
}

func (db *CDB) Misc() *misc.Misc {
	return &misc.Misc{DB: db.DB}
}

func Init(dsn string) (*CDB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to connect database error: %s", err.Error()))
	}

	// if table exist - do nothink, if not - create init structure with test data
	if !db.Migrator().HasTable("roles") {
		if err := db.AutoMigrate(&models.Role{}); err != nil {
			return nil, err
		}
	}
	if !db.Migrator().HasTable("contacts") {
		if err := db.AutoMigrate(&models.Contact{}); err != nil {
			return nil, err
		}
	}

	if !db.Migrator().HasTable("leads") {
		if err := db.AutoMigrate(&models.Lead{}); err != nil {
			return nil, err
		}
	}

	if !db.Migrator().HasTable("tasks") {
		if err := db.AutoMigrate(&models.Task{}); err != nil {
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
