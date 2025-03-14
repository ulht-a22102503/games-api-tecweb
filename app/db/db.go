package db

func Ping() bool {
	db, err := DB_CONNECTION.DB()
	if err != nil {
		return false
	}

	err = db.Ping()
	return err == nil
}
