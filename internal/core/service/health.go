package core

func (svc Service) Health() error {
	if err := svc.db.Health(); err != nil {
		return err
	}

	return nil
}
