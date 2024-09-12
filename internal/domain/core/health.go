package core

func (cr Core) Health() error {
	if err := cr.db.Health(); err != nil {
		return err
	}

	return nil
}
