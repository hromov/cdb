package cdb

func Sources() ([]Source, error) {
	var sources []Source
	if result := db.Find(&sources); result.Error != nil {
		return nil, result.Error
	}
	return sources, nil
}
