package gateway

import (
	"context"
	"database/sql"
	"thourus-api/domain/entity"
)

type ProjectGw interface {
	GetProjectById(ctx context.Context, projectId int) (entity.Project, error)
	GetProjectByUid(ctx context.Context, projectUid string) (entity.Project, error)
	GetDocumentsInProject(ctx context.Context, projectUid string) ([]entity.Document, error)
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

func (gw *ProjectGateway) GetDocumentsInProject(ctx context.Context, projectUid string) ([]entity.Document, error) {

	const query = `
	SELECT d.id, d.uid, d.name, d.path, 
	       u.uid, u.name, u.surname,
		   st.id, st.uid, st.name,
		   p.id, p.uid, p.name,
		   s.id, s.uid, s.name,
	       d.date_created
	FROM thourus.document d
	INNER JOIN thourus.project p ON d.project_id = p.id
	INNER JOIN thourus.space s ON p.space_id = s.id
	INNER JOIN thourus.status st ON d.status_id = st.id
	INNER JOIN thourus.user u ON d.creator_id = u.id
	WHERE p.uid = ?;
`
	documents := []entity.Document{}

	rows, err := gw.db.QueryContext(ctx, query, projectUid)

	for rows.Next() {
		document := entity.Document{}
		if err = rows.Scan(
			&document.Id,
			&document.Uid,
			&document.Name,
			&document.Path,
			&document.Creator.Uid,
			&document.Creator.Name,
			&document.Creator.Surname,
			&document.Status.Id,
			&document.Status.Uid,
			&document.Status.Name,
			&document.Project.Id,
			&document.Project.Uid,
			&document.Project.Name,
			&document.Space.Id,
			&document.Space.Uid,
			&document.Space.Name,
			&document.DateCreated,
		); err != nil {
			return documents, err
		}
		document.DateCreatedString = document.DateCreated.Format(dateFormat)
		documents = append(documents, document)
	}

	return documents, nil
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

func (gw *ProjectGateway) RenameProjectById(ctx context.Context, newProjectName string, projectId int) error {

	const query = `
	UPDATE thourus.project p SET p.name = ? WHERE p.id = ?;
`
	_, err := gw.db.ExecContext(ctx, query, newProjectName, projectId)
	if err != nil {
		return err
	}

	return nil
}

func (gw *ProjectGateway) RenameProjectByUid(ctx context.Context, newProjectName string, projectUid string) error {

	const query = `
	UPDATE thourus.project p SET p.name = ? WHERE p.uid = ?;
`
	_, err := gw.db.ExecContext(ctx, query, newProjectName, projectUid)
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
