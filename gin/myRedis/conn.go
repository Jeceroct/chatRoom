package myRedis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func Connect(addr string, password string, db int) redis.Conn {
	conn, err := redis.Dial("tcp", addr, redis.DialPassword(password), redis.DialDatabase(db))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Redis连接成功")
	return conn
}

func Close(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		return
	}
}
