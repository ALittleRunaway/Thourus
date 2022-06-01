package gateway

import (
	"context"
	"database/sql"
	"thourus-api/domain/entity"
)

type UserGw interface {
}

type UserGateway struct {
	db *sql.DB
}

func NewUserGateway(db *sql.DB) *UserGateway {
	return &UserGateway{
		db: db,
	}
}

func (gw *UserGateway) GetUserById(ctx context.Context, documentId int) (entity.Document, error) {

	const query = `
	SELECT d.id, d.uid, d.name, d.path, u.uid,
		   p.id, p.uid, p.name,
		   s.id, s.uid, s.name
	FROM thourus.document d
	INNER JOIN thourus.project p ON d.project_id = p.id
	INNER JOIN thourus.space s ON p.space_id = s.id
	INNER JOIN thourus.user u ON d.creator_id = u.id
	WHERE d.id = ?;
`
	document := entity.Document{}

	rows, err := gw.db.QueryContext(ctx, query, documentId)

	for rows.Next() {
		if err = rows.Scan(
			&document.Id,
			&document.Uid,
			&document.Name,
			&document.Path,
			&document.Creator.Uid,
			&document.Project.Id,
			&document.Project.Uid,
			&document.Project.Name,
			&document.Space.Id,
			&document.Space.Uid,
			&document.Space.Name,
		); err != nil {
			return document, err
		}
	}

	return document, nil
}

func (gw *UserGateway) GetUserByUid(ctx context.Context, documentUid string) (entity.Document, error) {

	const query = `
	SELECT d.id, d.uid, d.name, d.path, u.uid,
		   p.id, p.uid, p.name,
		   s.id, s.uid, s.name
	FROM thourus.document d
	INNER JOIN thourus.project p ON d.project_id = p.id
	INNER JOIN thourus.space s ON p.space_id = s.id
	INNER JOIN thourus.user u ON d.creator_id = u.id
	WHERE d.uid = ?;
`
	document := entity.Document{}

	rows, err := gw.db.QueryContext(ctx, query, documentUid)

	for rows.Next() {
		if err = rows.Scan(
			&document.Id,
			&document.Uid,
			&document.Name,
			&document.Path,
			&document.Creator.Uid,
			&document.Project.Id,
			&document.Project.Uid,
			&document.Project.Name,
			&document.Space.Id,
			&document.Space.Uid,
			&document.Space.Name,
		); err != nil {
			return document, err
		}
	}

	return document, nil
}

func (gw *UserGateway) GetUsersInProjectById(ctx context.Context, projectId int) ([]entity.Document, error) {

	const query = `
	SELECT d.id, d.uid, d.name, d.path, u.uid,
		   p.id, p.uid, p.name,
		   s.id, s.uid, s.name
	FROM thourus.document d
	INNER JOIN thourus.project p ON d.project_id = p.id
	INNER JOIN thourus.space s ON p.space_id = s.id
	INNER JOIN thourus.user u ON d.creator_id = u.id
	WHERE p.id = ?;
`
	documents := []entity.Document{}

	rows, err := gw.db.QueryContext(ctx, query, projectId)

	for rows.Next() {
		document := entity.Document{}
		if err = rows.Scan(
			&document.Id,
			&document.Uid,
			&document.Name,
			&document.Path,
			&document.Creator.Uid,
			&document.Project.Id,
			&document.Project.Uid,
			&document.Project.Name,
			&document.Space.Id,
			&document.Space.Uid,
			&document.Space.Name,
		); err != nil {
			return documents, err
		}
		documents = append(documents, document)
	}

	return documents, nil
}

func (gw *UserGateway) GetUsersInProjectByUid(ctx context.Context, projectUid string) ([]entity.Document, error) {

	const query = `
	SELECT d.id, d.uid, d.name, d.path, u.uid,
		   p.id, p.uid, p.name,
		   s.id, s.uid, s.name
	FROM thourus.document d
	INNER JOIN thourus.project p ON d.project_id = p.id
	INNER JOIN thourus.space s ON p.space_id = s.id
	INNER JOIN thourus.user u ON d.creator_id = u.id
	WHERE p.id = ?;
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
			&document.Project.Id,
			&document.Project.Uid,
			&document.Project.Name,
			&document.Space.Id,
			&document.Space.Uid,
			&document.Space.Name,
		); err != nil {
			return documents, err
		}
		documents = append(documents, document)
	}

	return documents, nil
}

func (gw *UserGateway) CreateUser(ctx context.Context, documentName string, documentPath string, creatorId int, projectId int) (int, error) {

	const query = `
	INSERT INTO thourus.document (name, path, creator_id, status_id, project_id)
	VALUES (?, ?, ?, 1, ?);
`
	var createdDocumentId int

	res, err := gw.db.ExecContext(ctx, query, documentName, documentPath, creatorId, projectId)
	if err != nil {
		return createdDocumentId, err
	}

	createdProjectId64, err := res.LastInsertId()
	if err != nil {
		return createdDocumentId, err
	}
	createdDocumentId = int(createdProjectId64)

	return createdDocumentId, nil
}

func (gw *UserGateway) RenameUserById(ctx context.Context, newDocumentName string, documentId int) error {

	const query = `
	UPDATE thourus.document d SET d.name = ? WHERE d.id = ?;
`
	_, err := gw.db.ExecContext(ctx, query, newDocumentName, documentId)
	if err != nil {
		return err
	}

	return nil
}

func (gw *UserGateway) RenameUserByUid(ctx context.Context, newProjectName string, documentUid string) error {

	const query = `
	UPDATE thourus.document d SET d.name = ? WHERE d.uid = ?;
`
	_, err := gw.db.ExecContext(ctx, query, newProjectName, documentUid)
	if err != nil {
		return err
	}

	return nil
}

func (gw *UserGateway) DeleteUserById(ctx context.Context, documentId int) error {

	const query = `
	DELETE FROM thourus.document d WHERE d.id = ?;
`
	_, err := gw.db.ExecContext(ctx, query, documentId)
	if err != nil {
		return err
	}

	return nil
}

func (gw *UserGateway) DeleteUserByUid(ctx context.Context, documentUid string) error {

	const query = `
	DELETE FROM thourus.document d WHERE d.uid = ?;
`
	_, err := gw.db.ExecContext(ctx, query, documentUid)
	if err != nil {
		return err
	}

	return nil
}
