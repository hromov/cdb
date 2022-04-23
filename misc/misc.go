package misc

import (
	"github.com/hromov/cdb"

	"gorm.io/gorm"
)

type Misc struct {
	*gorm.DB
}

func (m *Misc) Sources() ([]cdb.Source, error) {
	var sources []cdb.Source
	if result := m.DB.Find(&sources); result.Error != nil {
		return nil, result.Error
	}
	return sources, nil
}

func (m *Misc) Users() ([]cdb.User, error) {
	var users []cdb.User
	if result := m.DB.Joins("Role").Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (m *Misc) Roles() ([]cdb.Role, error) {
	var roles []cdb.Role
	if result := m.DB.Find(&roles); result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}
