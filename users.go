package cdb

func Users() ([]User, error) {
	var users []User
	if result := db.Joins("Role").Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func Roles() ([]Role, error) {
	var roles []Role
	if result := db.Find(&roles); result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}
