package dbClient

import (
	"projects/ProjectOneAPI_3/configs"
	"projects/ProjectOneAPI_3/repositories"
)

func Client(dbName, cName string) repositories.Client {
	dbClient := configs.ConnectionMDB(dbName, cName)
	return repositories.Client{Collection: dbClient}
}
