package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type APIConfig struct {
	S3Client *s3.Client
	S3Bucket string
	S3Region string
}

// LoadConfig initializes APIConfig using .env and AWS SDK
func LoadConfig() *APIConfig {
	// Load .env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	s3Region := os.Getenv("S3_REGION")

	// Load AWS SDK config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(s3Region))
	if err != nil {
		log.Fatalf("Unable to load AWS config: %v", err)
	}

	// Create S3 client
	s3Client := s3.NewFromConfig(cfg)

	return &APIConfig{
		S3Client: s3Client,
		S3Bucket: s3Bucket,
		S3Region: s3Region,
	}
}
