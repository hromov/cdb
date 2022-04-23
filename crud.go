package cdb

func (c *CDB) Create(i interface{}) (interface{}, error) {
	if err := c.DB.Create(i).Error; err != nil {
		return nil, err
	}
	return i, nil
}

func (c *CDB) Update(i interface{}) (interface{}, error) {
	if err := c.DB.Save(i).Error; err != nil {
		return nil, err
	}
	return i, nil
}

func (c *CDB) Delete(i interface{}) (interface{}, error) {
	if err := c.DB.Delete(i).Error; err != nil {
		return nil, err
	}
	return i, nil
}

func (c *CDB) List(limit, offset int) (list []interface{}, err error) {
	if result := c.DB.Limit(limit).Offset(offset).Find(&list); result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
