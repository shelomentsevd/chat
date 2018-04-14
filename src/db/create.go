package db

func Create(model interface{}) error {
	if err := Pool.Create(model).Error; err != nil {
		return err
	}

	return nil
}
