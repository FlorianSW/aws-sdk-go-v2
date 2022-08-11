// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Refreshes the Trusted Advisor check that you specify using the check ID. You can
// get the check IDs by calling the DescribeTrustedAdvisorChecks operation. Some
// checks are refreshed automatically. If you call the RefreshTrustedAdvisorCheck
// operation to refresh them, you might see the InvalidParameterValue error. The
// response contains a TrustedAdvisorCheckRefreshStatus object.
//
// * You must have a
// Business, Enterprise On-Ramp, or Enterprise Support plan to use the Amazon Web
// Services Support API.
//
// * If you call the Amazon Web Services Support API from an
// account that does not have a Business, Enterprise On-Ramp, or Enterprise Support
// plan, the SubscriptionRequiredException error message appears. For information
// about changing your support plan, see Amazon Web Services Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) RefreshTrustedAdvisorCheck(ctx context.Context, params *RefreshTrustedAdvisorCheckInput, optFns ...func(*Options)) (*RefreshTrustedAdvisorCheckOutput, error) {
	if params == nil {
		params = &RefreshTrustedAdvisorCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RefreshTrustedAdvisorCheck", params, optFns, c.addOperationRefreshTrustedAdvisorCheckMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RefreshTrustedAdvisorCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RefreshTrustedAdvisorCheckInput struct {

	// The unique identifier for the Trusted Advisor check to refresh. Specifying the
	// check ID of a check that is automatically refreshed causes an
	// InvalidParameterValue error.
	//
	// This member is required.
	CheckId *string

	noSmithyDocumentSerde
}

// The current refresh status of a Trusted Advisor check.
type RefreshTrustedAdvisorCheckOutput struct {

	// The current refresh status for a check, including the amount of time until the
	// check is eligible for refresh.
	//
	// This member is required.
	Status *types.TrustedAdvisorCheckRefreshStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRefreshTrustedAdvisorCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRefreshTrustedAdvisorCheck{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRefreshTrustedAdvisorCheck{}, middleware.After)
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
	if err = addOpRefreshTrustedAdvisorCheckValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRefreshTrustedAdvisorCheck(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRefreshTrustedAdvisorCheck(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "RefreshTrustedAdvisorCheck",
	}
}
