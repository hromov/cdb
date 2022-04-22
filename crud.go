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

func List(limit, offset int) (list []interface{}, err error) {
	if result := db.Limit(limit).Offset(offset).Find(&list); result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
