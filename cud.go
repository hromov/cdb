package cdb

func Create(i interface{}) (interface{}, error) {
	if err := db.Create(i).Error; err != nil {
		return nil, err
	}
	return i, nil
}

func Update(i interface{}) (interface{}, error) {
	if err := db.Save(i).Error; err != nil {
		return nil, err
	}
	return i, nil
}

func Delete(i interface{}) (interface{}, error) {
	if err := db.Delete(i).Error; err != nil {
		return nil, err
	}
	return i, nil
}
