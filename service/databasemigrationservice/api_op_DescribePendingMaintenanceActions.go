// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For internal use only
func (c *Client) DescribePendingMaintenanceActions(ctx context.Context, params *DescribePendingMaintenanceActionsInput, optFns ...func(*Options)) (*DescribePendingMaintenanceActionsOutput, error) {
	if params == nil {
		params = &DescribePendingMaintenanceActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePendingMaintenanceActions", params, optFns, c.addOperationDescribePendingMaintenanceActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePendingMaintenanceActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePendingMaintenanceActionsInput struct {

	//
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn *string

	noSmithyDocumentSerde
}

type DescribePendingMaintenanceActionsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The pending maintenance action.
	PendingMaintenanceActions []types.ResourcePendingMaintenanceActions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePendingMaintenanceActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePendingMaintenanceActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePendingMaintenanceActions{}, middleware.After)
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
	if err = addOpDescribePendingMaintenanceActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePendingMaintenanceActions(options.Region), middleware.Before); err != nil {
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

// DescribePendingMaintenanceActionsAPIClient is a client that implements the
// DescribePendingMaintenanceActions operation.
type DescribePendingMaintenanceActionsAPIClient interface {
	DescribePendingMaintenanceActions(context.Context, *DescribePendingMaintenanceActionsInput, ...func(*Options)) (*DescribePendingMaintenanceActionsOutput, error)
}

var _ DescribePendingMaintenanceActionsAPIClient = (*Client)(nil)

// DescribePendingMaintenanceActionsPaginatorOptions is the paginator options for
// DescribePendingMaintenanceActions
type DescribePendingMaintenanceActionsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePendingMaintenanceActionsPaginator is a paginator for
// DescribePendingMaintenanceActions
type DescribePendingMaintenanceActionsPaginator struct {
	options   DescribePendingMaintenanceActionsPaginatorOptions
	client    DescribePendingMaintenanceActionsAPIClient
	params    *DescribePendingMaintenanceActionsInput
	nextToken *string
	firstPage bool
}

// NewDescribePendingMaintenanceActionsPaginator returns a new
// DescribePendingMaintenanceActionsPaginator
func NewDescribePendingMaintenanceActionsPaginator(client DescribePendingMaintenanceActionsAPIClient, params *DescribePendingMaintenanceActionsInput, optFns ...func(*DescribePendingMaintenanceActionsPaginatorOptions)) *DescribePendingMaintenanceActionsPaginator {
	if params == nil {
		params = &DescribePendingMaintenanceActionsInput{}
	}

	options := DescribePendingMaintenanceActionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePendingMaintenanceActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePendingMaintenanceActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePendingMaintenanceActions page.
func (p *DescribePendingMaintenanceActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePendingMaintenanceActionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribePendingMaintenanceActions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribePendingMaintenanceActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribePendingMaintenanceActions",
	}
}
