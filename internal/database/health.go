package database

func (pDB postgresDB) Health(dbString string) error {
	_, err := openAndCheck(dbString)
	return err
}
