// config/config.go
package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	RedisPassword string `json:"redisPassword"`
	RedisAddr     string `json:"redisAddr"`
	RedisDB       int    `json:"redisDB"`
	RoomName      string `json:"roomName"`
	RoomKey       string `json:"roomKey"`
}

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

var (
	cfg        *Config
	redisCli   *RedisClient
	configPath = "./chatRoom.conf.json"
)

func Init() error {
	// 加载配置文件
	if err := loadConfig(); err != nil {
		return err
	}

	// 初始化Redis客户端
	cli, err := initRedisClient()
	if err != nil {
		return err
	}
	redisCli = cli

	return nil
}

func GetRedis() *RedisClient {
	return redisCli
}

func GetConfig() *Config {
	return cfg
}

func loadConfig() error {
	cfg = &Config{
		RedisAddr: "localhost:6379", // 默认值
	}

	content, err := os.ReadFile(configPath)
	if os.IsNotExist(err) {
		return createDefaultConfig()
	} else if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	if err := json.Unmarshal(content, cfg); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	return validateConfig()
}

func createDefaultConfig() error {
	defaultCfg := Config{
		RedisAddr: "localhost:6379",
		RoomName:  "默认聊天室",
		RoomKey:   "chat_room",
	}

	content, _ := json.MarshalIndent(defaultCfg, "", "  ")
	if err := os.WriteFile(configPath, content, 0644); err != nil {
		return fmt.Errorf("创建配置文件失败: %w", err)
	}

	cfg = &defaultCfg
	return nil
}

func validateConfig() error {
	if cfg.RedisAddr == "" {
		return fmt.Errorf("redis地址不能为空")
	}
	return nil
}

func initRedisClient() (*RedisClient, error) {
	options := &redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	}

	client := redis.NewClient(options)
	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("Redis连接测试失败: %w", err)
	}

	return &RedisClient{
		client: client,
		ctx:    ctx,
	}, nil
}

// 以下是Redis客户端方法
func (rc *RedisClient) PublishMessage(channel string, message []byte) error {
	return rc.client.Publish(rc.ctx, channel, message).Err()
}

func (rc *RedisClient) SaveMessage(message []byte) error {
	return rc.client.RPush(rc.ctx, cfg.RoomKey, message).Err()
}

func (rc *RedisClient) GetHistory() ([]string, error) {
	return rc.client.LRange(rc.ctx, cfg.RoomKey, 0, -1).Result()
}

func (rc *RedisClient) Subscribe() *redis.PubSub {
	return rc.client.Subscribe(rc.ctx, cfg.RoomKey)
}

func (rc *RedisClient) Close() error {
	return rc.client.Close()
}
