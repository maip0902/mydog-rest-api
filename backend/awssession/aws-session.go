package awssession

import (
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
)

func StartSession() (sess *session.Session, err error) {
	AccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
    SecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
    MyRegion := os.Getenv("AWS_REGION")
    sess, err = session.NewSession(
     &aws.Config{
      Endpoint: aws.String(os.Getenv("S3_ENDPOINT")),
      Region: aws.String(MyRegion),
      S3ForcePathStyle: aws.Bool(true),
      Credentials: credentials.NewStaticCredentials(
       AccessKeyID,
       SecretAccessKey,
       "", 
      ),
	})
	return sess,err
}