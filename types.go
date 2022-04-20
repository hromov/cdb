package cdb

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	IsPerson   bool
	Name       string
	SecondName string
	//implement
	ResponsibleID uint
	Responsible   User `gorm:"foreignKey:ResponsibleID"`
	CreatedID     uint
	Created       User `gorm:"foreignKey:CreatedID"`

	Tags        []Tag `gorm:"many2many:lead_tag;"`
	Tasks       []Task
	Phone       string
	SecondPhone string
	Email       string
	SecondEmail string
	URL         string

	City    string
	Address string

	SourceID uint
	Source   Source `gorm:"foreignKey:SourceID"`
	Position string

	//google analytics
	CID string
	UID string
	TID string
}

type Lead struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Budget    int
	Profit    int

	//implement
	ResponsibleID uint
	Responsible   User `gorm:"foreignKey:ResponsibleID"`
	CreatedID     uint
	Created       User `gorm:"foreignKey:CreatedID"`
	StepID        uint
	Step          Step
	//implement
	ProductID uint
	Product   Product
	//implement
	ManufacturerID uint
	Manufacturer   Manufacturer
	SourceID       uint
	Source         Source
	//google analytics
	Tags  []Tag `gorm:"many2many:lead_tag;"`
	Tasks []Task

	CID string
	UID string
	TID string

	UtmID       string
	UtmSource   string
	UtmMedium   string
	UtmCampaign string

	Domain string
}

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Email     string `gorm:"unique"`
	// Events    []Event
	RoleID uint
	Role   Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Role string `gorm:"unique"`
}

type Step struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"unique"`
	//1st, 2nd etc
	Order uint
}

type Event struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Implement
	UserID uint
	// User   User

	// And corresponding ID, one of it is not null - so we know how to show changes in contact, leas or task card
	ContactID uint
	TaskID    uint
	LeadID    uint
}

type Tag struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"unique"`
}

//Task & Notice
type Task struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeadLine  time.Time
	//if not - notice
	TaskTypeID uint
	TaskType   TaskType

	UserID    uint
	LeadID    uint
	ContactID uint

	//just links
	Files string
}

type TaskType struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"unique"`
}

type Source struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"unique"`
}

type Product struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"unique"`
}

type Manufacturer struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"unique"`
}
