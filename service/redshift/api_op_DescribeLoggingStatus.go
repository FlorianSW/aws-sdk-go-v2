// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes whether information, such as queries and connection attempts, is being
// logged for the specified Amazon Redshift cluster.
func (c *Client) DescribeLoggingStatus(ctx context.Context, params *DescribeLoggingStatusInput, optFns ...func(*Options)) (*DescribeLoggingStatusOutput, error) {
	if params == nil {
		params = &DescribeLoggingStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoggingStatus", params, optFns, c.addOperationDescribeLoggingStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoggingStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLoggingStatusInput struct {

	// The identifier of the cluster from which to get the logging status. Example:
	// examplecluster
	//
	// This member is required.
	ClusterIdentifier *string

	noSmithyDocumentSerde
}

// Describes the status of logging for a cluster.
type DescribeLoggingStatusOutput struct {

	// The name of the S3 bucket where the log files are stored.
	BucketName *string

	// The message indicating that logs failed to be delivered.
	LastFailureMessage *string

	// The last time when logs failed to be delivered.
	LastFailureTime *time.Time

	// The last time that logs were delivered.
	LastSuccessfulDeliveryTime *time.Time

	// The log destination type. An enum with possible values of s3 and cloudwatch.
	LogDestinationType types.LogDestinationType

	// The collection of exported log types. Log types include the connection log, user
	// log and user activity log.
	LogExports []string

	// true if logging is on, false if logging is off.
	LoggingEnabled bool

	// The prefix applied to the log file names.
	S3KeyPrefix *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLoggingStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoggingStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoggingStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeLoggingStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoggingStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeLoggingStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeLoggingStatus",
	}
}
