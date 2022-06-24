package usecase

import (
	"context"
	"go.uber.org/zap"
	"thourus-api/domain/entity"
	"thourus-api/gateway"
)

type CompanyUseCase struct {
	companyGw gateway.CompanyGw
	logger    *zap.SugaredLogger
}

func NewCompanyUseCase(companyGw gateway.CompanyGw, logger *zap.SugaredLogger) *CompanyUseCase {
	return &CompanyUseCase{
		companyGw: companyGw,
		logger:    logger,
	}
}

func (uc *CompanyUseCase) GetSpacesInCompany(companyUid string) ([]entity.Space, error) {
	spaces, err := uc.companyGw.GetSpacesInCompany(context.Background(), companyUid)
	if err != nil {
		return []entity.Space{}, err
	}
	return spaces, nil
}

func (uc *CompanyUseCase) GetCompanyInfo(companyUid string) (entity.Company, error) {
	company, err := uc.companyGw.GetCompanyByUid(context.Background(), companyUid)
	if err != nil {
		return entity.Company{}, err
	}
	return company, nil
}
