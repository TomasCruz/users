package database

func (pDB postgresDB) Close() {
	pDB.db.Close()
}
