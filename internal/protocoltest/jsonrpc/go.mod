module github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.14.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.5
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.3.0
	github.com/aws/smithy-go v1.11.0
	github.com/google/go-cmp v0.5.7
)

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/
