package core

func (c Core) Health() error {
	if err := c.db.Health(); err != nil {
		return err
	}

	return nil
}
