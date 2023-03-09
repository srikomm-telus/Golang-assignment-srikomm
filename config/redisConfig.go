package config

import (
	"Golang-assignment-srikomm/constants"
	"encoding/json"
	"log"
	"os"
)

type RedisConfig struct {
	Production  RedisEnvironmentConfig `json:"PRODUCTION"`
	Development RedisEnvironmentConfig `json:"DEVELOPMENT"`
}

type RedisEnvironmentConfig struct {
	ClientAddress string `json:"RedisClientAddress"`
	Password      string `json:"RedisClientPassword"`
	DB            int    `json:"RedisDB"`
}

func GetRedisConfig(environment string) RedisEnvironmentConfig {
	redisConfig := readRedisConfigFromFile()
	switch environment {
	case constants.PRODUCTION:
		return redisConfig.Production
	case constants.DEVELOPMENT:
		return redisConfig.Development
	default:
		return redisConfig.Development
	}
}

func readRedisConfigFromFile() RedisConfig {
	configJson, err := os.ReadFile("./config/redisConfig.json")
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
