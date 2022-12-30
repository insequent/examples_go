package main

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func NewS3Client() (s *S3Client) {
	ctx := context.TODO()
	// Use the config below to log all AWS Requests. Useful for debugging
	//cfg, err := config.LoadDefaultConfig(ctx, config.WithClientLogMode(aws.LogRequest|aws.LogRetries))
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}
	client := s3.NewFromConfig(cfg)
	s = &S3Client{
		client: client,
		config: cfg,
		ctx:    ctx,
	}

	return s
}

type S3Client struct {
	client *s3.Client
	config aws.Config
	ctx    context.Context
}

func (s *S3Client) Cleanup(bucket string) {
	objs := s.ListObjects(bucket)

	for _, obj := range objs {
		s.DeleteObject(bucket, obj)
	}
	s.DeleteBucket(bucket)
}

func (s *S3Client) CreateBucket(bucket string) {
	// Bucket creation requires a location constraint to be added
	_, err := s.client.CreateBucket(s.ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(s.config.Region),
		},
	})

	if err != nil {
		log.Fatalf("Failed to create bucket %v: %v", bucket, err)
	}
}

func (s *S3Client) DeleteBucket(bucket string) {
	_, err := s.client.DeleteBucket(s.ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		log.Fatalf("Failed to delete bucket %v: %v", bucket, err)
	}
}

func (s *S3Client) CreateObject(bucket string, obj string, body io.ReadSeeker, sseAlg, sseKey *string) {
	var input *s3.PutObjectInput
	if sseAlg != nil && sseKey != nil {
		sum := md5.Sum([]byte(*sseKey))
		input = &s3.PutObjectInput{
			Bucket:               aws.String(bucket),
			Key:                  aws.String(obj),
			Body:                 body,
			SSECustomerAlgorithm: aws.String(*sseAlg),
			SSECustomerKey:       aws.String(base64.StdEncoding.EncodeToString([]byte(*sseKey))),
			SSECustomerKeyMD5:    aws.String(base64.StdEncoding.EncodeToString(sum[:])),
		}
	} else {
		input = &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(obj),
			Body:   body,
		}
	}

	_, err := s.client.PutObject(s.ctx, input)

	if err != nil {
		log.Fatalf("Failed to create object %v: %v", obj, err)
	}
}

func (s *S3Client) DeleteObject(bucket string, obj string) {
	_, err := s.client.DeleteObject(s.ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})

	if err != nil {
		log.Fatalf("Failed to delete object %v: %v", obj, err)
	}
}

func (s *S3Client) ListObjects(bucket string) (objs []string) {
	l, err := s.client.ListObjects(s.ctx, &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		log.Fatalf("Failed to list objects for bucket %v: %v", bucket, err)
	}

	objs = make([]string, len(l.Contents))
	for num, obj := range l.Contents {
		objs[num] = *obj.Key
	}

	return objs
}

func (s *S3Client) PresignedURL(bucket string, obj string, sseAlg, sseKey *string) (url string) {
	var input *s3.GetObjectInput
	if sseAlg != nil && sseKey != nil {
		input = &s3.GetObjectInput{
			Bucket:               aws.String(bucket),
			Key:                  aws.String(obj),
			SSECustomerAlgorithm: aws.String(*sseAlg),
			SSECustomerKey:       aws.String(base64.StdEncoding.EncodeToString([]byte(*sseKey))),
		}
	} else {
		input = &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(obj),
		}
	}

	psClient := s3.NewPresignClient(s.client)
	req, err := psClient.PresignGetObject(s.ctx, input)
	if err != nil {
		log.Fatalf("Failed to sign request: %v", err)
	}

	return req.URL
}
