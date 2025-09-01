package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Buckt  string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			"your-access-key-id",     // replace with your access key id
			"your-secret-access-key", // replace with your secret access key
			"",                       // a token will be created when the session it's used.
		),
	})

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Buckt = "your-bucket-name"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	// control the number of concurrent uploads
	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)

	// Retry mechanism for failed uploads
	go func() {
		for {
			select {
			case fileName := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(fileName, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.Readdir(1)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		wg.Add(1)
		// acquire a slot
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}

	wg.Wait()

}

func uploadFile(fileName string, uploadControl <-chan struct{}, errorFileUpload chan<- string) error {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", fileName)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: \n", completeFileName)
		<-uploadControl // esvazia o canal
		errorFileUpload <- completeFileName
		return err
	}

	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Buckt),
		Key:    aws.String(fileName),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s: \n", completeFileName)
		<-uploadControl
	}

	fmt.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl
	return nil
}
