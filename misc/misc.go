package misc

import (
	"github.com/hromov/cdb/models"
	"gorm.io/gorm"
)

type Misc struct {
	*gorm.DB
}

func (m *Misc) Sources() ([]models.Source, error) {
	var sources []models.Source
	if result := m.DB.Find(&sources); result.Error != nil {
		return nil, result.Error
	}
	return sources, nil
}

func (m *Misc) Source(ID uint8) (*models.Source, error) {
	var source models.Source
	if result := m.DB.First(&source, ID); result.Error != nil {
		return nil, result.Error
	}
	return &source, nil
}

func (m *Misc) Users() ([]models.User, error) {
	var users []models.User
	if result := m.DB.Joins("Role").Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (m *Misc) User(ID uint) (*models.User, error) {
	var user models.User
	if result := m.DB.Joins("Role").First(&user, ID); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (m *Misc) Roles() ([]models.Role, error) {
	var roles []models.Role
	if result := m.DB.Find(&roles); result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (m *Misc) Role(ID uint8) (*models.Role, error) {
	var role models.Role
	if result := m.DB.First(&role, ID); result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

func (m *Misc) Products() ([]models.Product, error) {
	var items []models.Product
	if result := m.DB.Find(&items); result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (m *Misc) Product(ID uint32) (*models.Product, error) {
	var item models.Product
	if result := m.DB.First(&item, ID); result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (m *Misc) Manufacturers() ([]models.Manufacturer, error) {
	var items []models.Manufacturer
	if result := m.DB.Find(&items); result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (m *Misc) Manufacturer(ID uint16) (*models.Manufacturer, error) {
	var item models.Manufacturer
	if result := m.DB.First(&item, ID); result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (m *Misc) Steps() ([]models.Step, error) {
	var items []models.Step
	if result := m.DB.Find(&items); result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (m *Misc) Step(ID uint8) (*models.Step, error) {
	var item models.Step
	if result := m.DB.First(&item, ID); result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}
