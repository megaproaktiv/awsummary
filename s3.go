package awsummary

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"log"
)

func ListS3(region string, verbose bool) (number int, err error) {
	var bucketNumber int
	svc := s3.New(session.New(&aws.Config{Region: aws.String(region)}))
	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return bucketNumber, err
	}

	if verbose {
		log.Println("Buckets:")
	}
	for _, bucket := range result.Buckets {
		bucketNumber++
		if verbose {
			log.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
		}
	}
	return bucketNumber, err
}
