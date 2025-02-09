package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Region    string
	Bucket    string
	Endpoint  string
	AccessKey string
	SecretKey string
}

type S3Repository struct {
	client *s3.Client
	config S3Config
}

func NewS3Repository(cfg S3Config) (*S3Repository, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     cfg.AccessKey,
				SecretAccessKey: cfg.SecretKey,
			}, nil
		})),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS config: %w", err)
	}

	return &S3Repository{
		client: s3.NewFromConfig(awsCfg),
		config: cfg,
	}, nil
}

func (r *S3Repository) Store(ctx context.Context, data []byte) error {
	key := fmt.Sprintf("logs/%s/%d", time.Now().Format("2006-01-02"), time.Now().UnixNano())

	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &r.config.Bucket,
		Key:    &key,
		Body:   bytes.NewReader(data),
	})

	return err
}