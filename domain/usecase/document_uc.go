package usecase

import (
	"bytes"
	"context"
	"fmt"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"thourus-api/domain/entity"
	"thourus-api/gateway"
)

type DocumentUseCase struct {
	documentGw gateway.DocumentGw
	storageGw  gateway.StorageGw
	cryptoGw   gateway.CryptoGw
	userGw     gateway.UserGw
	ProjectGw  gateway.ProjectGw
	logger     *zap.SugaredLogger
}

func NewDocumentUseCase(documentGw gateway.DocumentGw, storageGw gateway.StorageGw, cryptoGw gateway.CryptoGw, userGw gateway.UserGw, ProjectGw gateway.ProjectGw, logger *zap.SugaredLogger) *DocumentUseCase {
	return &DocumentUseCase{
		documentGw: documentGw,
		storageGw:  storageGw,
		cryptoGw:   cryptoGw,
		userGw:     userGw,
		ProjectGw:  ProjectGw,
		logger:     logger,
	}
}

func (uc *DocumentUseCase) DownloadDocument(documentUid string) (entity.Document, error) {
	document, err := uc.documentGw.GetDocumentByUid(context.Background(), documentUid)
	if err != nil {
		return entity.Document{}, err
	}
	return document, nil
}

func (uc *DocumentUseCase) UploadNewDocument(fileHeader *multipart.FileHeader, userUid string, projectUid string) error {
	newFileName := fileHeader.Filename
	path := fmt.Sprintf("./storage/%s", newFileName)

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}

	err = uc.storageGw.SaveDocument(path, buf.Bytes())
	if err != nil {
		return err
	}

	err = uc.storageGw.ChangeRights()
	if err != nil {
		return err
	}

	PoW, err := uc.cryptoGw.GetDocumentPoW(buf.Bytes())
	if err != nil {
		return err
	}
	hash := uc.cryptoGw.GenerateHash(buf.Bytes())

	user, err := uc.userGw.GetUserByUid(context.Background(), userUid)
	if err != nil {
		return err
	}
	project, err := uc.ProjectGw.GetProjectByUid(context.Background(), projectUid)
	if err != nil {
		return err
	}

	documentId, err := uc.documentGw.CreateDocument(context.Background(), newFileName, path, user.Id, project.Id)
	if err != nil {
		return err
	}

	docHistoryId, err := uc.documentGw.AddDocumentHistory(context.Background(), documentId, hash, PoW, user.Id)
	if err != nil {
		return err
	}
	uc.logger.Debug(docHistoryId)

	byteContainer, err := ioutil.ReadAll(file) // why the long names though?
	uc.logger.Debug("size:%d", len(byteContainer))
	return nil

}

func (uc *DocumentUseCase) DeleteDocument(documentUid string) error {
	document, err := uc.documentGw.GetDocumentByUid(context.Background(), documentUid)
	if err != nil {
		return err
	}

	err = uc.storageGw.DeleteDocument(document.Path)
	if err != nil {
		return err
	}

	err = uc.documentGw.DeleteDocumentByUid(context.Background(), documentUid)
	if err != nil {
		return err
	}

	return nil
}
