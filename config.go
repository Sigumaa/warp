package main

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

var (
	// ErrNoDBURI is returned when the DB_URI environment variable is not set.
	ErrNoDBURI = errors.New("DB_URI not set")

	// ErrLoadEnv is returned when the .env file cannot be loaded.
	ErrLoadEnv = errors.New("error loading .env file")
)

// loadEnv loads the .env file.
func loadURI() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", ErrLoadEnv
	}
	uri := os.Getenv("DB_URI")
	if uri == "" {
		return "", ErrNoDBURI
	}
	return uri, nil
}
