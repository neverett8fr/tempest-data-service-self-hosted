package storage

import (
	"context"
	"fmt"
)

type StorageProvider struct {
	APIURLBase   string
	ProjectName  string
	DatabaseName string
	Key          string
}

func InitialiseStorageProvider(ctx context.Context, APIURLBase string, projectName string, databaseName string, key string) (StorageProvider, error) {
	if projectName == "" || databaseName == "" {
		return StorageProvider{}, fmt.Errorf("project name or database name were empty")
	}

	return StorageProvider{
		Key:          key,
		APIURLBase:   APIURLBase,
		ProjectName:  projectName,
		DatabaseName: databaseName,
	}, nil
}
