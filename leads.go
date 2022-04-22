package cdb

import "gorm.io/gorm/clause"

type LeadsResponse struct {
	Leads []Lead
	Total int64
}

func Leads(limit, offset int, query string) (*LeadsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &LeadsResponse{}
	//How to make joins work?.Joins("Contacts")
	if query != "" {
		if result := db.Preload(clause.Associations).Order("updated_at desc").Where("name LIKE ?", "%"+query+"%").
			Limit(limit).Offset(offset).Find(&cr.Leads).Count(&cr.Total); result.Error != nil {
			return nil, result.Error

		}
		return cr, nil
	}
	if result := db.Preload(clause.Associations).Order("updated_at desc").Limit(limit).Offset(offset).Find(&cr.Leads).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func LeadsByContact(ID uint) (*LeadsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &LeadsResponse{}
	//How to make joins work?.Joins("Contacts")
	if result := db.Preload(clause.Associations).Order("updated_at desc").Where("contact_id = ?", ID).Find(&cr.Leads).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func LeadByID(ID uint64) (*Lead, error) {
	// log.Println(limit, offset, query, query == "")
	var lead Lead

	if result := db.Find(&lead, ID); result.Error != nil {
		return nil, result.Error
	}
	return &lead, nil
}
