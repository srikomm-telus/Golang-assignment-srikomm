package config

import (
	"Golang-assignment-srikomm/constants"
	"encoding/json"
	"log"
	"os"
)

type RedisConfig struct {
	Config RedisEnvironmentConfig `json:"CONFIG"`
}

type RedisEnvironmentConfig struct {
	Environment   string `json:"Environment"`
	ClientAddress string `json:"RedisClientAddress"`
	Password      string `json:"RedisClientPassword"`
	DB            int    `json:"RedisDB"`
}

func GetRedisConfig(environment string) RedisEnvironmentConfig {
	var fileName string
	switch environment {
	case constants.PRODUCTION:
		fileName = "./config/redisProductionConfig.json"
	case constants.DEVELOPMENT:
		fileName = "./config/redisDevelopmentConfig.json"
	default:
		fileName = "./config/redisDevelopmentConfig.json"
	}
	redisConfig := readRedisConfigFromFile(fileName)
	return redisConfig.Config
}

func readRedisConfigFromFile(fileName string) RedisConfig {
	configJson, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var redisConfig RedisConfig
	err = json.Unmarshal(configJson, &redisConfig)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return redisConfig
}
