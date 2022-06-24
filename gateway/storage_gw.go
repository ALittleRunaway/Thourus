package gateway

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type StorageGw interface {
	SaveDocument(path string, bytes []byte) error
	ChangeRights() error
	DeleteDocument(path string) error
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

func (gw *StorageGateway) ChangeRights() error {

	cmd := exec.Command("chmod", "-R", "777", "./storage")
	_, err := cmd.Output()

	if err != nil {
		return err
	}
	return nil
}

func (gw *StorageGateway) DeleteDocument(path string) error {
	err := os.Remove(path)

	if err != nil {
		return err
	}

	return nil
}
