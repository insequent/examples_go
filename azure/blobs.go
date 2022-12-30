package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// From the Storage Account: Access Keys in the Azure Portal
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT")
	accountKey := os.Getenv("AZURE_STORAGE_ACCESS_KEY")

	// Create a default request pipeline
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Either the AZURE_STORAGE_ACCOUNT or AZURE_STORAGE_ACCESS_KEY " +
			"environment variables are not set correctly")
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// Need to ensure the container exists
	containerName := "test-container"
	u, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))
	containerUrl := azblob.NewContainerURL(*u, p)
	fmt.Printf("Checking on container named %s\n", containerName)
	ctx := context.Background()
	_, err = containerUrl.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
	if err != nil {
		if serr, ok := err.(azblob.StorageError); ok {
			if serr.ServiceCode() == azblob.ServiceCodeContainerAlreadyExists {
				fmt.Printf("Containter named %s already exists. Continuing...\n", containerName)
			}
		} else {
			handleErrors(err)
		}
	}

	// Prepare a blob
	blob := []byte("I'm a blob!")
	fileName := "test-file"
	err = ioutil.WriteFile(fileName, blob, 0700)
	handleErrors(err)

	// Upload our blob
	blobUrl := containerUrl.NewBlockBlobURL(fileName)
	file, err := os.Open(fileName)
	handleErrors(err)

	fmt.Printf("Uploading the file with blob name: %s\n", fileName)
	_, err = azblob.UploadFileToBlockBlob(ctx, file, blobUrl, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16})
	handleErrors(err)

	// Download our blob
	resp, err := blobUrl.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	handleErrors(err)

	bodyStream := resp.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})
	dbuff := bytes.Buffer{}
	_, err = dbuff.ReadFrom(bodyStream)
	fmt.Println("Downloaded: " + dbuff.String())

	// Print out our blob!
	fmt.Printf("This is your blob: %v\n", blob)
}
