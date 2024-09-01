package core

func (c Core) Health() error {
	if err := c.db.Health(c.config.DbURL); err != nil {
		return err
	}

	return nil
}
