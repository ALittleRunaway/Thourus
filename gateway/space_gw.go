package gateway

import (
	"context"
	"database/sql"
	"thourus-api/domain/entity"
)

type SpaceGw interface {
	GetSpaceById(ctx context.Context, spaceId int) (entity.Space, error)
	GetSpaceByUid(ctx context.Context, spaceUid string) (entity.Space, error)
	GetSpacesInCompany(ctx context.Context, companyId int) ([]entity.Space, error)
	CreateSpace(ctx context.Context, spaceName string, companyId int) (int, error)
	RenameSpaceById(ctx context.Context, newSpaceName string, spaceId int) error
	RenameSpaceByUid(ctx context.Context, newSpaceName string, spaceUid string) error
	DeleteSpaceById(ctx context.Context, id int) error
	DeleteSpaceByUid(ctx context.Context, uid string) error
	AddUserToSpace(ctx context.Context, spaceUid string, userUid string) error
}

type SpaceGateway struct {
	db *sql.DB
}

func NewSpaceGateway(db *sql.DB) *SpaceGateway {
	return &SpaceGateway{
		db: db,
	}
}

func (gw *SpaceGateway) GetSpaceById(ctx context.Context, spaceId int) (entity.Space, error) {

	const query = `
	SELECT s.id, s.uid, s.name, c.id, c.uid, c.name FROM thourus.space s
	INNER JOIN thourus.company c ON s.company_id = c.id
	WHERE s.id = ?;
`
	space := entity.Space{}

	rows, err := gw.db.QueryContext(ctx, query, spaceId)

	for rows.Next() {
		if err = rows.Scan(
			&space.Id,
			&space.Uid,
			&space.Name,
			&space.Company.Id,
			&space.Company.Uid,
			&space.Company.Name,
		); err != nil {
			return space, err
		}
	}

	return space, nil
}

func (gw *SpaceGateway) GetSpaceByUid(ctx context.Context, spaceUid string) (entity.Space, error) {

	const query = `
	SELECT s.id, s.uid, s.name, c.id, c.uid, c.name FROM thourus.space s
	INNER JOIN thourus.company c ON s.company_id = c.id
	WHERE s.uid = ?;
`
	space := entity.Space{}

	rows, err := gw.db.QueryContext(ctx, query, spaceUid)

	for rows.Next() {
		if err = rows.Scan(
			&space.Id,
			&space.Uid,
			&space.Name,
			&space.Company.Id,
			&space.Company.Uid,
			&space.Company.Name,
		); err != nil {
			return space, err
		}
	}

	return space, nil
}

func (gw *SpaceGateway) GetSpacesInCompany(ctx context.Context, companyId int) ([]entity.Space, error) {

	const query = `
	SELECT s.id, s.uid, s.name FROM thourus.space s
	INNER JOIN thourus.company c ON s.company_id = c.id
	WHERE c.id = ?;
`
	spaces := []entity.Space{}

	rows, err := gw.db.QueryContext(ctx, query, companyId)

	for rows.Next() {
		space := entity.Space{}
		if err = rows.Scan(
			&space.Id,
			&space.Uid,
			&space.Name,
			&space.Company.Id,
			&space.Company.Uid,
			&space.Company.Name,
		); err != nil {
			return spaces, err
		}
		spaces = append(spaces, space)
	}

	return spaces, nil
}

func (gw *SpaceGateway) CreateSpace(ctx context.Context, spaceName string, companyId int) (int, error) {

	const query = `
	INSERT INTO thourus.space (name, company_id) VALUES (?, ?);
`
	var createdSpaceId int

	res, err := gw.db.ExecContext(ctx, query, spaceName, companyId)
	if err != nil {
		return createdSpaceId, err
	}

	createdProjectId64, err := res.LastInsertId()
	if err != nil {
		return createdSpaceId, err
	}
	createdSpaceId = int(createdProjectId64)

	return createdSpaceId, nil
}

func (gw *SpaceGateway) RenameSpaceById(ctx context.Context, newSpaceName string, spaceId int) error {

	const query = `
	UPDATE thourus.space s SET s.name = ? WHERE s.id = 1;
`
	_, err := gw.db.ExecContext(ctx, query, newSpaceName, spaceId)
	if err != nil {
		return err
	}

	return nil
}

func (gw *SpaceGateway) RenameSpaceByUid(ctx context.Context, newSpaceName string, spaceUid string) error {

	const query = `
	UPDATE thourus.space s SET s.name = ? WHERE s.uid = 1;
`
	_, err := gw.db.ExecContext(ctx, query, newSpaceName, spaceUid)
	if err != nil {
		return err
	}

	return nil
}

func (gw *SpaceGateway) DeleteSpaceById(ctx context.Context, id int) error {

	const query = `
	DELETE FROM thourus.company c WHERE c.id = ?;
`
	_, err := gw.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (gw *SpaceGateway) DeleteSpaceByUid(ctx context.Context, uid string) error {

	const query = `
	DELETE FROM thourus.space c WHERE c.uid = ?;
`
	_, err := gw.db.ExecContext(ctx, query, uid)
	if err != nil {
		return err
	}

	return nil
}

func (gw *SpaceGateway) AddUserToSpace(ctx context.Context, spaceUid string, userUid string) error {
	const query = `
	INSERT INTO thourus.space_user (space_id, user_id) VALUES (?, ?);
`
	_, err := gw.db.ExecContext(ctx, query, spaceUid, userUid)
	if err != nil {
		return err
	}

	return nil
}
