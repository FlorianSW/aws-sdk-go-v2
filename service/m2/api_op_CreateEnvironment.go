// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a runtime environment for a given runtime engine.
func (c *Client) CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) {
	if params == nil {
		params = &CreateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironment", params, optFns, c.addOperationCreateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentInput struct {

	// The engine type for the runtime environment.
	//
	// This member is required.
	EngineType types.EngineType

	// The type of instance for the runtime environment.
	//
	// This member is required.
	InstanceType *string

	// The name of the runtime environment. Must be unique within the account.
	//
	// This member is required.
	Name *string

	// Unique, case-sensitive identifier you provide to ensure the idempotency of the
	// request to create an environment. The service generates the clientToken when the
	// API call is triggered. The token expires after one hour, so if you retry the API
	// within this timeframe with the same clientToken, you will get the same response.
	// The service also handles deleting the clientToken after it expires.
	ClientToken *string

	// The description of the runtime environment.
	Description *string

	// The version of the engine type for the runtime environment.
	EngineVersion *string

	// The details of a high availability configuration for this runtime environment.
	HighAvailabilityConfig *types.HighAvailabilityConfig

	// The identifier of a customer managed key.
	KmsKeyId *string

	// Configures the maintenance window you want for the runtime environment. If you
	// do not provide a value, a random system-generated value will be assigned.
	PreferredMaintenanceWindow *string

	// Specifies whether the runtime environment is publicly accessible.
	PubliclyAccessible bool

	// The list of security groups for the VPC associated with this runtime
	// environment.
	SecurityGroupIds []string

	// Optional. The storage configurations for this runtime environment.
	StorageConfigurations []types.StorageConfiguration

	// The list of subnets associated with the VPC for this runtime environment.
	SubnetIds []string

	// The tags for the runtime environment.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateEnvironmentOutput struct {

	// The unique identifier of the runtime environment.
	//
	// This member is required.
	EnvironmentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEnvironment{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateEnvironmentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironment(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateEnvironment struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateEnvironment) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateEnvironmentInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateEnvironmentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateEnvironment{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "CreateEnvironment",
	}
}
