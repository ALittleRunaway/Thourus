package usecase

import (
	"context"
	"go.uber.org/zap"
	"thourus-api/domain/entity"
	"thourus-api/gateway"
)

type UserUseCase struct {
	userGw gateway.UserGw
	logger *zap.SugaredLogger
}

func NewUserUseCase(userGw gateway.UserGw, logger *zap.SugaredLogger) *UserUseCase {
	return &UserUseCase{
		userGw: userGw,
		logger: logger,
	}
}

func (uc *UserUseCase) LoginUser(userPassword string, userLogin string) (entity.User, error) {

	user, err := uc.userGw.GetUserByPasswordAndLogin(context.Background(), userPassword, userLogin)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
