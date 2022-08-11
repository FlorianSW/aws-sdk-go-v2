// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a mission profile. dataflowEdges is a list of lists of strings. Each
// lower level list of strings has two elements: a from ARN and a to ARN.
func (c *Client) CreateMissionProfile(ctx context.Context, params *CreateMissionProfileInput, optFns ...func(*Options)) (*CreateMissionProfileOutput, error) {
	if params == nil {
		params = &CreateMissionProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMissionProfile", params, optFns, c.addOperationCreateMissionProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMissionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMissionProfileInput struct {

	// A list of lists of ARNs. Each list of ARNs is an edge, with a from Config and a
	// to Config.
	//
	// This member is required.
	DataflowEdges [][]string

	// Smallest amount of time in seconds that you’d like to see for an available
	// contact. AWS Ground Station will not present you with contacts shorter than this
	// duration.
	//
	// This member is required.
	MinimumViableContactDurationSeconds *int32

	// Name of a mission profile.
	//
	// This member is required.
	Name *string

	// ARN of a tracking Config.
	//
	// This member is required.
	TrackingConfigArn *string

	// Amount of time after a contact ends that you’d like to receive a CloudWatch
	// event indicating the pass has finished.
	ContactPostPassDurationSeconds *int32

	// Amount of time prior to contact start you’d like to receive a CloudWatch event
	// indicating an upcoming pass.
	ContactPrePassDurationSeconds *int32

	// Tags assigned to a mission profile.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateMissionProfileOutput struct {

	// UUID of a mission profile.
	MissionProfileId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMissionProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMissionProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMissionProfile{}, middleware.After)
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
	if err = addOpCreateMissionProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMissionProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMissionProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "groundstation",
		OperationName: "CreateMissionProfile",
	}
}
