package usecase

import (
	"go.uber.org/zap"
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

func (uc *CompanyUseCase) UploadNewDocument(projectUid string) {
	//func (uc *CompanyUseCase) UploadNewDocument(projectUid string) ([]entity.Space, error) {
	//spaces, err := uc.companyGw.GetSpacesInCompany(context.Background(), companyUid)
	//if err != nil {
	//	return nil, err
	//}
	//return spaces, nil

}
