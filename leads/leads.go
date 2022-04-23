package leads

import (
	"time"

	"github.com/hromov/cdb/contacts"
	"github.com/hromov/cdb/misc"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Lead struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ClosedAt  *time.Time     `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:64"`
	Budget    uint32
	Profit    int32

	//implement
	ContactID *uint
	Contact   contacts.Contact `gorm:"foreignKey:ContactID"`

	ResponsibleID *uint
	Responsible   misc.User `gorm:"foreignKey:ResponsibleID"`
	CreatedID     *uint
	Created       misc.User `gorm:"foreignKey:CreatedID"`
	StepID        *uint8
	Step          misc.Step
	//implement
	ProductID *uint32
	Product   misc.Product
	//implement
	ManufacturerID *uint16
	Manufacturer   misc.Manufacturer
	SourceID       *uint8
	Source         *misc.Source
	//google analytics
	Tags  []misc.Tag `gorm:"many2many:leads_tags;"`
	Tasks []misc.Task

	Analytics misc.Analytics `gorm:"embedded;embeddedPrefix:analytics_"`
}

type LeadsResponse struct {
	Leads []Lead
	Total int64
}

type Leads struct {
	DB *gorm.DB
}

func (l *Leads) List(limit, offset int, query string) (*LeadsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &LeadsResponse{}
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

func (l *Leads) ByContact(ID uint) (*LeadsResponse, error) {
	// log.Println(limit, offset, query, query == "")
	cr := &LeadsResponse{}
	//How to make joins work?.Joins("Contacts")
	if result := l.DB.Preload(clause.Associations).Order("updated_at desc").Where("contact_id = ?", ID).Find(&cr.Leads).Count(&cr.Total); result.Error != nil {
		return nil, result.Error
	}
	return cr, nil
}

func (l *Leads) ByID(ID uint64) (*Lead, error) {
	// log.Println(limit, offset, query, query == "")
	var lead Lead

	if result := l.DB.Find(&lead, ID); result.Error != nil {
		return nil, result.Error
	}
	return &lead, nil
}
