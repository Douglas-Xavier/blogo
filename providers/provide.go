package providers

import (
	"yakki/blogo/api"
	"yakki/blogo/data"
)

func InitUserApi() api.UserInterface {
	ds := dbProvider()
	userRepoImpl := data.ProvideUserRepository(ds)
	userApiImpl := api.ProvideUserApi(userRepoImpl)
	return userApiImpl
}

func dbProvider() data.Datasource {
	return data.GetDBInstance()
}
