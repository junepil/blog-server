package s3

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Uploader struct {
	Client     *s3.Client
	BucketName string
}

func NewUploader() (*Uploader, error) {
	bucketName := os.Getenv("AWS_S3_BUCKET")
	if bucketName == "" {
		return nil, fmt.Errorf("AWS_S3_BUCKET environment variable not set")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	return &Uploader{
		Client:     client,
		BucketName: bucketName,
	}, nil
}

func (u *Uploader) UploadFile(ctx context.Context, key string, file io.Reader) (string, error) {
	_, err := u.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &u.BucketName,
		Key:    &key,
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	// Construct the public URL
	// Note: For this URL to be publicly accessible, your S3 bucket must have public access enabled
	// and you might need to configure object ACLs or bucket policies.
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.BucketName, os.Getenv("AWS_REGION"), key)
	return url, nil
}
