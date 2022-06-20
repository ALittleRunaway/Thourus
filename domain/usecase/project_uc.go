package usecase

import (
	"context"
	"go.uber.org/zap"
	"thourus-api/domain/entity"
	"thourus-api/gateway"
)

type ProjectUseCase struct {
	projectGw  gateway.ProjectGw
	documentGw gateway.DocumentGw
	logger     *zap.SugaredLogger
}

func NewProjectUseCase(projectGw gateway.ProjectGw, documentGw gateway.DocumentGw, logger *zap.SugaredLogger) *ProjectUseCase {
	return &ProjectUseCase{
		projectGw:  projectGw,
		documentGw: documentGw,
		logger:     logger,
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

func (uc *ProjectUseCase) DeleteProject(projectUid string) error {
	ctx := context.Background()

	documents, err := uc.projectGw.GetDocumentsInProject(ctx, projectUid)
	if err != nil {
		return err
	}

	// delete every document in the project
	for _, document := range documents {
		err = uc.documentGw.DeleteDocumentByUid(ctx, document.Uid)
		if err != nil {
			return err
		}
	}

	err = uc.projectGw.DeleteProjectByUid(ctx, projectUid)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ProjectUseCase) AddProject(projectName string, spaceId int) (entity.Project, error) {
	projectId, err := uc.projectGw.CreateProject(context.Background(), projectName, spaceId)
	if err != nil {
		return entity.Project{}, err
	}
	project, err := uc.projectGw.GetProjectById(context.Background(), projectId)
	if err != nil {
		return entity.Project{}, err
	}
	return project, nil
}
