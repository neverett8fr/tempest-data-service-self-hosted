package storage

import (
	"context"
)

type StorageProvider struct {
	StorageLocation string
}

func InitialiseStorageProvider(ctx context.Context, storageLocation string) (StorageProvider, error) {

	return StorageProvider{
		StorageLocation: storageLocation,
	}, nil
}
