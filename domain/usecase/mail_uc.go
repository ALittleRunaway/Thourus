package usecase

import (
	"go.uber.org/zap"
	"thourus-api/gateway"
)

type MailUseCase struct {
	mailGw gateway.MailGw
	logger *zap.SugaredLogger
}

func NewMailUseCase(mailGw gateway.MailGw, logger *zap.SugaredLogger) *MailUseCase {
	return &MailUseCase{
		mailGw: mailGw,
		logger: logger,
	}
}

func (uc *MailUseCase) SendUpdates(filename string) error {
	err := uc.mailGw.SendUpdate(filename)
	if err != nil {
		return err
	}
	return nil
}
