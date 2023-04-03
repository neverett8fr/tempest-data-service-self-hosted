package storage

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	application "tempest-data-service/pkg/application/entities"
)

func (sp *StorageProvider) GetAllFileInformation(ctx context.Context, username string) ([]application.File, error) {

	path := fmt.Sprintf("%s/%s", sp.StorageLocation, username)
	files, err := os.ReadDir(path)
	if err != nil {
		return []application.File{}, fmt.Errorf("Error reading directory %s: %v\n", path, err)
	}

	var filesOut []application.File
	for _, file := range files {
		if !file.IsDir() {
			filesOut = append(filesOut, application.File{Key: file.Name()})
		}
	}

	return filesOut, nil

}

func (sp *StorageProvider) GetFileContent(ctx context.Context, username string, key string) (application.File, error) {
	path := fmt.Sprintf("%s/%s/%s", sp.StorageLocation, username, key)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return application.File{}, fmt.Errorf("error reading file, err: %v", err)
	}

	return application.File{
		Key:  key,
		User: username,
		Data: content,
	}, nil

}

func (sp *StorageProvider) UploadSmallFile(ctx context.Context, username string, fileName string, fileSize int, fileContent []byte) error {
	path := fmt.Sprintf("%s/%s", sp.StorageLocation, username)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory, err: %v", err)
		}
	}

	filePath := fmt.Sprintf("%s/%s", path, fileName)
	err := os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return fmt.Errorf("error writing file, err: %v", err)
	}
	return nil
}
