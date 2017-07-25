package s3client

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3Client() (s3c *S3Client) {
	opts := session.Options{Profile: "amazon"}
	ses := session.Must(session.NewSessionWithOptions(opts))
	s3 := s3.New(ses)
	s3c := &S3Client{
		client:  s3,
		session: ses,
	}

	return s3c
}

type S3Client struct {
	client  *s3.S3
	session *aws.Session
}

func (s *S3Client) Cleanup(bucket) {
	objs := s.ListObjects(bucket)

	for obj := range objs {
		s.DeleteObject(bucket, obj)
	}
	s.DeleteBucket(bucket)
}

func (s *S3Client) CreateBucket(bucket string) {
	_, err := s.s3c.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to create bucket %v: %v", bucket, err))
	}
}

func (s *S3Client) DeleteBucket(bucket string) {
	_, err := s.s3c.DeleteBucket(&s3.DeletBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to delete bucket %v: %v", bucket, err))
	}
}

func (s *S3Client) CreateObject(bucket string, obj string, body io.ReadSeeker) {
	_, err := s.s3c.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(obj),
		Body:                 body,
		SSECustomerAlgorithm: aws.String(sseAlg),
		SSECustomerKey:       aws.String(sslKey),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to create object %v: %v", obj, err))
	}
}

func (s *S3Client) DeleteObject(bucket string, obj string, sseAlg string, sseKey string) {
	_, err := s.s3c.DeleteObject(&s3.DeletObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to delete object %v: %v", obj, err))
	}
}

func (s *S3Client) ListObjects(b string) (objs []string) {
	l, err := s.s3c.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to list objects for bucket %v: %v", bucket, err))
	}

	objs = make([]string, len(l.Contents))
	for obj := range l.Contents {
		objs.append(obj.Key)
	}
}

func (s *S3Client) PresignedURL(bucket string, obj string, sseAlg, string, sseKey string) (url string) {
	input := &s3.GetObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(obj),
		SSECustomerAlgorithm: aws.String(sseAlg),
		SSECustomerKey:       aws.String(sslKey),
	}

	req, _ := s.s3c.GetObjectRequest(input)
	url, err := req.Presign(30 * time.Minute)

	if err != nil {
		panic(fmt.Sprintf("Failed to sign request: %v", err))

	}

	return url
}
