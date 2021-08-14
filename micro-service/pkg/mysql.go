package pkg

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func NewDBEngine () (*xorm.Engine, error) {
	var err error
	engine, err := xorm.NewEngine("mysql", "root:@/micro?charset=utf8mb4")
	if err != nil {
		return nil, err
	}
	if err = engine.Ping(); err != nil {
		return nil, err
	}
	// engine.ShowSQL(true)
	return engine, nil
}
