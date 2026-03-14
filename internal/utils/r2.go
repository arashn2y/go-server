package utils

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	projectConfig "github.com/arashn0uri/go-server/internal/config"
)

func NewClient(ctx context.Context) (*s3.Client, error) {
	accountID := projectConfig.GetEnv(projectConfig.R2AccountID)
	accessKey := projectConfig.GetEnv(projectConfig.R2AccessKeyID)
	secretKey := projectConfig.GetEnv(projectConfig.R2SecretAccessKey)

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, fmt.Errorf("r2: load config: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(
			fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID),
		)
	})

	return client, nil
}

// UploadFile streams a file directly to R2.
// key         — object path, e.g. "products/123/cover.jpg"
// body        — the file reader (from multipart.File)
// contentType — e.g. "image/jpeg"
func UploadFile(ctx context.Context, key string, body io.Reader, contentType string) error {
	client, err := NewClient(ctx)
	if err != nil {
		return err
	}

	bucket := projectConfig.GetEnv(projectConfig.R2Bucket)

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return fmt.Errorf("r2: upload file: %w", err)
	}

	return nil
}

// PresignDownload returns a presigned GET URL the client uses to fetch an object from R2.
// key     — object path in bucket, e.g. "images/user-123/avatar.jpg"
// expires — how long the URL is valid
func PresignDownload(ctx context.Context, key string, expires time.Duration) (string, error) {
	client, err := NewClient(ctx)
	if err != nil {
		return "", err
	}

	bucket := projectConfig.GetEnv(projectConfig.R2Bucket)
	presigner := s3.NewPresignClient(client)

	result, err := presigner.PresignGetObject(ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(expires),
	)
	if err != nil {
		return "", fmt.Errorf("r2: presign download: %w", err)
	}

	return result.URL, nil
}
