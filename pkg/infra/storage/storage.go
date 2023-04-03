package storage

import (
	"context"
	application "tempest-data-service/pkg/application/entities"
)

func (sp *StorageProvider) GetAllFileInformation(ctx context.Context, username string) ([]application.File, error) {

	return []application.File{}, nil
}

func (sp *StorageProvider) GetFileContent(ctx context.Context, username string, key string) (application.File, error) {

	return application.File{}, nil

}

func (sp *StorageProvider) UploadSmallFile(ctx context.Context, username string, fileName string, fileExt string, fileSize int, fileContent []byte) error {

	return nil

}
