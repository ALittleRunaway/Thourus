package gateway

import (
	"context"
	"database/sql"
	"thourus-api/domain/entity"
)

type ProjectGw interface {
	GetProjectById(ctx context.Context, projectId int) (entity.Project, error)
	GetProjectByUid(ctx context.Context, projectUid string) (entity.Project, error)
	GetProjectsInSpace(ctx context.Context, spaceId int) ([]entity.Project, error)
	CreateProject(ctx context.Context, projectName string, spaceId int) (int, error)
	RenameProjectById(ctx context.Context, newProjectName string, spaceId int) error
	RenameProjectByUid(ctx context.Context, newProjectName string, spaceUid string) error
	DeleteProjectById(ctx context.Context, projectId int) error
	DeleteProjectByUid(ctx context.Context, projectUid string) error
	AddUserToProject(ctx context.Context, projectUid string, userUid string) error
}

type ProjectGateway struct {
	db *sql.DB
}

func NewProjectGateway(db *sql.DB) *ProjectGateway {
	return &ProjectGateway{
		db: db,
	}
}

func (gw *ProjectGateway) GetProjectById(ctx context.Context, projectId int) (entity.Project, error) {

	const query = `
	SELECT p.id, p.uid, p.name, s.id, s.uid, s.name, c.id, c.uid, c.name FROM thourus.project p
	INNER JOIN thourus.space s ON p.space_id = s.id
	INNER JOIN thourus.company c ON s.company_id = c.id
	WHERE p.id = ?;
`
	project := entity.Project{}

	rows, err := gw.db.QueryContext(ctx, query, projectId)

	for rows.Next() {
		if err = rows.Scan(
			&project.Id,
			&project.Uid,
			&project.Name,
			&project.Space.Id,
			&project.Space.Uid,
			&project.Space.Name,
			&project.Company.Id,
			&project.Company.Uid,
			&project.Company.Name,
		); err != nil {
			return project, err
		}
	}

	return project, nil
}

func (gw *ProjectGateway) GetProjectByUid(ctx context.Context, projectUid string) (entity.Project, error) {

	const query = `
	SELECT p.id, p.uid, p.name, s.id, s.uid, s.name, c.id, c.uid, c.name FROM thourus.project p
	INNER JOIN thourus.space s ON p.space_id = s.id
	INNER JOIN thourus.company c ON s.company_id = c.id
	WHERE p.uid = ?;
`
	project := entity.Project{}

	rows, err := gw.db.QueryContext(ctx, query, projectUid)

	for rows.Next() {
		if err = rows.Scan(
			&project.Id,
			&project.Uid,
			&project.Name,
			&project.Space.Id,
			&project.Space.Uid,
			&project.Space.Name,
			&project.Company.Id,
			&project.Company.Uid,
			&project.Company.Name,
		); err != nil {
			return project, err
		}
	}

	return project, nil
}

func (gw *ProjectGateway) GetProjectsInSpace(ctx context.Context, spaceId int) ([]entity.Project, error) {

	const query = `
	SELECT p.id, p.uid, p.name, s.id, s.uid, s.name, c.id, c.uid, c.name FROM thourus.project p
	INNER JOIN thourus.space s ON p.space_id = s.id
	INNER JOIN thourus.company c ON s.company_id = c.id
	WHERE s.id = ?;
`
	projects := []entity.Project{}

	rows, err := gw.db.QueryContext(ctx, query, spaceId)

	for rows.Next() {
		project := entity.Project{}
		if err = rows.Scan(
			&project.Id,
			&project.Uid,
			&project.Name,
			&project.Space.Id,
			&project.Space.Uid,
			&project.Space.Name,
			&project.Company.Id,
			&project.Company.Uid,
			&project.Company.Name,
		); err != nil {
			return projects, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (gw *ProjectGateway) CreateProject(ctx context.Context, projectName string, spaceId int) (int, error) {

	const query = `
	INSERT INTO thourus.project (name, space_id) VALUES (?, ?);
`
	var createdProjectId int

	res, err := gw.db.ExecContext(ctx, query, projectName, spaceId)
	if err != nil {
		return createdProjectId, err
	}

	createdProjectId64, err := res.LastInsertId()
	if err != nil {
		return createdProjectId, err
	}
	createdProjectId = int(createdProjectId64)

	return createdProjectId, nil
}

func (gw *ProjectGateway) RenameProjectById(ctx context.Context, newProjectName string, spaceId int) error {

	const query = `
	UPDATE thourus.project p SET p.name = ? WHERE p.id = 1;
`
	_, err := gw.db.ExecContext(ctx, query, newProjectName, spaceId)
	if err != nil {
		return err
	}

	return nil
}

func (gw *ProjectGateway) RenameProjectByUid(ctx context.Context, newProjectName string, spaceUid string) error {

	const query = `
	UPDATE thourus.project p SET p.name = ? WHERE p.uid = 1;
`
	_, err := gw.db.ExecContext(ctx, query, newProjectName, spaceUid)
	if err != nil {
		return err
	}

	return nil
}

func (gw *ProjectGateway) DeleteProjectById(ctx context.Context, projectId int) error {

	const query = `
	DELETE FROM thourus.project p WHERE p.id = ?;
`
	_, err := gw.db.ExecContext(ctx, query, projectId)
	if err != nil {
		return err
	}

	return nil
}

func (gw *ProjectGateway) DeleteProjectByUid(ctx context.Context, projectUid string) error {

	const query = `
	DELETE FROM thourus.project p WHERE p.uid = ?;
`
	_, err := gw.db.ExecContext(ctx, query, projectUid)
	if err != nil {
		return err
	}

	return nil
}

func (gw *ProjectGateway) AddUserToProject(ctx context.Context, projectUid string, userUid string) error {
	const query = `
	INSERT INTO thourus.project_user (project_id, user_id) VALUES (?, ?);
`
	_, err := gw.db.ExecContext(ctx, query, projectUid, userUid)
	if err != nil {
		return err
	}

	return nil
}
