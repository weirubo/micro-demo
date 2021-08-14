package dao

import "xorm.io/xorm"

type Dao struct {
	dbEngine *xorm.Engine
}

func New(engine *xorm.Engine) *Dao {
	return &Dao{dbEngine: engine}
}
