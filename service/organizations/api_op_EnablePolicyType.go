// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables a policy type in a root. After you enable a policy type in a root, you
// can attach policies of that type to the root, any organizational unit (OU), or
// account in that root. You can undo this by using the DisablePolicyType
// operation. This is an asynchronous request that Amazon Web Services performs in
// the background. Amazon Web Services recommends that you first use ListRoots to
// see the status of policy types for a specified root, and then use this
// operation. This operation can be called only from the organization's management
// account. You can enable a policy type in a root only if that policy type is
// available in the organization. To view the status of available policy types in
// the organization, use DescribeOrganization.
func (c *Client) EnablePolicyType(ctx context.Context, params *EnablePolicyTypeInput, optFns ...func(*Options)) (*EnablePolicyTypeOutput, error) {
	if params == nil {
		params = &EnablePolicyTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnablePolicyType", params, optFns, c.addOperationEnablePolicyTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnablePolicyTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnablePolicyTypeInput struct {

	// The policy type that you want to enable. You can specify one of the following
	// values:
	//
	// * AISERVICES_OPT_OUT_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	//
	// *
	// BACKUP_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html)
	//
	// *
	// SERVICE_CONTROL_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html)
	//
	// *
	// TAG_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
	//
	// This member is required.
	PolicyType types.PolicyType

	// The unique identifier (ID) of the root in which you want to enable a policy
	// type. You can get the ID from the ListRoots operation. The regex pattern
	// (http://wikipedia.org/wiki/regex) for a root ID string requires "r-" followed by
	// from 4 to 32 lowercase letters or digits.
	//
	// This member is required.
	RootId *string

	noSmithyDocumentSerde
}

type EnablePolicyTypeOutput struct {

	// A structure that shows the root with the updated list of enabled policy types.
	Root *types.Root

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnablePolicyTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnablePolicyType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnablePolicyType{}, middleware.After)
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
	if err = addOpEnablePolicyTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnablePolicyType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnablePolicyType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "EnablePolicyType",
	}
}
