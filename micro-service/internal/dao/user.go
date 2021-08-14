package dao

import (
	"micro-service/internal/model"
	protoUser "micro-service/proto/user"
)

func (d *Dao) Add (username, email, password string) (int64, error) {
	req := &protoUser.User{
		Username: username,
		Email: email,
		Password: password,
	}
	user := new(model.User)
	return user.Add(d.dbEngine, req)
}

func (d *Dao) Get (id int64) (*protoUser.User, error) {
	user := new(model.User)
	return user.Get(d.dbEngine, id)
}
