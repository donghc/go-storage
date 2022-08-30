module github.com/beyondstorage/go-storage/services/s3/v3

go 1.16

require (
	github.com/aws/aws-sdk-go-v2 v1.16.12
	github.com/aws/aws-sdk-go-v2/config v1.17.1
	github.com/aws/aws-sdk-go-v2/credentials v1.12.15
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.5
	github.com/aws/smithy-go v1.13.0
	github.com/beyondstorage/go-storage/credential v1.0.0
	github.com/beyondstorage/go-storage/endpoint v1.2.0
	github.com/beyondstorage/go-storage/v5 v5.0.0
	github.com/google/uuid v1.3.0
)

replace (
	github.com/beyondstorage/go-storage/credential => ../../credential
	github.com/beyondstorage/go-storage/endpoint => ../../endpoint
	github.com/beyondstorage/go-storage/v5 => ../../
)
