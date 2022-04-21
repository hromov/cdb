package cdb

func Contacts() ([]Contact, error) {
	var contacts []Contact
	if result := db.Find(&contacts); result.Error != nil {
		return nil, result.Error
	}
	return contacts, nil
}
