package config

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBSSLMode     string
	AppPort       string
	AppHostServer string
	AppPrefix     string
	AppEnv        string
	MongoDatabase string
	MongoUser     string
	MongoPassword string
}

var EnvConfig *Config

func Load() {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "staging"
	}

	var BaseDir = FindDir("go.mod")
	if BaseDir == "" {
		log.Fatalln("Please use `go mod init` to create project.")
		return
	}

	envFile := path.Join(BaseDir, ".env."+env)
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Failed to load %s: %+v\n", envFile, err)
		return
	}

	EnvConfig = &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_DATABASE"),
		DBSSLMode:     os.Getenv("DB_SSLMODE"),
		AppPort:       os.Getenv("APP_PORT"),
		AppPrefix:     os.Getenv("APP_PREFIX"),
		AppHostServer: os.Getenv("APP_HOST_SERVER"),
		MongoDatabase: os.Getenv("MONGO_DB_DATABASE"),
		MongoUser:     os.Getenv("MONGO_DB_USERNAME"),
		MongoPassword: os.Getenv("MONGO_DB_PASSWORD"),
		AppEnv:        env,
	}
}

// GetDSN PostgreSQL
func GetDSN() string {
	return "host=" + EnvConfig.DBHost +
		" port=" + EnvConfig.DBPort +
		" user=" + EnvConfig.DBUser +
		" password=" + EnvConfig.DBPassword +
		" dbname=" + EnvConfig.DBName +
		" sslmode=" + EnvConfig.DBSSLMode
}

func GetMongoURI() string {
	return "mongodb+srv://" + EnvConfig.MongoUser + ":" +
		EnvConfig.MongoPassword +
		"@" + EnvConfig.MongoDatabase +
		".cobqrxh.mongodb.net/?retryWrites=true&w=majority&appName=" + EnvConfig.MongoDatabase
}

func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func FindDir(filename string) string {
	dir, _ := os.Getwd()
	for !IsFile(path.Join(dir, filename)) {
		pdir := path.Dir(dir)
		if pdir == dir {
			return ""
		}
		dir = pdir
	}
	return dir
}
