package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitConfig() *ProgramConfig {
	godotenv.Load()

	var res = new(ProgramConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

type ProgramConfig struct {
	SECRET         			string
	REFRESH_SECRET 			string
	SERVER_PORT    			string
	CDN_CLOUD_NAME  		string
	CDN_API_KEY     		string
	CDN_API_SECRET  		string
	CDN_FOLDER_ARTICLES		string
	CDN_FOLDER_COURSE		string
	OPENAI_KEY				string
}

type DatabaseConfig struct {
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
}

type MongoConfig struct {
	MONGO_URI		string
	MONGO_DB_NAME	string
}
func LoadDBConfig() *DatabaseConfig {
	var res = new(DatabaseConfig)

	if val, found := os.LookupEnv("DB_USER"); found {
		res.DB_USER = val
	}

	if val, found := os.LookupEnv("DB_PASS"); found {
		res.DB_PASS = val
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		res.DB_HOST = val
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		res.DB_PORT = val
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		res.DB_NAME = val
	}

	return res
}

func loadConfig() *ProgramConfig {
	var res = new(ProgramConfig)

	if val, found := os.LookupEnv("SECRET"); found {
		res.SECRET = val
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.REFRESH_SECRET = val
	}

	if val, found := os.LookupEnv("SERVER_PORT"); found {
		res.SERVER_PORT = val
	}

	if val, found := os.LookupEnv("CDN_CLOUD_NAME"); found {
		res.CDN_CLOUD_NAME = val
	}

	if val, found := os.LookupEnv("CDN_API_KEY"); found {
		res.CDN_API_KEY = val
	}

	if val, found := os.LookupEnv("CDN_API_SECRET"); found {
		res.CDN_API_SECRET = val
	}

	if val, found := os.LookupEnv("CDN_FOLDER_ARTICLES"); found {
		res.CDN_FOLDER_ARTICLES = val
	}

	if val, found := os.LookupEnv("CDN_FOLDER_COURSE"); found {
		res.CDN_FOLDER_COURSE = val
	}

	if val, found := os.LookupEnv("OPENAI_KEY"); found {
		res.OPENAI_KEY = val
	}

	return res
}

func LoadMongoConfig() *MongoConfig {
	var res = new(MongoConfig)

	if val, found := os.LookupEnv("MONGO_URI"); found {
		res.MONGO_URI = val
	}

	if val, found := os.LookupEnv("MONGO_DB_NAME"); found {
		res.MONGO_DB_NAME = val
	}

	return res
}
