package usecase

import (
	"context"
	"go.uber.org/zap"
	"thourus-api/domain/entity"
	"thourus-api/gateway"
)

type ProjectUseCase struct {
	projectGw gateway.ProjectGw
	logger    *zap.SugaredLogger
}

func NewProjectUseCase(projectGw gateway.ProjectGw, logger *zap.SugaredLogger) *ProjectUseCase {
	return &ProjectUseCase{
		projectGw: projectGw,
		logger:    logger,
	}
}

func (uc *ProjectUseCase) GetDocumentsInProject(projectUid string) ([]entity.Document, error) {
	documents, err := uc.projectGw.GetDocumentsInProject(context.Background(), projectUid)
	if err != nil {
		return []entity.Document{}, err
	}
	return documents, nil
}

func (uc *ProjectUseCase) GetProjectInfo(projectUid string) (entity.Project, error) {
	project, err := uc.projectGw.GetProjectByUid(context.Background(), projectUid)
	if err != nil {
		return entity.Project{}, err
	}
	return project, nil
}
