package myRedis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func Set(conn redis.Conn, key, value string) {
	_, err := conn.Do("SET", key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Add(conn redis.Conn, key, value string) {
	_, err := conn.Do("LPUSH", key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Get(conn redis.Conn, key string) string {
	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return value
}

func GetList(conn redis.Conn, key string) []string {
	values, err := redis.Strings(conn.Do("LRANGE", key, 0, -1))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return values
}

func Pop(conn redis.Conn, key string) string {
	value, err := redis.String(conn.Do("LPOP", key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return value
}
