package config

import(
	"fmt"
	"os"
	"github.com/joho/godotenv"
) 

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    ServerPort string
}

func LoadConfig() (*Config,error) {
	godotenv.Load()

	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "user-api"),
        ServerPort: getEnv("SERVER_PORT", "3000"),
	}

	return config,nil
}

func (c *Config) GetDBConnectionString() string{
   return fmt.Sprintf(
	 "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
   )
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}