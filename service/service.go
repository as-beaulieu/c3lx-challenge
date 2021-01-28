package service

import (
	"c3lx-challenge/src/dao"
	"go.uber.org/zap"
)

type Service interface {
}

type service struct {
	//package	PackageType
	logger   zap.Logger
	postgres dao.DAO
}

type ServiceBuilder struct {
	service
}

func (sb ServiceBuilder) WithLogger(logger zap.Logger) ServiceBuilder {
	a := sb
	a.logger = logger
	return a
}

func (sb ServiceBuilder) WithPostgres(dao dao.DAO) ServiceBuilder { //Point to Interface of package to be injected
	a := sb
	a.postgres = dao
	return a
}

func (sb ServiceBuilder) Build() *service {
	return &sb.service
}
