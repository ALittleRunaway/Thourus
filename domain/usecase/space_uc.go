package usecase

import (
	"context"
	"go.uber.org/zap"
	"thourus-api/domain/entity"
	"thourus-api/gateway"
)

type SpaceUseCase struct {
	spaceGw gateway.SpaceGw
	logger  *zap.SugaredLogger
}

func NewSpaceUseCase(spaceGw gateway.SpaceGw, logger *zap.SugaredLogger) *SpaceUseCase {
	return &SpaceUseCase{
		spaceGw: spaceGw,
		logger:  logger,
	}
}

func (uc *SpaceUseCase) GetProjectsInSpace(spaceUid string) ([]entity.Project, error) {
	projects, err := uc.spaceGw.GetProjectsInSpace(context.Background(), spaceUid)
	if err != nil {
		return []entity.Project{}, err
	}
	return projects, nil
}

func (uc *SpaceUseCase) GetSpaceInfo(spaceUid string) (entity.Space, error) {
	space, err := uc.spaceGw.GetSpaceByUid(context.Background(), spaceUid)
	if err != nil {
		return entity.Space{}, err
	}
	return space, nil
}
