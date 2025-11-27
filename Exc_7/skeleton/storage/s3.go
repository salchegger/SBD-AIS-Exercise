package storage

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"ordersystem/secrets"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	S3EndpointEnvKey        = "S3_ENDPOINT"
	S3AccessKeyEnvKey       = "S3_ACCESS_KEY_ID"
	S3SecretAccessKeyEnvKey = "S3_SECRET_ACCESS_KEY"
	OrdersBucket            = "orders"
)

// CreateS3client creates an S3Client that implements all functions of the interfaces. Storage
// interface. The S3Client requires the following env variables to be set:
// - S3_ENDPOINT
// - S3_ACCESS_KEY_ID or S3_ACCESS_KEY_ID_FILE
// - S3_SECRET_ACCESS_KEY or S3_SECRET_ACCESS_KEY_FILE
func CreateS3client() (*minio.Client, error) {
	slog.Info("Connecting to S3")
	s3Endpoint, exists := os.LookupEnv(S3EndpointEnvKey)
	if !exists {
		return nil, errors.New(fmt.Sprintf("Environment variable %s not set", S3EndpointEnvKey))
	}
	s3AccessKeyId, err := secrets.LoadSecretOrEnv(S3AccessKeyEnvKey)
	if err != nil {
		return nil, err
	}
	secretAccessKey, err := secrets.LoadSecretOrEnv(S3SecretAccessKeyEnvKey)
	if err != nil {
		return nil, err
	}

	client, err := minio.New(s3Endpoint, &minio.Options{
		Secure: false,
		Creds: credentials.NewStaticV4(
			s3AccessKeyId,
			secretAccessKey,
			"",
		),
	})
	if err != nil {
		return nil, err
	}

	// start healthcheck on the endpoint
	cancelFn, err := client.HealthCheck(1 * time.Second)
	if err != nil {
		return nil, err
	}
	defer cancelFn()
	// continually check for 10 seconds if s3 is available
	alive := false
	deadline := time.Now().Add(10 * time.Second)
	for deadline.After(time.Now()) {
		if client.IsOnline() {
			alive = true
			break
		}
		time.Sleep(1 * time.Second)
	}
	if !alive {
		return nil, errors.New("S3 is not reachable, timeout after waiting for 10 seconds")
	}
	// create bucket if not exists
	exists, err = client.BucketExists(context.Background(), OrdersBucket)
	if err != nil {
		return nil, err
	}
	if exists {
		return client, nil
	}
	err = client.MakeBucket(context.Background(), OrdersBucket, minio.MakeBucketOptions{})
	if err != nil {
		return nil, err
	}
	return client, nil
}
