package lib

import (
	"fmt"
	"os"
)

func GeneratePublicURL(key string) string {
	bucket := os.Getenv("BUCKET_NAME")
	endpoint := os.Getenv("AWS_ENDPOINT_URL_S3")

	url := fmt.Sprintf("%s/%s/%s", endpoint, bucket, key)
	return url
}
