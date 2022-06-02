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

func (uc *MailUseCase) SendUpdates() error {
	err := uc.mailGw.SendUpdate()
	if err != nil {
		return err
	}
	return nil
}
