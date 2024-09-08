package database

func (pDB postgresDB) Health() error {
	_, err := openAndCheck(pDB.config.DBURL)
	return err
}
