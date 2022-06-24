package usecase

import (
	"bytes"
	"context"
	"fmt"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strings"
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

func (uc *DocumentUseCase) UploadNewDocument(fileHeader *multipart.FileHeader, userUid string, projectUid string) (entity.Document, error) {
	newFileName := fileHeader.Filename
	path := fmt.Sprintf("./storage/%s", newFileName)
	ctx := context.Background()

	file, err := fileHeader.Open()
	if err != nil {
		return entity.Document{}, err
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return entity.Document{}, err
	}

	err = uc.storageGw.SaveDocument(path, buf.Bytes())
	if err != nil {
		return entity.Document{}, err
	}

	err = uc.storageGw.ChangeRights()
	if err != nil {
		return entity.Document{}, err
	}

	PoW, err := uc.cryptoGw.GetDocumentPoW(buf.Bytes())
	if err != nil {
		return entity.Document{}, err
	}
	hash := uc.cryptoGw.GenerateHash(buf.Bytes())

	user, err := uc.userGw.GetUserByUid(ctx, userUid)
	if err != nil {
		return entity.Document{}, err
	}

	project, err := uc.ProjectGw.GetProjectByUid(ctx, projectUid)
	if err != nil {
		return entity.Document{}, err
	}

	documentId, err := uc.documentGw.CreateDocument(ctx, newFileName, path, user.Id, project.Id)
	if err != nil {
		return entity.Document{}, err
	}

	document, err := uc.documentGw.GetDocumentById(ctx, documentId)
	if err != nil {
		return entity.Document{}, err
	}

	docHistoryId, err := uc.documentGw.AddDocumentHistory(ctx, documentId, hash, PoW, user.Id)
	if err != nil {
		return entity.Document{}, err
	}
	uc.logger.Debug(docHistoryId)

	byteContainer, err := ioutil.ReadAll(file) // why the long names though?
	uc.logger.Debug("size:%d", len(byteContainer))
	return document, nil

}

func (uc *DocumentUseCase) UpdateDocument(fileHeader *multipart.FileHeader, userUid string, documentUid string, version string) error {
	document, err := uc.documentGw.GetDocumentByUid(context.Background(), documentUid)
	if err != nil {
		return err
	}
	nameExtention := strings.Split(document.Name, ".")
	path := fmt.Sprintf("./storage/%s-v%v.%s", nameExtention[0], version, nameExtention[1])

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

	err = uc.documentGw.ChangeDocumentStatusToPending(context.Background(), document.Id)
	if err != nil {
		return err
	}

	docHistoryId, err := uc.documentGw.AddDocumentHistory(context.Background(), document.Id, hash, PoW, user.Id)
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

func (uc *DocumentUseCase) GetDocumentHistory(documentUid string) ([]entity.DocumentHistory, error) {
	document, err := uc.documentGw.GetDocumentByUid(context.Background(), documentUid)
	if err != nil {
		return []entity.DocumentHistory{}, err
	}

	historyRows, err := uc.documentGw.GetDocumentHistory(context.Background(), document.Id)
	if err != nil {
		return []entity.DocumentHistory{}, err
	}
	return historyRows, nil
}
