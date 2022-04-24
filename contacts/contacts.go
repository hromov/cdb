package contacts

import (
	"database/sql"
	"unicode"

	"github.com/hromov/cdb/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const fullSearch = "name LIKE @query OR second_name LIKE @query OR phone LIKE @query OR second_phone LIKE @query OR email LIKE @query OR second_email LIKE @query OR url LIKE @query OR city LIKE @query OR address LIKE @query OR position LIKE @query"
const phonesOnly = "phone LIKE @query OR second_phone LIKE @query"

type Contacts struct {
	*gorm.DB
}

func digitsOnly(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func (c *Contacts) List(limit, offset int, query string) (*models.ContactsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &models.ContactsResponse{}
	if query != "" {
		searchType := ""
		if digitsOnly(query) {
			searchType = phonesOnly
		} else {
			searchType = fullSearch
		}
		if result := c.DB.Preload(clause.Associations).Limit(limit).Offset(offset).Where(searchType, sql.Named("query", "%"+query+"%")).Order("updated_at desc").Find(&cr.Contacts).Count(&cr.Total); result.Error != nil {
			return nil, result.Error
		}
		return cr, nil
	}

	if result := c.DB.Preload(clause.Associations).Order("updated_at desc").Limit(limit).Offset(offset).Find(&cr.Contacts).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func (c *Contacts) ByID(ID uint64) (*models.Contact, error) {
	// log.Println(limit, offset, query, query == "")
	var contact models.Contact

	if result := c.DB.Preload(clause.Associations).First(&contact, ID); result.Error != nil {
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
