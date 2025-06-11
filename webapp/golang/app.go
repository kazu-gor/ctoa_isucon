package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	host := os.Getenv("ISUCONP_DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("ISUCONP_DB_PORT")
	if port == "" {
		port = "3306"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Failed to read DB port number from an environment variable ISUCONP_DB_PORT.\nError: %s", err.Error())
	}
	user := os.Getenv("ISUCONP_DB_USER")
	if user == "" {
		user = "isuconp"
	}
	password := os.Getenv("ISUCONP_DB_PASSWORD")
	if password == "" {
		password = "isuconp"
	}
	dbname := os.Getenv("ISUCONP_DB_NAME")
	if dbname == "" {
		dbname = "isuconp"
	}
} 