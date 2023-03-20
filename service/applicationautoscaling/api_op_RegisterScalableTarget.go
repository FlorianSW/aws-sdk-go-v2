// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationautoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers or updates a scalable target, which is the resource that you want to
// scale. Scalable targets are uniquely identified by the combination of resource
// ID, scalable dimension, and namespace, which represents some capacity dimension
// of the underlying service. When you register a new scalable target, you must
// specify values for the minimum and maximum capacity. If the specified resource
// is not active in the target service, this operation does not change the
// resource's current capacity. Otherwise, it changes the resource's current
// capacity to a value that is inside of this range. If you add a scaling policy,
// current capacity is adjustable within the specified range when scaling starts.
// Application Auto Scaling scaling policies will not scale capacity to values that
// are outside of the minimum and maximum range. After you register a scalable
// target, you do not need to register it again to use other Application Auto
// Scaling operations. To see which resources have been registered, use
// DescribeScalableTargets
// (https://docs.aws.amazon.com/autoscaling/application/APIReference/API_DescribeScalableTargets.html).
// You can also view the scaling policies for a service namespace by using
// DescribeScalableTargets
// (https://docs.aws.amazon.com/autoscaling/application/APIReference/API_DescribeScalableTargets.html).
// If you no longer need a scalable target, you can deregister it by using
// DeregisterScalableTarget
// (https://docs.aws.amazon.com/autoscaling/application/APIReference/API_DeregisterScalableTarget.html).
// To update a scalable target, specify the parameters that you want to change.
// Include the parameters that identify the scalable target: resource ID, scalable
// dimension, and namespace. Any parameters that you don't specify are not changed
// by this update request. If you call the RegisterScalableTarget API operation to
// create a scalable target, there might be a brief delay until the operation
// achieves eventual consistency
// (https://en.wikipedia.org/wiki/Eventual_consistency). You might become aware of
// this brief delay if you get unexpected errors when performing sequential
// operations. The typical strategy is to retry the request, and some Amazon Web
// Services SDKs include automatic backoff and retry logic. If you call the
// RegisterScalableTarget API operation to update an existing scalable target,
// Application Auto Scaling retrieves the current capacity of the resource. If it's
// below the minimum capacity or above the maximum capacity, Application Auto
// Scaling adjusts the capacity of the scalable target to place it within these
// bounds, even if you don't include the MinCapacity or MaxCapacity request
// parameters.
func (c *Client) RegisterScalableTarget(ctx context.Context, params *RegisterScalableTargetInput, optFns ...func(*Options)) (*RegisterScalableTargetOutput, error) {
	if params == nil {
		params = &RegisterScalableTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterScalableTarget", params, optFns, c.addOperationRegisterScalableTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterScalableTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterScalableTargetInput struct {

	// The identifier of the resource that is associated with the scalable target. This
	// string consists of the resource type and unique identifier.
	//
	// * ECS service - The
	// resource type is service and the unique identifier is the cluster name and
	// service name. Example: service/default/sample-webapp.
	//
	// * Spot Fleet - The
	// resource type is spot-fleet-request and the unique identifier is the Spot Fleet
	// request ID. Example:
	// spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	// * EMR cluster -
	// The resource type is instancegroup and the unique identifier is the cluster ID
	// and instance group ID. Example:
	// instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0.
	//
	// * AppStream 2.0 fleet - The
	// resource type is fleet and the unique identifier is the fleet name. Example:
	// fleet/sample-fleet.
	//
	// * DynamoDB table - The resource type is table and the
	// unique identifier is the table name. Example: table/my-table.
	//
	// * DynamoDB global
	// secondary index - The resource type is index and the unique identifier is the
	// index name. Example: table/my-table/index/my-table-index.
	//
	// * Aurora DB cluster -
	// The resource type is cluster and the unique identifier is the cluster name.
	// Example: cluster:my-db-cluster.
	//
	// * SageMaker endpoint variant - The resource
	// type is variant and the unique identifier is the resource ID. Example:
	// endpoint/my-end-point/variant/KMeansClustering.
	//
	// * Custom resources are not
	// supported with a resource type. This parameter must specify the OutputValue from
	// the CloudFormation template stack used to access the resources. The unique
	// identifier is defined by the service provider. More information is available in
	// our GitHub repository
	// (https://github.com/aws/aws-auto-scaling-custom-resource).
	//
	// * Amazon Comprehend
	// document classification endpoint - The resource type and unique identifier are
	// specified using the endpoint ARN. Example:
	// arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE.
	//
	// *
	// Amazon Comprehend entity recognizer endpoint - The resource type and unique
	// identifier are specified using the endpoint ARN. Example:
	// arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE.
	//
	// *
	// Lambda provisioned concurrency - The resource type is function and the unique
	// identifier is the function name with a function version or alias name suffix
	// that is not $LATEST. Example: function:my-function:prod or
	// function:my-function:1.
	//
	// * Amazon Keyspaces table - The resource type is table
	// and the unique identifier is the table name. Example:
	// keyspace/mykeyspace/table/mytable.
	//
	// * Amazon MSK cluster - The resource type and
	// unique identifier are specified using the cluster ARN. Example:
	// arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5.
	//
	// *
	// Amazon ElastiCache replication group - The resource type is replication-group
	// and the unique identifier is the replication group name. Example:
	// replication-group/mycluster.
	//
	// * Neptune cluster - The resource type is cluster
	// and the unique identifier is the cluster name. Example: cluster:mycluster.
	//
	// This member is required.
	ResourceId *string

	// The scalable dimension associated with the scalable target. This string consists
	// of the service namespace, resource type, and scaling property.
	//
	// *
	// ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	// *
	// elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR
	// Instance Group.
	//
	// * ec2:spot-fleet-request:TargetCapacity - The target capacity
	// of a Spot Fleet.
	//
	// * appstream:fleet:DesiredCapacity - The desired capacity of an
	// AppStream 2.0 fleet.
	//
	// * dynamodb:table:ReadCapacityUnits - The provisioned read
	// capacity for a DynamoDB table.
	//
	// * dynamodb:table:WriteCapacityUnits - The
	// provisioned write capacity for a DynamoDB table.
	//
	// *
	// dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a DynamoDB
	// global secondary index.
	//
	// * dynamodb:index:WriteCapacityUnits - The provisioned
	// write capacity for a DynamoDB global secondary index.
	//
	// *
	// rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB
	// cluster. Available for Aurora MySQL-compatible edition and Aurora
	// PostgreSQL-compatible edition.
	//
	// * sagemaker:variant:DesiredInstanceCount - The
	// number of EC2 instances for a SageMaker model endpoint variant.
	//
	// *
	// custom-resource:ResourceType:Property - The scalable dimension for a custom
	// resource provided by your own application or service.
	//
	// *
	// comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number of
	// inference units for an Amazon Comprehend document classification endpoint.
	//
	// *
	// comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of
	// inference units for an Amazon Comprehend entity recognizer endpoint.
	//
	// *
	// lambda:function:ProvisionedConcurrency - The provisioned concurrency for a
	// Lambda function.
	//
	// * cassandra:table:ReadCapacityUnits - The provisioned read
	// capacity for an Amazon Keyspaces table.
	//
	// * cassandra:table:WriteCapacityUnits -
	// The provisioned write capacity for an Amazon Keyspaces table.
	//
	// *
	// kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for
	// brokers in an Amazon MSK cluster.
	//
	// * elasticache:replication-group:NodeGroups -
	// The number of node groups for an Amazon ElastiCache replication group.
	//
	// *
	// elasticache:replication-group:Replicas - The number of replicas per node group
	// for an Amazon ElastiCache replication group.
	//
	// * neptune:cluster:ReadReplicaCount
	// - The count of read replicas in an Amazon Neptune DB cluster.
	//
	// This member is required.
	ScalableDimension types.ScalableDimension

	// The namespace of the Amazon Web Services service that provides the resource. For
	// a resource provided by your own application or service, use custom-resource
	// instead.
	//
	// This member is required.
	ServiceNamespace types.ServiceNamespace

	// The maximum value that you plan to scale out to. When a scaling policy is in
	// effect, Application Auto Scaling can scale out (expand) as needed to the maximum
	// capacity limit in response to changing demand. This property is required when
	// registering a new scalable target. Although you can specify a large maximum
	// capacity, note that service quotas might impose lower limits. Each service has
	// its own default quotas for the maximum capacity of the resource. If you want to
	// specify a higher limit, you can request an increase. For more information,
	// consult the documentation for that service. For information about the default
	// quotas for each service, see Service endpoints and quotas
	// (https://docs.aws.amazon.com/general/latest/gr/aws-service-information.html) in
	// the Amazon Web Services General Reference.
	MaxCapacity *int32

	// The minimum value that you plan to scale in to. When a scaling policy is in
	// effect, Application Auto Scaling can scale in (contract) as needed to the
	// minimum capacity limit in response to changing demand. This property is required
	// when registering a new scalable target. For the following resources, the minimum
	// value allowed is 0.
	//
	// * AppStream 2.0 fleets
	//
	// * Aurora DB clusters
	//
	// * ECS
	// services
	//
	// * EMR clusters
	//
	// * Lambda provisioned concurrency
	//
	// * SageMaker endpoint
	// variants
	//
	// * Spot Fleets
	//
	// * custom resources
	//
	// It's strongly recommended that you
	// specify a value greater than 0. A value greater than 0 means that data points
	// are continuously reported to CloudWatch that scaling policies can use to scale
	// on a metric like average CPU utilization. For all other resources, the minimum
	// allowed value depends on the type of resource that you are using. If you provide
	// a value that is lower than what a resource can accept, an error occurs. In which
	// case, the error message will provide the minimum value that the resource can
	// accept.
	MinCapacity *int32

	// This parameter is required for services that do not support service-linked roles
	// (such as Amazon EMR), and it must specify the ARN of an IAM role that allows
	// Application Auto Scaling to modify the scalable target on your behalf. If the
	// service supports service-linked roles, Application Auto Scaling uses a
	// service-linked role, which it creates if it does not yet exist. For more
	// information, see Application Auto Scaling IAM roles
	// (https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-roles).
	RoleARN *string

	// An embedded object that contains attributes and attribute values that are used
	// to suspend and resume automatic scaling. Setting the value of an attribute to
	// true suspends the specified scaling activities. Setting it to false (default)
	// resumes the specified scaling activities. Suspension Outcomes
	//
	// * For
	// DynamicScalingInSuspended, while a suspension is in effect, all scale-in
	// activities that are triggered by a scaling policy are suspended.
	//
	// * For
	// DynamicScalingOutSuspended, while a suspension is in effect, all scale-out
	// activities that are triggered by a scaling policy are suspended.
	//
	// * For
	// ScheduledScalingSuspended, while a suspension is in effect, all scaling
	// activities that involve scheduled actions are suspended.
	//
	// For more information,
	// see Suspending and resuming scaling
	// (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html)
	// in the Application Auto Scaling User Guide.
	SuspendedState *types.SuspendedState

	// Assigns one or more tags to the scalable target. Use this parameter to tag the
	// scalable target when it is created. To tag an existing scalable target, use the
	// TagResource operation. Each tag consists of a tag key and a tag value. Both the
	// tag key and the tag value are required. You cannot have more than one tag on a
	// scalable target with the same tag key. Use tags to control access to a scalable
	// target. For more information, see Tagging support for Application Auto Scaling
	// (https://docs.aws.amazon.com/autoscaling/application/userguide/resource-tagging-support.html)
	// in the Application Auto Scaling User Guide.
	Tags map[string]string

	noSmithyDocumentSerde
}

type RegisterScalableTargetOutput struct {

	// The ARN of the scalable target.
	ScalableTargetARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterScalableTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterScalableTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterScalableTarget{}, middleware.After)
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
	if err = addOpRegisterScalableTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterScalableTarget(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterScalableTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "application-autoscaling",
		OperationName: "RegisterScalableTarget",
	}
}
