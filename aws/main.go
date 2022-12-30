package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,:;<>!?")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}

	return string(b)
}

func main() {
	client := NewS3Client()
	epoch := time.Now().Unix()
	bucket := "my-test-" + strconv.FormatInt(epoch, 10)
	defer func() {
		fmt.Println("Cleaning up...")
		client.Cleanup(bucket)
	}()

	// Create the bucket
	fmt.Println("Creating a new bucket", bucket)
	client.CreateBucket(bucket)

	// Create a clear object
	clearObj := "test-clear-object"
	data, err := os.Open("files/object.txt")
	if err != nil {
		panic(fmt.Sprintf("Error attempting to read files/object.txt: %v", err))
	}
	defer data.Close()
	fmt.Println("Creating a new unencrypted object", clearObj)
	client.CreateObject(bucket, clearObj, data, nil, nil)

	// Create a presigned URL for clear object
	fmt.Println("Creating a presigned URL for the unencrypted object", clearObj)
	url := client.PresignedURL(bucket, clearObj, nil, nil)
	fmt.Println("\tPresigned URL:", url)

	// Perform GET on pre-signed URL for clear object
	resp, err := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("\tResponse (%v): %s\n", resp.Status, body)

	// Create an encrypted object
	encObj := "test-encrypted-object"
	sseAlg := "AES256"
	sseKey := RandString(32)
	//sseKey := "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456"
	fmt.Println("Created customer key:", sseKey)
	fmt.Println("Creating a new SSE-C encrypted object", encObj)
	data.Seek(0, io.SeekStart) // Reset the file pointer so we can read again
	client.CreateObject(bucket, encObj, data, &sseAlg, &sseKey)

	// Create a presigned URL for encrypted object
	fmt.Println("Creating a presigned URL for the SSE-C encrypted object", encObj)
	url = client.PresignedURL(bucket, encObj, &sseAlg, &sseKey)
	fmt.Println("\tPresigned URL:", url)

	// Perform a GET on pre-signed URL for encrypted object
	// WARNING: Do not convert md5 from hex to string before base64 encoding
	sseMD5 := md5.Sum([]byte(sseKey))
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-amz-server-side-encryption-customer-algorithm", "AES256")
	req.Header.Add("x-amz-server-side-encryption-customer-key", base64.StdEncoding.EncodeToString([]byte(sseKey)))
	req.Header.Add("x-amz-server-side-encryption-customer-key-md5", base64.StdEncoding.EncodeToString(sseMD5[:]))
	for _, header := range req.Header {
		fmt.Println("\t", header)
	}
	resp, _ = http.DefaultClient.Do(req)
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("\tResponse (%v): %s\n", resp.Status, body)
}
