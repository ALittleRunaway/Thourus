package usecase

import (
	"errors"
	"go.uber.org/zap"
	"strconv"
	"thourus-api/gateway"
)

type CacheUseCase struct {
	cacheGw gateway.CacheGw
	logger  *zap.SugaredLogger
}

func NewCacheUseCase(cacheGw gateway.CacheGw, logger *zap.SugaredLogger) *CacheUseCase {
	return &CacheUseCase{
		cacheGw: cacheGw,
		logger:  logger,
	}
}

func (uc *CacheUseCase) IncrementDocumentVersion(documentUid string) error {
	currentVersionString, err := uc.cacheGw.GetDocumentVersion(documentUid)
	if err != nil {
		return err
	}
	var currentVersion int
	if currentVersionString != "" {

		currentVersion, err = strconv.Atoi(currentVersionString)
		if err != nil {
			return err
		}

		currentVersion += 1
		err := uc.cacheGw.SaveDocumentVersion(documentUid, strconv.Itoa(currentVersion))
		if err != nil {
			return err
		}
		return nil

	} else {
		return errors.New("cannot find the document actual version")
	}
}

func (uc *CacheUseCase) SaveDocumentVersion(documentUid string, version string) error {

	err := uc.cacheGw.SaveDocumentVersion(documentUid, version)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CacheUseCase) GetDocumentVersion(documentUid string) (string, error) {

	version, err := uc.cacheGw.GetDocumentVersion(documentUid)
	if err != nil {
		return version, err
	}
	return version, nil
}
