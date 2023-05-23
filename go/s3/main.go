package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// initialize a session in ap-south-1 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})

	// create s3 service client
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("antis-filestore"),
		Key:    aws.String("test.png"),
	})

	url, err := req.Presign(5 * time.Minute)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
