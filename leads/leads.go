package leads

import (
	"github.com/hromov/cdb"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Leads struct {
	*gorm.DB
}

func (l *Leads) List(limit, offset int, query string) (*cdb.LeadsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &cdb.LeadsResponse{}
	//How to make joins work?.Joins("Contacts")
	if query != "" {
		if result := l.DB.Preload(clause.Associations).Order("updated_at desc").Where("name LIKE ?", "%"+query+"%").
			Limit(limit).Offset(offset).Find(&cr.Leads).Count(&cr.Total); result.Error != nil {
			return nil, result.Error

		}
		return cr, nil
	}
	if result := l.DB.Preload(clause.Associations).Order("updated_at desc").Limit(limit).Offset(offset).Find(&cr.Leads).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func (l *Leads) ByContact(ID uint) (*cdb.LeadsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &cdb.LeadsResponse{}
	//How to make joins work?.Joins("Contacts")
	if result := l.DB.Preload(clause.Associations).Order("updated_at desc").Where("contact_id = ?", ID).Find(&cr.Leads).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func (l *Leads) ByID(ID uint64) (*cdb.Lead, error) {
	// log.Println(limit, offset, query, query == "")
	var lead cdb.Lead

	if result := l.DB.Find(&lead, ID); result.Error != nil {
		return nil, result.Error
	}
	return &lead, nil
}
