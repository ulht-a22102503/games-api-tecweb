package rest

import "github.com/martim-lusofona/games-api/db"

func PingDB() bool {
	return db.Ping()
}
