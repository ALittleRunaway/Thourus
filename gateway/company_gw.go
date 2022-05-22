package gateway

import (
	"context"
	"database/sql"
	"thourus-api/domain/entity"
)

type CompanyGw interface {
	GetCompany(ctx context.Context, uid string) (entity.Company, error)
}

type CompanyGateway struct {
	db *sql.DB
}

func NewCompanyGateway(db *sql.DB) *CompanyGateway {
	return &CompanyGateway{
		db: db,
	}
}

func (gw *CompanyGateway) GetCompany(ctx context.Context, uid string) (entity.Company, error) {
	return entity.Company{}, nil
}
