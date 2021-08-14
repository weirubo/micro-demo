package service

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"micro-service/pkg"
	protoUser "micro-service/proto/user"
)

func (s service) Add (param *protoUser.User) (int64, error) {
	return s.dao.Add(param.Username, param.Email, param.Password)
}

func (s service) Get (param *protoUser.User) (*protoUser.User, error) {
	// redis
	key := fmt.Sprintf("user_%v", param.Id)
	c := pkg.NewRedisClient()
	defer c.Close()

	// 查 Redis
	v, err := redis.Values(c.Do("HGETALL", key))
	if err != nil {
		fmt.Println("查", err)
		return nil, err
	}

	u := new(protoUser.User)
	if err := redis.ScanStruct(v, u); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if u != nil {
		return u, nil
	}

	// 查 MySQL
	data, err := s.dao.Get(param.Id)
	if err != nil {
		fmt.Println(err)
	}

	// 写 Redis
	fmt.Printf("data=%+v", data)
	if _, err := c.Do("HMSET", redis.Args{}.Add(key).AddFlat(data)...); err != nil {
		fmt.Println("写", err)
		return nil, err
	}
	return data, nil
}