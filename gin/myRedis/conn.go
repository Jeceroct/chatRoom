package myRedis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Connect(addr string, password string, db string, dbIndex int, reConnectTimes int) redis.Conn {
	database, _ := strconv.Atoi(db)
	for range reConnectTimes {
		conn, err := redis.Dial("tcp", addr, redis.DialPassword(password), redis.DialDatabase(database+dbIndex))
		if err != nil {
			fmt.Println("连接失败，正在重试...", err)
			time.Sleep(3 * time.Second)
			continue
		}
		fmt.Println("Redis连接成功")
		return conn
	}
	return nil
}

func Close(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		return
	}
}
