// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about an Amazon Kendra data source connector.
func (c *Client) DescribeDataSource(ctx context.Context, params *DescribeDataSourceInput, optFns ...func(*Options)) (*DescribeDataSourceOutput, error) {
	if params == nil {
		params = &DescribeDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataSource", params, optFns, c.addOperationDescribeDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDataSourceInput struct {

	// The identifier of the data source connector.
	//
	// This member is required.
	Id *string

	// The identifier of the index used with the data source connector.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DescribeDataSourceOutput struct {

	// Configuration details for the data source connector. This shows how the data
	// source is configured. The configuration options for a data source depend on the
	// data source provider.
	Configuration *types.DataSourceConfiguration

	// The Unix timestamp of when the data source connector was created.
	CreatedAt *time.Time

	// Configuration information for altering document metadata and content during the
	// document ingestion process when you describe a data source. For more information
	// on how to create, modify and delete document metadata, or make other content
	// alterations when you ingest documents into Amazon Kendra, see Customizing
	// document metadata during the ingestion process
	// (https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html).
	CustomDocumentEnrichmentConfiguration *types.CustomDocumentEnrichmentConfiguration

	// The description for the data source connector.
	Description *string

	// When the Status field value is FAILED, the ErrorMessage field contains a
	// description of the error that caused the data source to fail.
	ErrorMessage *string

	// The identifier of the data source connector.
	Id *string

	// The identifier of the index used with the data source connector.
	IndexId *string

	// The code for a language. This shows a supported language for all documents in
	// the data source. English is supported by default. For more information on
	// supported languages, including their codes, see Adding documents in languages
	// other than English
	// (https://docs.aws.amazon.com/kendra/latest/dg/in-adding-languages.html).
	LanguageCode *string

	// The name for the data source connector.
	Name *string

	// The Amazon Resource Name (ARN) of the role with permission to access the data
	// source and required resources.
	RoleArn *string

	// The schedule for Amazon Kendra to update the index.
	Schedule *string

	// The current status of the data source connector. When the status is ACTIVE the
	// data source is ready to use. When the status is FAILED, the ErrorMessage field
	// contains the reason that the data source failed.
	Status types.DataSourceStatus

	// The type of the data source. For example, SHAREPOINT.
	Type types.DataSourceType

	// The Unix timestamp of when the data source connector was last updated.
	UpdatedAt *time.Time

	// Configuration information for an Amazon Virtual Private Cloud to connect to your
	// data source. For more information, see Configuring a VPC
	// (https://docs.aws.amazon.com/kendra/latest/dg/vpc-configuration.html).
	VpcConfiguration *types.DataSourceVpcConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDataSource{}, middleware.After)
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
	if err = addOpDescribeDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DescribeDataSource",
	}
}
