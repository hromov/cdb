package leads

import (
	"github.com/hromov/cdb/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Leads struct {
	*gorm.DB
}

func (l *Leads) List(limit, offset int, query string) (*models.LeadsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &models.LeadsResponse{}
	//How to make joins work?.Joins("Contacts")
	if query != "" {
		q := l.DB.Preload(clause.Associations).Order("updated_at desc")
		q = q.Where("name LIKE ?", "%"+query+"%")
		q = q.Limit(limit).Offset(offset).Find(&cr.Leads)
		if result := q.Count(&cr.Total); result.Error != nil {
			return nil, result.Error

		}
		return cr, nil
	}
	if result := l.DB.Preload(clause.Associations).Order("updated_at desc").Limit(limit).Offset(offset).Find(&cr.Leads).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func (l *Leads) ByTag(limit, offset int, TagID uint8) (*models.LeadsResponse, error) {
	cr := &models.LeadsResponse{}

	IDs := []uint{}
	l.DB.Raw("select lead_id from leads_tags WHERE tag_id = ?", 2).Scan(&IDs)

	//How to make joins work?.Joins("Contacts")
	if result := l.DB.Preload(clause.Associations).Order("updated_at desc").
		Limit(limit).Offset(offset).Find(&cr.Leads, IDs).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func (l *Leads) ByContact(ID uint) (*models.LeadsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &models.LeadsResponse{}
	//How to make joins work?.Joins("Contacts")
	if result := l.DB.Preload(clause.Associations).Order("updated_at desc").Where("contact_id = ?", ID).Find(&cr.Leads).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func (l *Leads) ByID(ID uint64) (*models.Lead, error) {
	// log.Println(limit, offset, query, query == "")
	var lead models.Lead

	if result := l.DB.Preload(clause.Associations).First(&lead, ID); result.Error != nil {
		return nil, result.Error
	}
	return &lead, nil
}
