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

func (m *Misc) Source(ID uint64) (*models.Source, error) {
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

func (m *Misc) User(ID uint64) (*models.User, error) {
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

func (m *Misc) Role(ID uint64) (*models.Role, error) {
	var role models.Role
	if result := m.DB.First(&role, ID); result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}
