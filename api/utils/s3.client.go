package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var bucketRegion string
var bucketName string

func init() {
	godotenv.Load()
	bucketRegion = os.Getenv("BUCKET_REGION")
	bucketName = os.Getenv("BUCKET_NAME")
}

func UploadFile(file *multipart.FileHeader, fileName string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	fileData, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileData.Close()

	key := "documents/" + uuid.New().String() + "/" + fileName

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:   aws.String(key),
		Body:   fileData,
	})
	if err != nil {
		return "", err
	}

    fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
	
	return fileURL, nil
}
