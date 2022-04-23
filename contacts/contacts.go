package contacts

import (
	"database/sql"
	"time"
	"unicode"

	"github.com/hromov/cdb/misc"

	"gorm.io/gorm"
)

const fullSearch = "name LIKE @query OR second_name LIKE @query OR phone LIKE @query OR second_phone LIKE @query OR email LIKE @query OR second_email LIKE @query OR url LIKE @query OR city LIKE @query OR address LIKE @query OR position LIKE @query"
const phonesOnly = "phone LIKE @query OR second_phone LIKE @query"

type Contact struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//or company
	IsPerson   bool
	Name       string `gorm:"size:32"`
	SecondName string `gorm:"size:32"`
	//implement
	ResponsibleID *uint
	Responsible   misc.User `gorm:"foreignKey:ResponsibleID"`
	CreatedID     *uint
	Created       misc.User `gorm:"foreignKey:CreatedID"`

	Tags        []misc.Tag `gorm:"many2many:contacts_tags;"`
	Tasks       []misc.Task
	Phone       string `gorm:"size:32"`
	SecondPhone string `gorm:"size:32"`
	Email       string `gorm:"size:128"`
	SecondEmail string `gorm:"size:128"`
	URL         string `gorm:"size:128"`

	City    string `gorm:"size:128"`
	Address string `gorm:"size:256"`

	SourceID *uint8
	Source   misc.Source `gorm:"foreignKey:SourceID"`
	Position string      `gorm:"size:128"`

	Analytics misc.Analytics `gorm:"embedded;embeddedPrefix:analytics_"`
}

type ContactsResponse struct {
	Contacts []Contact
	Total    int64
}

type Contacts struct {
	DB *gorm.DB
}

func digitsOnly(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func (c *Contacts) List(limit, offset int, query string) (*ContactsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &ContactsResponse{}
	if query != "" {
		searchType := ""
		if digitsOnly(query) {
			searchType = phonesOnly
		} else {
			searchType = fullSearch
		}
		if result := c.DB.Limit(limit).Offset(offset).Where(searchType, sql.Named("query", "%"+query+"%")).Order("updated_at desc").Find(&cr.Contacts).Count(&cr.Total); result.Error != nil {
			return nil, result.Error
		}
		return cr, nil
	}

	if result := c.DB.Order("updated_at desc").Limit(limit).Offset(offset).Find(&cr.Contacts).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func (c *Contacts) ByID(ID uint64) (*Contact, error) {
	// log.Println(limit, offset, query, query == "")
	var contact Contact

	if result := c.DB.Find(&contact, ID); result.Error != nil {
		return nil, result.Error
	}
	return &contact, nil
}

// | name                   | varchar(32)      | YES  |     | NULL    |                |
// | second_name            | varchar(32)      | YES  |     | NULL    |                |
// | phone                  | varchar(32)      | YES  |     | NULL    |                |
// | second_phone           | varchar(32)      | YES  |     | NULL    |                |
// | email                  | varchar(128)     | YES  |     | NULL    |                |
// | second_email           | varchar(128)     | YES  |     | NULL    |                |
// | url                    | varchar(128)     | YES  |     | NULL    |                |
// | city                   | varchar(128)     | YES  |     | NULL    |                |
// | address                | varchar(256)     | YES  |     | NULL    |                |
// | position				| varchar(128)     | YES  |     | NULL    |                |

// func ContactsPhone(limit, offset int, query string) ([]Contact, error) {
// 	// log.Println(limit, offset, query, query == "")
// 	var contacts []Contact
// 	if query != "" {
// 		if result := db.Limit(limit).Offset(offset).Where("phone LIKE @query OR second_phone LIKE @query", sql.Named("query", "%"+query+"%")).Order("updated_at desc").Find(&contacts); result.Error != nil {
// 			return nil, result.Error
// 		}
// 		return contacts, nil
// 	}

// 	if result := db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&contacts); result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return contacts, nil
// }

// func ContactsNamesAndPhone(limit, offset int, query string) ([]Contact, error) {
// 	// log.Println(limit, offset, query, query == "")
// 	var contacts []Contact
// 	if query != "" {
// 		if result := db.Limit(limit).Offset(offset).Where("name LIKE @query OR second_name LIKE @query OR phone LIKE @query OR second_phone LIKE @query", sql.Named("query", "%"+query+"%")).Order("updated_at desc").Find(&contacts); result.Error != nil {
// 			return nil, result.Error
// 		}
// 		return contacts, nil
// 	}

// 	if result := db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&contacts); result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return contacts, nil
// }
