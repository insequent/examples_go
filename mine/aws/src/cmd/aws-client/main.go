package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"aws/s3client"
)

func main() {
	client := s3client.NewS3Client()
	defer client.Cleanup()

	data := os.Open("files/object.txt")

	epoch := time.Now().Unix()
	bucket := "joshc_test_" + strconv.Itoa(epoch)
	fmt.Println("Creating a new bucket", bucket)
	client.CreateBucket(bucket)

	clearObj := "test_clear_object"
	fmt.Println("Creating a new unencrypted object", clearObj)
	client.CreateObject(bucket, clearObj, data, nil, nil)

	fmt.Println("Creating a presigned URL for the unencrypted object", clearObj)
	url := client.PresignedURL(bucket, clearObj)
	fmt.Println("\tPresigned URL:", url)
	resp, err := http.Get(url)
	fmt.Printf("\tResponse (%v): %v", resp.status, resp.body)

	encObj := "test_encrypted_object"
	sseAlg := "AES256"
	sseKey := make([]byte, 32)
	if _, err := rand.Read(sseKey); err != nil {
		panic("Error attempting to generating random SSE-C Key:", err)
	}
	fmt.Println("Creating a new SSE-C encrypted object", encObj)
	client.CreateObject(bucket, encObj, data, sseAlg, sseKey)

	fmt.Println("Creating a presigned URL for the SSE-C encrypted object", encObj)
	url = client.PresignedURL(bucket, encObj)
	fmt.Println("\tPresigned URL:", url)
	resp, err = http.Get(url)
	fmt.Printf("\tResponse (%v): %v", resp.status, resp.body)
}
