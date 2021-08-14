package model

import (
	protoUser "micro-service/proto/user"
	"xorm.io/xorm"
)

type User struct {}

func (u User) Add (engine *xorm.Engine, req *protoUser.User) (int64, error){
	if err := engine.Sync2(new(protoUser.User)); err != nil {
		return 0, err
	}
	return engine.Insert(req)
}

func (u User) Get (engine *xorm.Engine, id int64) (*protoUser.User, error){
	user := &protoUser.User{}
	_, err := engine.Where("id = ?", id).Get(user)
	if err != nil {
		return user, err
	}
	return user, nil
}
