package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		println("Usage: go run test_bucket_owner.go ACCOUNT_ID BUCKET_NAME OBJECT_KEY")
		os.Exit(1)
	}

	account := os.Args[1]
	bucket := os.Args[2]
	key := os.Args[3]
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if err != nil {
		panic(err)
	}
	c := s3.NewFromConfig(cfg)
	p := s3.NewPresignClient(c)
	i := s3.GetObjectInput{
		ExpectedBucketOwner: aws.String(account),
		Bucket:              aws.String(bucket),
		Key:                 aws.String(key),
	}
	u, err := p.PresignGetObject(ctx, &i)
	if err != nil {
		panic(err)
	}
	println(u.URL)
	for k, vs := range u.SignedHeader {
		fmt.Printf("%s: %s\n", k, strings.Join(vs, ","))
	}
	println(u.Method)
}
