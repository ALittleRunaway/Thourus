package gateway

import (
	"context"
	"database/sql"
	"thourus-api/domain/entity"
)

type UserGw interface {
	GetUserByPasswordAndLogin(ctx context.Context, UserPassword string, userLogin string) (entity.User, error)
	GetUserByUid(ctx context.Context, userUid string) (entity.User, error)
}

type UserGateway struct {
	db *sql.DB
}

func NewUserGateway(db *sql.DB) *UserGateway {
	return &UserGateway{
		db: db,
	}
}

func (gw *UserGateway) GetUserByPasswordAndLogin(ctx context.Context, UserPassword string, userLogin string) (entity.User, error) {

	const query = `
	SELECT u.id, u.uid, u.name, surname, email, r.id, r.uid, r.name, c.id, c.uid, c.name
	FROM thourus.user u
	INNER JOIN thourus.role r ON u.role_id = r.id
	INNER JOIN thourus.space_user su ON u.id = su.user_id
	INNER JOIN thourus.space s ON su.space_id = s.id
	INNER JOIN thourus.company c ON s.company_id = c.id
	WHERE u.login = ? AND u.password = ?;
`
	user := entity.User{}

	rows, err := gw.db.QueryContext(ctx, query, userLogin, UserPassword)

	for rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Uid,
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.Role.Id,
			&user.Role.Uid,
			&user.Role.Name,
			&user.Company.Id,
			&user.Company.Uid,
			&user.Company.Name,
		); err != nil {
			return user, err
		}
	}

	return user, nil
}

func (gw *UserGateway) GetUserByUid(ctx context.Context, userUid string) (entity.User, error) {

	const query = `
	SELECT u.id, u.uid, u.name, surname, email, r.id, r.uid, r.name, c.id, c.uid, c.name
	FROM thourus.user u
	INNER JOIN thourus.role r ON u.role_id = r.id
	INNER JOIN thourus.space_user su ON u.id = su.user_id
	INNER JOIN thourus.space s ON su.space_id = s.id
	INNER JOIN thourus.company c ON s.company_id = c.id
	WHERE u.uid = ?;
`
	user := entity.User{}

	rows, err := gw.db.QueryContext(ctx, query, userUid)

	for rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Uid,
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.Role.Id,
			&user.Role.Uid,
			&user.Role.Name,
			&user.Company.Id,
			&user.Company.Uid,
			&user.Company.Name,
		); err != nil {
			return user, err
		}
	}

	return user, nil
}

//func (gw *UserGateway) GetUsersInProjectById(ctx context.Context, projectId int) ([]entity.Document, error) {
//
//	const query = `
//	SELECT d.id, d.uid, d.name, d.path, u.uid,
//		   p.id, p.uid, p.name,
//		   s.id, s.uid, s.name
//	FROM thourus.document d
//	INNER JOIN thourus.project p ON d.project_id = p.id
//	INNER JOIN thourus.space s ON p.space_id = s.id
//	INNER JOIN thourus.user u ON d.creator_id = u.id
//	WHERE p.id = ?;
//`
//	documents := []entity.Document{}
//
//	rows, err := gw.db.QueryContext(ctx, query, projectId)
//
//	for rows.Next() {
//		document := entity.Document{}
//		if err = rows.Scan(
//			&document.Id,
//			&document.Uid,
//			&document.Name,
//			&document.Path,
//			&document.Creator.Uid,
//			&document.Project.Id,
//			&document.Project.Uid,
//			&document.Project.Name,
//			&document.Space.Id,
//			&document.Space.Uid,
//			&document.Space.Name,
//		); err != nil {
//			return documents, err
//		}
//		documents = append(documents, document)
//	}
//
//	return documents, nil
//}
//
//func (gw *UserGateway) GetUsersInProjectByUid(ctx context.Context, projectUid string) ([]entity.Document, error) {
//
//	const query = `
//	SELECT d.id, d.uid, d.name, d.path, u.uid,
//		   p.id, p.uid, p.name,
//		   s.id, s.uid, s.name
//	FROM thourus.document d
//	INNER JOIN thourus.project p ON d.project_id = p.id
//	INNER JOIN thourus.space s ON p.space_id = s.id
//	INNER JOIN thourus.user u ON d.creator_id = u.id
//	WHERE p.id = ?;
//`
//	documents := []entity.Document{}
//
//	rows, err := gw.db.QueryContext(ctx, query, projectUid)
//
//	for rows.Next() {
//		document := entity.Document{}
//		if err = rows.Scan(
//			&document.Id,
//			&document.Uid,
//			&document.Name,
//			&document.Path,
//			&document.Creator.Uid,
//			&document.Project.Id,
//			&document.Project.Uid,
//			&document.Project.Name,
//			&document.Space.Id,
//			&document.Space.Uid,
//			&document.Space.Name,
//		); err != nil {
//			return documents, err
//		}
//		documents = append(documents, document)
//	}
//
//	return documents, nil
//}
//
//func (gw *UserGateway) CreateUser(ctx context.Context, documentName string, documentPath string, creatorId int, projectId int) (int, error) {
//
//	const query = `
//	INSERT INTO thourus.document (name, path, creator_id, status_id, project_id)
//	VALUES (?, ?, ?, 1, ?);
//`
//	var createdDocumentId int
//
//	res, err := gw.db.ExecContext(ctx, query, documentName, documentPath, creatorId, projectId)
//	if err != nil {
//		return createdDocumentId, err
//	}
//
//	createdProjectId64, err := res.LastInsertId()
//	if err != nil {
//		return createdDocumentId, err
//	}
//	createdDocumentId = int(createdProjectId64)
//
//	return createdDocumentId, nil
//}
//
//func (gw *UserGateway) RenameUserById(ctx context.Context, newDocumentName string, documentId int) error {
//
//	const query = `
//	UPDATE thourus.document d SET d.name = ? WHERE d.id = ?;
//`
//	_, err := gw.db.ExecContext(ctx, query, newDocumentName, documentId)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (gw *UserGateway) RenameUserByUid(ctx context.Context, newProjectName string, documentUid string) error {
//
//	const query = `
//	UPDATE thourus.document d SET d.name = ? WHERE d.uid = ?;
//`
//	_, err := gw.db.ExecContext(ctx, query, newProjectName, documentUid)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (gw *UserGateway) DeleteUserById(ctx context.Context, documentId int) error {
//
//	const query = `
//	DELETE FROM thourus.document d WHERE d.id = ?;
//`
//	_, err := gw.db.ExecContext(ctx, query, documentId)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (gw *UserGateway) DeleteUserByUid(ctx context.Context, documentUid string) error {
//
//	const query = `
//	DELETE FROM thourus.document d WHERE d.uid = ?;
//`
//	_, err := gw.db.ExecContext(ctx, query, documentUid)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
