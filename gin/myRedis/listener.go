package myRedis

import (
	"chatroom/config"
	"chatroom/postType"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/gomodule/redigo/redis"
)

func listener(conn redis.Conn, key string, channel chan postType.PostRequest) {
	// 上次检查时列表的长度
	var lastLength int64 = config.ListenerLastLen()

	for {
		currentLength, err := redis.Int64(conn.Do("LLEN", key))
		if err != nil {
			fmt.Println("获取列表长度失败，redis连接失效：", err)
			channel <- postType.ParseError(postType.ErrorMsg("500", "Redis连接失效"))
			for {
				conn = Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB())
				if conn != nil {
					break
				}
				time.Sleep(5 * time.Second)
			}
			continue
		}

		if currentLength == lastLength {
			continue
		}

		if currentLength > lastLength {
			newElements, err := redis.Values(conn.Do("LRANGE", key, lastLength, currentLength-1))
			if err != nil {
				fmt.Println("获取新增元素失败，redis连接失效：", err)
				channel <- postType.ParseError(postType.ErrorMsg("500", "Redis连接失效"))
				for {
					conn = Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB())
					if conn != nil {
						break
					}
					time.Sleep(5 * time.Second)
				}
				continue
			}
			var elements []json.RawMessage
			for _, element := range newElements {
				if raw, ok := element.([]byte); ok {
					elements = append(elements, json.RawMessage(raw))
				} else {
					fmt.Println("不支持的消息类型:", reflect.TypeOf(element))
					continue
				}
			}
			for _, element := range elements {
				var postReq postType.PostRequest
				if err := json.Unmarshal(element, &postReq); err != nil {
					fmt.Println("消息转换时出错:", err)
					continue
				}
				channel <- postReq
			}
		}

		lastLength = currentLength
		config.UpdateListenerLastLen(lastLength)
	}
}

func StartListen(conn redis.Conn, key string, channel chan postType.PostRequest) {
	fmt.Println("开始监听[", key, "]...")
	go listener(conn, key, channel)
}
