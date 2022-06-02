package gateway

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type StorageGw interface {
	SaveDocument(path string, bytes []byte) error
}

type StorageGateway struct {
	storagePath string
}

func NewStorageGateway(storagePath string) *StorageGateway {
	return &StorageGateway{
		storagePath: storagePath,
	}
}

func (gw *StorageGateway) SaveDocument(path string, bytes []byte) error {

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := ioutil.WriteFile(path, bytes, 0)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (gw *StorageGateway) DownloadDocument() error {
	return nil
}

func (gw *StorageGateway) DeleteDocument(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
