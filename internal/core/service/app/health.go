package app

func (svc AppUserService) Health() error {
	if err := svc.db.Health(); err != nil {
		return err
	}

	return nil
}
