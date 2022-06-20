package usecase

import (
	"context"
	"go.uber.org/zap"
	"thourus-api/domain/entity"
	"thourus-api/gateway"
)

type SpaceUseCase struct {
	spaceGw    gateway.SpaceGw
	projectGw  gateway.ProjectGw
	documentGw gateway.DocumentGw
	logger     *zap.SugaredLogger
}

func NewSpaceUseCase(spaceGw gateway.SpaceGw, projectGw gateway.ProjectGw, documentGw gateway.DocumentGw, logger *zap.SugaredLogger) *SpaceUseCase {
	return &SpaceUseCase{
		spaceGw:    spaceGw,
		projectGw:  projectGw,
		documentGw: documentGw,
		logger:     logger,
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

func (uc *SpaceUseCase) DeleteSpace(spaceUid string) error {
	ctx := context.Background()

	// get projects in the space
	projects, err := uc.spaceGw.GetProjectsInSpace(ctx, spaceUid)
	if err != nil {
		return err
	}

	for _, project := range projects {
		// get documents in every project
		documents, err := uc.projectGw.GetDocumentsInProject(ctx, project.Uid)
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

		// delete the project after
		err = uc.projectGw.DeleteProjectByUid(ctx, project.Uid)
		if err != nil {
			return err
		}
	}
	// finally, delete the space
	err = uc.spaceGw.DeleteSpaceByUid(ctx, spaceUid)
	if err != nil {
		return err
	}
	return nil
}

func (uc *SpaceUseCase) AddSpace(companyId int, spaceName string) (entity.Space, error) {
	spaceId, err := uc.spaceGw.CreateSpace(context.Background(), spaceName, companyId)
	if err != nil {
		return entity.Space{}, err
	}
	space, err := uc.spaceGw.GetSpaceById(context.Background(), spaceId)
	if err != nil {
		return entity.Space{}, err
	}
	return space, nil
}
