package cdb

import (
	"time"

	"gorm.io/gorm"
)

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
	Responsible   User `gorm:"foreignKey:ResponsibleID"`
	CreatedID     *uint
	Created       User `gorm:"foreignKey:CreatedID"`

	Tags        []Tag `gorm:"many2many:contacts_tags;"`
	Tasks       []Task
	Phone       string `gorm:"size:32"`
	SecondPhone string `gorm:"size:32"`
	Email       string `gorm:"size:128"`
	SecondEmail string `gorm:"size:128"`
	URL         string `gorm:"size:128"`

	City    string `gorm:"size:128"`
	Address string `gorm:"size:256"`

	SourceID *uint8
	Source   Source `gorm:"foreignKey:SourceID"`
	Position string `gorm:"size:128"`

	Analytics Analytics `gorm:"embedded;embeddedPrefix:analytics_"`
}

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
	Contact   Contact `gorm:"foreignKey:ContactID"`

	ResponsibleID *uint
	Responsible   User `gorm:"foreignKey:ResponsibleID"`
	CreatedID     *uint
	Created       User `gorm:"foreignKey:CreatedID"`
	StepID        *uint8
	Step          Step
	//implement
	ProductID *uint32
	Product   Product
	//implement
	ManufacturerID *uint16
	Manufacturer   Manufacturer
	SourceID       *uint8
	Source         *Source
	//google analytics
	Tags  []Tag `gorm:"many2many:leads_tags;"`
	Tasks []Task

	Analytics Analytics `gorm:"embedded;embeddedPrefix:analytics_"`
}

type Analytics struct {
	CID string `gorm:"size:64"`
	UID string `gorm:"size:64"`
	TID string `gorm:"size:64"`

	UtmID       string `gorm:"size:64"`
	UtmSource   string `gorm:"size:64"`
	UtmMedium   string `gorm:"size:64"`
	UtmCampaign string `gorm:"size:64"`

	Domain string `gorm:"size:128"`
}

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:32"`
	Email     string         `gorm:"size:128; unique"`
	// Events    []Event
	RoleID *uint8
	Role   Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Role struct {
	ID   uint8  `gorm:"primaryKey"`
	Role string `gorm:"unique;size:32"`
}

type Step struct {
	ID        uint8 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"unique;size:32"`
	//1st, 2nd etc
	Order uint8
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

	Description string `gorm:"size:256"`
}

type Tag struct {
	ID        uint8 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:32;unique"`
}

//Task & Notice
type Task struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeadLine  time.Time
	//if not - notice
	TaskTypeID *uint8
	TaskType   TaskType

	UserID    uint
	LeadID    uint
	ContactID uint

	//just links
	Files string `gorm:"size:512"`
}

type TaskType struct {
	ID        uint8 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:32;unique"`
}

type Source struct {
	ID        uint8 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:32;unique"`
}

type Product struct {
	ID        uint32 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:64;unique"`
}

type Manufacturer struct {
	ID        uint16 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:32;unique"`
}
