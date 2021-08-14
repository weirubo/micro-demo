package service

import (
	"micro-service/internal/dao"
	"micro-service/pkg"
)

type service struct {
	dao *dao.Dao
}

func New () service {
	svc := service{
		dao: dao.New(pkg.DBEngine),
	}
	return svc
}