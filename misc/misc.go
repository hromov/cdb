package misc

import "gorm.io/gorm"

type Misc struct {
	DB *gorm.DB
}

func (m *Misc) Sources() ([]Source, error) {
	var sources []Source
	if result := m.DB.Find(&sources); result.Error != nil {
		return nil, result.Error
	}
	return sources, nil
}

func (m *Misc) Users() ([]User, error) {
	var users []User
	if result := m.DB.Joins("Role").Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (m *Misc) Roles() ([]Role, error) {
	var roles []Role
	if result := m.DB.Find(&roles); result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}
