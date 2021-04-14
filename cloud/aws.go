package cloud

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Aws struct {
	sess *session.Session
}

func NewAws() *Aws {
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(os.Getenv("AWS_REGION")),
	}))

	return &Aws{
		sess,
	}
}

func (a *Aws) UploadToS3(key string, file *os.File) error {
	// sess := getAwsSession()
	uploader := s3manager.NewUploader(a.sess)

	// resultはlocationをもっているので、アップロード後のURLを取得できるが、時限付きではないと思うのでつかえない
	if file == nil {
		dir, _ := os.Getwd()
		var err error
		if file, err = os.Open(dir + "/pkg/cloud/asset/default.png"); err != nil {
			return err
		}
	}
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *Aws) GetURLToPutInS3(key string) (string, error) {
	svc := s3.New(a.sess)
	r, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(key),
	})
	url, err := r.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (a *Aws) GetURLToGetInS3(key string) (string, error) {
	svc := s3.New(a.sess)
	r, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(key),
	})
	url, err := r.Presign(30 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil
}
