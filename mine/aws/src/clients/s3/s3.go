package s3

import (
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	s3sdk "github.com/aws/aws-sdk-go/service/s3"
)

func New() (s *S3Client) {
	cfg := aws.Config{
		LogLevel: aws.LogLevel(aws.LogDebugWithSigning),
		Region:   aws.String("us-west-2"),
	}
	opts := session.Options{
		Config:  cfg,
		Profile: "amazon",
	}
	ses := session.Must(session.NewSessionWithOptions(opts))
	svc := s3sdk.New(ses)
	s = &S3Client{
		client:  svc,
		session: ses,
	}

	return s
}

type S3Client struct {
	client  *s3sdk.S3
	session *session.Session
}

func (s *S3Client) Cleanup(bucket string) {
	objs := s.ListObjects(bucket)

	for _, obj := range objs {
		s.DeleteObject(bucket, obj)
	}
	s.DeleteBucket(bucket)
}

func (s *S3Client) CreateBucket(bucket string) {
	_, err := s.client.CreateBucket(&s3sdk.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to create bucket %v: %v", bucket, err))
	}
}

func (s *S3Client) DeleteBucket(bucket string) {
	_, err := s.client.DeleteBucket(&s3sdk.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to delete bucket %v: %v", bucket, err))
	}
}

func (s *S3Client) CreateObject(bucket string, obj string, body io.ReadSeeker, sseAlg interface{}, sseKey interface{}) {
	alg, ok := sseAlg.(string)
	key, ok2 := sseKey.(string)

	var input *s3sdk.PutObjectInput
	if ok && ok2 {
		input = &s3sdk.PutObjectInput{
			Bucket:               aws.String(bucket),
			Key:                  aws.String(obj),
			Body:                 body,
			SSECustomerAlgorithm: aws.String(alg),
			SSECustomerKey:       aws.String(key),
		}
	} else {
		input = &s3sdk.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(obj),
			Body:   body,
		}
	}

	_, err := s.client.PutObject(input)

	if err != nil {
		panic(fmt.Sprintf("Failed to create object %v: %v", obj, err))
	}
}

func (s *S3Client) DeleteObject(bucket string, obj string) {
	_, err := s.client.DeleteObject(&s3sdk.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to delete object %v: %v", obj, err))
	}
}

func (s *S3Client) ListObjects(bucket string) (objs []string) {
	l, err := s.client.ListObjects(&s3sdk.ListObjectsInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to list objects for bucket %v: %v", bucket, err))
	}

	objs = make([]string, len(l.Contents))
	for num, obj := range l.Contents {
		objs[num] = *obj.Key
	}

	return objs
}

func (s *S3Client) PresignedURL(bucket string, obj string, sseAlg interface{}, sseKey interface{}) (url string) {
	alg, ok := sseAlg.(string)
	key, ok2 := sseKey.(string)

	var input *s3sdk.GetObjectInput
	if ok && ok2 {
		input = &s3sdk.GetObjectInput{
			Bucket:               aws.String(bucket),
			Key:                  aws.String(obj),
			SSECustomerAlgorithm: aws.String(alg),
			SSECustomerKey:       aws.String(key),
		}
	} else {
		input = &s3sdk.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(obj),
		}
	}

	req, _ := s.client.GetObjectRequest(input)
	url, err := req.Presign(30 * time.Minute)

	if err != nil {
		panic(fmt.Sprintf("Failed to sign request: %v", err))

	}

	return url
}
