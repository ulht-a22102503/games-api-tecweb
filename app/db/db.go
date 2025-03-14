package db

func Ping() bool {
	dbConnection := GetDbConnection()
	if dbConnection == nil {
		return false
	}

	db, err := dbConnection.DB()
	if err != nil {
		return false
	}

	err = db.Ping()
	return err == nil
}
