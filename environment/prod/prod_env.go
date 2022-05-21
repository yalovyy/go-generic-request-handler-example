package prod

import (
	"ghandler/clients/storage"
	"ghandler/environment"
)

type ProdEnvironment struct {
	storage *storage.Storage
}

func NewProdEnvironment() *ProdEnvironment {
	return &ProdEnvironment{storage.New(readFromEnv())}
}

func readFromEnv() string {
	return "URL"
}

func (env *ProdEnvironment) GetStorage() environment.Storage {
	return env.storage
}
