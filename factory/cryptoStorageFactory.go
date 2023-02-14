package factory

import (
	"Golang-assignment-srikomm/store"
	"context"
)

func CryptoStorageFactory(ctx context.Context) (store.CryptoStorageInterface, error) {
	var client, err = store.NewRedisClient(ctx)
	if err != nil {
		return nil, err
	}
	return store.NewCryptoCacheStorage(client), nil
}
