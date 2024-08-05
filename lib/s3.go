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

// Generates a presigned URL for the given key. Returns the pre-signed URL.
func GeneratePublicURL(key string) string {
	// Get the S3 bucket name
	bucket := os.Getenv("BUCKET_NAME")
	if bucket == "" {
		fmt.Println("No S3 bucket specified.")
		return ""
	}

	// Create the S3 session
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

	// Generate the presigned URL
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

// Generates a list of presigned URLs for all items in the given directory. Returns an array of pre-signed URLs.
func GeneratePublicURLsFromDirectory(directory string) []string {
	// Get the S3 bucket name
	bucket := os.Getenv("BUCKET_NAME")
	if bucket == "" {
		fmt.Println("No S3 bucket specified.")
		return []string{}
	}

	// Create the S3 session
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
		return []string{}
	}

	svc := s3.New(sess)

	// Create the input for the ListObjectsV2 call
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(directory),
	}

	// Get the list of items in the directory
	result, err := svc.ListObjectsV2(input)
	if err != nil {
		return []string{}
	}

	//Remove the items in results that are directories
	for i := 0; i < len(result.Contents); i++ {
		if *result.Contents[i].Key == directory {
			result.Contents = append(result.Contents[:i], result.Contents[i+1:]...)
			i--
		}
	}

	// Generate the presigned URLs
	urls := []string{}
	for _, item := range result.Contents {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    item.Key,
		})
		urlStr, err := req.Presign(15 * time.Minute)

		if err != nil {
			continue
		}

		urls = append(urls, urlStr)
	}

	return urls
}
