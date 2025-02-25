package myRedis

import (
	"fmt"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func Connect(addr string, password string, db string, dbIndex int) redis.Conn {
	database, _ := strconv.Atoi(db)
	conn, err := redis.Dial("tcp", addr, redis.DialPassword(password), redis.DialDatabase(database+dbIndex))
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
