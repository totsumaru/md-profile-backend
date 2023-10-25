package cloudflare

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// cloudflareに画像をアップロードします
func Upload() error {
	var bucketName = "profio"
	var accessKeyId = os.Getenv("CLOUDFLARE_S3_ACCESS_KEY")
	var accessKeySecret = os.Getenv("CLOUDFLARE_S3_SECRET_ACCESS_KEY")

	var f = func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: os.Getenv("CLOUDFLARE_S3_API_URL")}, nil
	}

	r2Resolver := aws.EndpointResolverWithOptionsFunc(f)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				accessKeyId, accessKeySecret, "",
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	ff, err := os.Open("./test-image.png")
	if err != nil {
		panic(err)
	}

	// MIMEタイプの検出
	buffer := make([]byte, 512) // ヘッダの最大長
	_, err = ff.Read(buffer)
	if err != nil {
		panic(err)
	}
	contentType := http.DetectContentType(buffer)
	_, err = ff.Seek(0, 0) // ファイルのポインタを先頭に戻す
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)
	param := s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String("test-image2.png"),
		Body:               ff,
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("inline"),
	}
	_, err = client.PutObject(context.Background(), &param)
	if err != nil {
		panic(err)
	}

	return nil
}
