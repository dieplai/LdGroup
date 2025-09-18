package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "perfume_quiz")

	log.Printf("Connecting to database: %s@%s:%s/%s", user, host, port, dbname)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		log.Printf("DSN (without password): %s:****@tcp(%s:%s)/%s", user, host, port, dbname)
		log.Fatal("Database connection failed")
	}

	log.Println("Database connected successfully!")
	return db
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}