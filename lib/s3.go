package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GeneratePublicURL(key string) string {
	bucket := os.Getenv("BUCKET_NAME")
	if bucket == "" {
		fmt.Println("No S3 bucket specified, skipping upload.")
		return ""
	}

	endpoint := os.Getenv("AWS_ENDPOINT_URL_S3")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{
		Region: &region,
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			"",
		),
		Endpoint: aws.String(endpoint),
	})
	if err != nil {
		return ""
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		return ""
	}

	return urlStr
}
