package myGin

import (
	"chatroom/config"
	"chatroom/myRedis"
	"chatroom/postType"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gomodule/redigo/redis"
)

func listener(conn redis.Conn, key string, channel chan postType.PostRequest, closeServer_start chan bool) {
	// 上次检查时列表的长度
	var lastLength int64 = config.ListenerLastLen()

	for {
		currentLength, err := redis.Int64(conn.Do("LLEN", key))
		if err != nil {
			fmt.Println("获取列表长度失败，redis连接失效：", err)
			// channel <- postType.ParseError(postType.ErrorMsg("902", "Redis连接失效"))
			conn = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0, 3)
			if conn == nil {
				closeServer_start <- true
				Page <- RoutePage.ADDRESS_PAGE
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
				// channel <- postType.ParseError(postType.ErrorMsg("902", "Redis连接失效"))
				conn = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0, 3)
				if conn == nil {
					closeServer_start <- true
					Page <- RoutePage.ADDRESS_PAGE
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

func StartListen(conn redis.Conn, key string, channel chan postType.PostRequest, closeServer_start chan bool) {
	fmt.Println("开始监听[", key, "]...")
	go listener(conn, key, channel, closeServer_start)
}
