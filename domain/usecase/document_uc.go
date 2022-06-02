package usecase

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"thourus-api/gateway"
)

type DocumentUseCase struct {
	documentGw gateway.DocumentGw
	storageGw  gateway.StorageGw
	cryptoGw   gateway.CryptoGw
	logger     *zap.SugaredLogger
}

func NewDocumentUseCase(documentGw gateway.DocumentGw, storageGw gateway.StorageGw, cryptoGw gateway.CryptoGw, logger *zap.SugaredLogger) *DocumentUseCase {
	return &DocumentUseCase{
		documentGw: documentGw,
		storageGw:  storageGw,
		cryptoGw:   cryptoGw,
		logger:     logger,
	}
}

func (uc *DocumentUseCase) UploadNewDocument(fileHeader *multipart.FileHeader) error {
	newFileName := fileHeader.Filename

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}

	err = uc.storageGw.SaveDocument(fmt.Sprintf("./storage/%s", newFileName), buf.Bytes())
	if err != nil {
		return err
	}

	byteContainer, err := ioutil.ReadAll(file) // why the long names though?
	fmt.Printf("size:%d", len(byteContainer))
	return nil

}
