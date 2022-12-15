// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backupgateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the bandwidth rate limit schedule for a specified gateway. By default,
// gateways do not have bandwidth rate limit schedules, which means no bandwidth
// rate limiting is in effect. Use this to get a gateway's bandwidth rate limit
// schedule.
func (c *Client) GetBandwidthRateLimitSchedule(ctx context.Context, params *GetBandwidthRateLimitScheduleInput, optFns ...func(*Options)) (*GetBandwidthRateLimitScheduleOutput, error) {
	if params == nil {
		params = &GetBandwidthRateLimitScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBandwidthRateLimitSchedule", params, optFns, c.addOperationGetBandwidthRateLimitScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBandwidthRateLimitScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBandwidthRateLimitScheduleInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways
	// (https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BGW_ListGateways.html)
	// operation to return a list of gateways for your account and Amazon Web Services
	// Region.
	//
	// This member is required.
	GatewayArn *string

	noSmithyDocumentSerde
}

type GetBandwidthRateLimitScheduleOutput struct {

	// An array containing bandwidth rate limit schedule intervals for a gateway. When
	// no bandwidth rate limit intervals have been scheduled, the array is empty.
	BandwidthRateLimitIntervals []types.BandwidthRateLimitInterval

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways
	// (https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BGW_ListGateways.html)
	// operation to return a list of gateways for your account and Amazon Web Services
	// Region.
	GatewayArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBandwidthRateLimitScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetBandwidthRateLimitSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetBandwidthRateLimitSchedule{}, middleware.After)
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
	if err = addOpGetBandwidthRateLimitScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBandwidthRateLimitSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBandwidthRateLimitSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-gateway",
		OperationName: "GetBandwidthRateLimitSchedule",
	}
}
