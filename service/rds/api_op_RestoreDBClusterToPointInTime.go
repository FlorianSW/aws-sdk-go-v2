// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restores a DB cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target DB cluster is created from the source DB cluster with the same
// configuration as the original DB cluster, except that the new DB cluster is
// created with the default DB security group. For Aurora, this action only
// restores the DB cluster, not the DB instances for that DB cluster. You must
// invoke the CreateDBInstance action to create DB instances for the restored DB
// cluster, specifying the identifier of the restored DB cluster in
// DBClusterIdentifier. You can create DB instances only after the
// RestoreDBClusterToPointInTime action has completed and the DB cluster is
// available. For more information on Amazon Aurora DB clusters, see  What is
// Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see  Multi-AZ deployments with two readable standby DB instances
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) RestoreDBClusterToPointInTime(ctx context.Context, params *RestoreDBClusterToPointInTimeInput, optFns ...func(*Options)) (*RestoreDBClusterToPointInTimeOutput, error) {
	if params == nil {
		params = &RestoreDBClusterToPointInTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBClusterToPointInTime", params, optFns, c.addOperationRestoreDBClusterToPointInTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBClusterToPointInTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreDBClusterToPointInTimeInput struct {

	// The name of the new DB cluster to be created. Constraints:
	//
	// * Must contain from
	// 1 to 63 letters, numbers, or hyphens
	//
	// * First character must be a letter
	//
	// *
	// Can't end with a hyphen or contain two consecutive hyphens
	//
	// Valid for: Aurora DB
	// clusters and Multi-AZ DB clusters
	//
	// This member is required.
	DBClusterIdentifier *string

	// The identifier of the source DB cluster from which to restore. Constraints:
	//
	// *
	// Must match the identifier of an existing DBCluster.
	//
	// Valid for: Aurora DB
	// clusters and Multi-AZ DB clusters
	//
	// This member is required.
	SourceDBClusterIdentifier *string

	// The target backtrack window, in seconds. To disable backtracking, set this value
	// to 0. Default: 0 Constraints:
	//
	// * If specified, this value must be set to a
	// number from 0 to 259,200 (72 hours).
	//
	// Valid for: Aurora MySQL DB clusters only
	BacktrackWindow *int64

	// A value that indicates whether to copy all tags from the restored DB cluster to
	// snapshots of the restored DB cluster. The default is not to copy them. Valid
	// for: Aurora DB clusters and Multi-AZ DB clusters
	CopyTagsToSnapshot *bool

	// The compute and memory capacity of the each DB instance in the Multi-AZ DB
	// cluster, for example db.m6g.xlarge. Not all DB instance classes are available in
	// all Amazon Web Services Regions, or for all database engines. For the full list
	// of DB instance classes, and availability for your engine, see DB instance class
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. Valid for: Multi-AZ DB clusters only
	DBClusterInstanceClass *string

	// The name of the DB cluster parameter group to associate with this DB cluster. If
	// this argument is omitted, the default DB cluster parameter group for the
	// specified engine is used. Constraints:
	//
	// * If supplied, must match the name of an
	// existing DB cluster parameter group.
	//
	// * Must be 1 to 255 letters, numbers, or
	// hyphens.
	//
	// * First character must be a letter.
	//
	// * Can't end with a hyphen or
	// contain two consecutive hyphens.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	DBClusterParameterGroupName *string

	// The DB subnet group name to use for the new DB cluster. Constraints: If
	// supplied, must match the name of an existing DBSubnetGroup. Example:
	// mydbsubnetgroup Valid for: Aurora DB clusters and Multi-AZ DB clusters
	DBSubnetGroupName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection isn't enabled. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	DeletionProtection *bool

	// Specify the Active Directory directory ID to restore the DB cluster in. The
	// domain must be created prior to this operation. For Amazon Aurora DB clusters,
	// Amazon RDS can use Kerberos Authentication to authenticate users that connect to
	// the DB cluster. For more information, see Kerberos Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service. Valid for: Aurora DB clusters only
	DomainIAMRoleName *string

	// The list of logs that the restored DB cluster is to export to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. RDS for MySQL
	// Possible values are error, general, and slowquery. RDS for PostgreSQL Possible
	// values are postgresql and upgrade. Aurora MySQL Possible values are audit,
	// error, general, and slowquery. Aurora PostgreSQL Possible value is postgresql.
	// For more information about exporting CloudWatch Logs for Amazon RDS, see
	// Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide. For more information about exporting CloudWatch
	// Logs for Amazon Aurora, see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	EnableCloudwatchLogsExports []string

	// A value that indicates whether to enable mapping of Amazon Web Services Identity
	// and Access Management (IAM) accounts to database accounts. By default, mapping
	// isn't enabled. For more information, see  IAM Database Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	EnableIAMDatabaseAuthentication *bool

	// The engine mode of the new cluster. Specify provisioned or serverless, depending
	// on the type of the cluster you are creating. You can create an Aurora Serverless
	// v1 clone from a provisioned cluster, or a provisioned clone from an Aurora
	// Serverless v1 cluster. To create a clone that is an Aurora Serverless v1
	// cluster, the original cluster must be an Aurora Serverless v1 cluster or an
	// encrypted provisioned cluster. Valid for: Aurora DB clusters only
	EngineMode *string

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for each DB instance in the Multi-AZ DB cluster. For
	// information about valid Iops values, see Amazon RDS Provisioned IOPS storage to
	// improve performance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide. Constraints: Must be a multiple between .5 and 50
	// of the storage amount for the DB instance. Valid for: Multi-AZ DB clusters only
	Iops *int32

	// The Amazon Web Services KMS key identifier to use when restoring an encrypted DB
	// cluster from an encrypted DB cluster. The Amazon Web Services KMS key identifier
	// is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS
	// key in a different Amazon Web Services account, specify the key ARN or alias
	// ARN. You can restore to a new DB cluster and encrypt the new DB cluster with a
	// KMS key that is different from the KMS key used to encrypt the source DB
	// cluster. The new DB cluster is encrypted with the KMS key identified by the
	// KmsKeyId parameter. If you don't specify a value for the KmsKeyId parameter,
	// then the following occurs:
	//
	// * If the DB cluster is encrypted, then the restored
	// DB cluster is encrypted using the KMS key that was used to encrypt the source DB
	// cluster.
	//
	// * If the DB cluster isn't encrypted, then the restored DB cluster
	// isn't encrypted.
	//
	// If DBClusterIdentifier refers to a DB cluster that isn't
	// encrypted, then the restore request is rejected. Valid for: Aurora DB clusters
	// and Multi-AZ DB clusters
	KmsKeyId *string

	// The name of the option group for the new DB cluster. DB clusters are associated
	// with a default option group that can't be modified.
	OptionGroupName *string

	// The port number on which the new DB cluster accepts connections. Constraints: A
	// value from 1150-65535. Default: The default port for the engine. Valid for:
	// Aurora DB clusters and Multi-AZ DB clusters
	Port *int32

	// A value that indicates whether the DB cluster is publicly accessible. When the
	// DB cluster is publicly accessible, its Domain Name System (DNS) endpoint
	// resolves to the private IP address from within the DB cluster's virtual private
	// cloud (VPC). It resolves to the public IP address from outside of the DB
	// cluster's VPC. Access to the DB cluster is ultimately controlled by the security
	// group it uses. That public access is not permitted if the security group
	// assigned to the DB cluster doesn't permit it. When the DB cluster isn't publicly
	// accessible, it is an internal DB cluster with a DNS name that resolves to a
	// private IP address. Default: The default behavior varies depending on whether
	// DBSubnetGroupName is specified. If DBSubnetGroupName isn't specified, and
	// PubliclyAccessible isn't specified, the following applies:
	//
	// * If the default VPC
	// in the target Region doesn’t have an internet gateway attached to it, the DB
	// cluster is private.
	//
	// * If the default VPC in the target Region has an internet
	// gateway attached to it, the DB cluster is public.
	//
	// If DBSubnetGroupName is
	// specified, and PubliclyAccessible isn't specified, the following applies:
	//
	// * If
	// the subnets are part of a VPC that doesn’t have an internet gateway attached to
	// it, the DB cluster is private.
	//
	// * If the subnets are part of a VPC that has an
	// internet gateway attached to it, the DB cluster is public.
	//
	// Valid for: Multi-AZ
	// DB clusters only
	PubliclyAccessible *bool

	// The date and time to restore the DB cluster to. Valid Values: Value must be a
	// time in Universal Coordinated Time (UTC) format Constraints:
	//
	// * Must be before
	// the latest restorable time for the DB instance
	//
	// * Must be specified if
	// UseLatestRestorableTime parameter isn't provided
	//
	// * Can't be specified if the
	// UseLatestRestorableTime parameter is enabled
	//
	// * Can't be specified if the
	// RestoreType parameter is copy-on-write
	//
	// Example: 2015-03-07T23:45:00Z Valid for:
	// Aurora DB clusters and Multi-AZ DB clusters
	RestoreToTime *time.Time

	// The type of restore to be performed. You can specify one of the following
	// values:
	//
	// * full-copy - The new DB cluster is restored as a full copy of the
	// source DB cluster.
	//
	// * copy-on-write - The new DB cluster is restored as a clone
	// of the source DB cluster.
	//
	// Constraints: You can't specify copy-on-write if the
	// engine version of the source DB cluster is earlier than 1.11. If you don't
	// specify a RestoreType value, then the new DB cluster is restored as a full copy
	// of the source DB cluster. Valid for: Aurora DB clusters and Multi-AZ DB clusters
	RestoreType *string

	// For DB clusters in serverless DB engine mode, the scaling properties of the DB
	// cluster. Valid for: Aurora DB clusters only
	ScalingConfiguration *types.ScalingConfiguration

	// Contains the scaling configuration of an Aurora Serverless v2 DB cluster. For
	// more information, see Using Amazon Aurora Serverless v2
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html)
	// in the Amazon Aurora User Guide.
	ServerlessV2ScalingConfiguration *types.ServerlessV2ScalingConfiguration

	// Specifies the storage type to be associated with the each DB instance in the
	// Multi-AZ DB cluster. Valid values: io1 When specified, a value for the Iops
	// parameter is required. Default: io1 Valid for: Multi-AZ DB clusters only
	StorageType *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []types.Tag

	// A value that indicates whether to restore the DB cluster to the latest
	// restorable backup time. By default, the DB cluster isn't restored to the latest
	// restorable backup time. Constraints: Can't be specified if RestoreToTime
	// parameter is provided. Valid for: Aurora DB clusters and Multi-AZ DB clusters
	UseLatestRestorableTime bool

	// A list of VPC security groups that the new DB cluster belongs to. Valid for:
	// Aurora DB clusters and Multi-AZ DB clusters
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBClusterToPointInTimeOutput struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster. For
	// an Amazon Aurora DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster, DeleteDBCluster, DescribeDBClusters,
	// FailoverDBCluster, ModifyDBCluster, PromoteReadReplicaDBCluster,
	// RestoreDBClusterFromS3, RestoreDBClusterFromSnapshot,
	// RestoreDBClusterToPointInTime, StartDBCluster, and StopDBCluster. For a Multi-AZ
	// DB cluster, this data type is used as a response element in the operations
	// CreateDBCluster, DeleteDBCluster, DescribeDBClusters, FailoverDBCluster,
	// ModifyDBCluster, RebootDBCluster, RestoreDBClusterFromSnapshot, and
	// RestoreDBClusterToPointInTime. For more information on Amazon Aurora DB
	// clusters, see  What is Amazon Aurora?
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
	// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
	// see  Multi-AZ deployments with two readable standby DB instances
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
	// in the Amazon RDS User Guide.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBClusterToPointInTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterToPointInTime{}, middleware.After)
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
	if err = addOpRestoreDBClusterToPointInTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterToPointInTime(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBClusterToPointInTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RestoreDBClusterToPointInTime",
	}
}
