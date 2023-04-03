package storage

import (
	"context"
	"fmt"
	"os"
	application "tempest-data-service/pkg/application/entities"
)

func (sp *StorageProvider) GetAllFileInformation(ctx context.Context, username string) ([]application.File, error) {

	return []application.File{}, nil
}

func (sp *StorageProvider) GetFileContent(ctx context.Context, username string, key string) (application.File, error) {

	return application.File{}, nil

}

func (sp *StorageProvider) UploadSmallFile(ctx context.Context, username string, fileName string, fileExt string, fileSize int, fileContent []byte) error {
	path := fmt.Sprintf("%s/%s", sp.StorageLocation, username)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory, err: %v", err)
		}
	}

	filePath := fmt.Sprintf("%s/%s.%s", path, fileName, fileExt)
	err := os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return fmt.Errorf("error writing file, err: %v", err)
	}
	return nil
}
