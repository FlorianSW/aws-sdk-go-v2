// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Entity that comprises information on categorical values in data.
type CategoricalValues struct {

	// Indicates whether there is a potential data issue related to categorical values.
	//
	// This member is required.
	Status StatisticalIssueStatus

	// Indicates the number of categories in the data.
	NumberOfCategory *int32

	noSmithyDocumentSerde
}

// Entity that comprises information of count and percentage.
type CountPercent struct {

	// Indicates the count of occurences of the given statistic.
	//
	// This member is required.
	Count *int32

	// Indicates the percentage of occurances of the given statistic.
	//
	// This member is required.
	Percentage float32

	noSmithyDocumentSerde
}

// Provides information about a specified data ingestion job, including dataset
// information, data ingestion configuration, and status.
type DataIngestionJobSummary struct {

	// The Amazon Resource Name (ARN) of the dataset used in the data ingestion job.
	DatasetArn *string

	// The name of the dataset used for the data ingestion job.
	DatasetName *string

	// Specifies information for the input data for the data inference job, including
	// data Amazon S3 location parameters.
	IngestionInputConfiguration *IngestionInputConfiguration

	// Indicates the job ID of the data ingestion job.
	JobId *string

	// Indicates the status of the data ingestion job.
	Status IngestionJobStatus

	noSmithyDocumentSerde
}

// The configuration is the TargetSamplingRate , which is the sampling rate of the
// data after post processing by Amazon Lookout for Equipment. For example, if you
// provide data that has been collected at a 1 second level and you want the system
// to resample the data at a 1 minute rate before training, the TargetSamplingRate
// is 1 minute. When providing a value for the TargetSamplingRate , you must attach
// the prefix "PT" to the rate you want. The value for a 1 second rate is therefore
// PT1S, the value for a 15 minute rate is PT15M, and the value for a 1 hour rate
// is PT1H
type DataPreProcessingConfiguration struct {

	// The sampling rate of the data after post processing by Amazon Lookout for
	// Equipment. For example, if you provide data that has been collected at a 1
	// second level and you want the system to resample the data at a 1 minute rate
	// before training, the TargetSamplingRate is 1 minute. When providing a value for
	// the TargetSamplingRate , you must attach the prefix "PT" to the rate you want.
	// The value for a 1 second rate is therefore PT1S, the value for a 15 minute rate
	// is PT15M, and the value for a 1 hour rate is PT1H
	TargetSamplingRate TargetSamplingRate

	noSmithyDocumentSerde
}

// DataQualitySummary gives aggregated statistics over all the sensors about a
// completed ingestion job. It primarily gives more information about statistics
// over different incorrect data like MissingCompleteSensorData, MissingSensorData,
// UnsupportedDateFormats, InsufficientSensorData, DuplicateTimeStamps.
type DataQualitySummary struct {

	// Parameter that gives information about duplicate timestamps in the input data.
	//
	// This member is required.
	DuplicateTimestamps *DuplicateTimestamps

	// Parameter that gives information about insufficient data for sensors in the
	// dataset. This includes information about those sensors that have complete data
	// missing and those with a short date range.
	//
	// This member is required.
	InsufficientSensorData *InsufficientSensorData

	// Parameter that gives information about data that is invalid over all the
	// sensors in the input data.
	//
	// This member is required.
	InvalidSensorData *InvalidSensorData

	// Parameter that gives information about data that is missing over all the
	// sensors in the input data.
	//
	// This member is required.
	MissingSensorData *MissingSensorData

	// Parameter that gives information about unsupported timestamps in the input data.
	//
	// This member is required.
	UnsupportedTimestamps *UnsupportedTimestamps

	noSmithyDocumentSerde
}

// Provides information about the data schema used with the given dataset.
type DatasetSchema struct {

	// The data schema used within the given dataset.
	//
	// This value conforms to the media type: application/json
	InlineDataSchema *string

	noSmithyDocumentSerde
}

// Contains information about the specific data set, including name, ARN, and
// status.
type DatasetSummary struct {

	// The time at which the dataset was created in Amazon Lookout for Equipment.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the specified dataset.
	DatasetArn *string

	// The name of the dataset.
	DatasetName *string

	// Indicates the status of the dataset.
	Status DatasetStatus

	noSmithyDocumentSerde
}

// Entity that comprises information abount duplicate timestamps in the dataset.
type DuplicateTimestamps struct {

	// Indicates the total number of duplicate timestamps.
	//
	// This member is required.
	TotalNumberOfDuplicateTimestamps *int32

	noSmithyDocumentSerde
}

// Contains information about the specific inference event, including start and
// end time, diagnostics information, event duration and so on.
type InferenceEventSummary struct {

	// An array which specifies the names and values of all sensors contributing to an
	// inference event.
	Diagnostics *string

	// Indicates the size of an inference event in seconds.
	EventDurationInSeconds *int64

	// Indicates the ending time of an inference event.
	EventEndTime *time.Time

	// Indicates the starting time of an inference event.
	EventStartTime *time.Time

	// The Amazon Resource Name (ARN) of the inference scheduler being used for the
	// inference event.
	InferenceSchedulerArn *string

	// The name of the inference scheduler being used for the inference events.
	InferenceSchedulerName *string

	noSmithyDocumentSerde
}

// Contains information about the specific inference execution, including input
// and output data configuration, inference scheduling information, status, and so
// on.
type InferenceExecutionSummary struct {

	// The S3 object that the inference execution results were uploaded to.
	CustomerResultObject *S3Object

	// Indicates the time reference in the dataset at which the inference execution
	// stopped.
	DataEndTime *time.Time

	// Specifies configuration information for the input data for the inference
	// scheduler, including delimiter, format, and dataset location.
	DataInputConfiguration *InferenceInputConfiguration

	// Specifies configuration information for the output results from for the
	// inference execution, including the output Amazon S3 location.
	DataOutputConfiguration *InferenceOutputConfiguration

	// Indicates the time reference in the dataset at which the inference execution
	// began.
	DataStartTime *time.Time

	// Specifies the reason for failure when an inference execution has failed.
	FailedReason *string

	// The Amazon Resource Name (ARN) of the inference scheduler being used for the
	// inference execution.
	InferenceSchedulerArn *string

	// The name of the inference scheduler being used for the inference execution.
	InferenceSchedulerName *string

	// The Amazon Resource Name (ARN) of the machine learning model used for the
	// inference execution.
	ModelArn *string

	// The name of the machine learning model being used for the inference execution.
	ModelName *string

	// The model version used for the inference execution.
	ModelVersion *int64

	// The Amazon Resource Number (ARN) of the model version used for the inference
	// execution.
	ModelVersionArn *string

	// Indicates the start time at which the inference scheduler began the specific
	// inference execution.
	ScheduledStartTime *time.Time

	// Indicates the status of the inference execution.
	Status InferenceExecutionStatus

	noSmithyDocumentSerde
}

// Specifies configuration information for the input data for the inference,
// including Amazon S3 location of input data..
type InferenceInputConfiguration struct {

	// Specifies configuration information for the input data for the inference,
	// including timestamp format and delimiter.
	InferenceInputNameConfiguration *InferenceInputNameConfiguration

	// Indicates the difference between your time zone and Coordinated Universal Time
	// (UTC).
	InputTimeZoneOffset *string

	// Specifies configuration information for the input data for the inference,
	// including Amazon S3 location of input data.
	S3InputConfiguration *InferenceS3InputConfiguration

	noSmithyDocumentSerde
}

// Specifies configuration information for the input data for the inference,
// including timestamp format and delimiter.
type InferenceInputNameConfiguration struct {

	// Indicates the delimiter character used between items in the data.
	ComponentTimestampDelimiter *string

	// The format of the timestamp, whether Epoch time, or standard, with or without
	// hyphens (-).
	TimestampFormat *string

	noSmithyDocumentSerde
}

// Specifies configuration information for the output results from for the
// inference, including KMS key ID and output S3 location.
type InferenceOutputConfiguration struct {

	// Specifies configuration information for the output results from for the
	// inference, output S3 location.
	//
	// This member is required.
	S3OutputConfiguration *InferenceS3OutputConfiguration

	// The ID number for the KMS key key used to encrypt the inference output.
	KmsKeyId *string

	noSmithyDocumentSerde
}

// Specifies configuration information for the input data for the inference,
// including input data S3 location.
type InferenceS3InputConfiguration struct {

	// The bucket containing the input dataset for the inference.
	//
	// This member is required.
	Bucket *string

	// The prefix for the S3 bucket used for the input data for the inference.
	Prefix *string

	noSmithyDocumentSerde
}

// Specifies configuration information for the output results from the inference,
// including output S3 location.
type InferenceS3OutputConfiguration struct {

	// The bucket containing the output results from the inference
	//
	// This member is required.
	Bucket *string

	// The prefix for the S3 bucket used for the output results from the inference.
	Prefix *string

	noSmithyDocumentSerde
}

// Contains information about the specific inference scheduler, including data
// delay offset, model name and ARN, status, and so on.
type InferenceSchedulerSummary struct {

	// A period of time (in minutes) by which inference on the data is delayed after
	// the data starts. For instance, if an offset delay time of five minutes was
	// selected, inference will not begin on the data until the first data measurement
	// after the five minute mark. For example, if five minutes is selected, the
	// inference scheduler will wake up at the configured frequency with the additional
	// five minute delay time to check the customer S3 bucket. The customer can upload
	// data at the same frequency and they don't need to stop and restart the scheduler
	// when uploading new data.
	DataDelayOffsetInMinutes *int64

	// How often data is uploaded to the source S3 bucket for the input data. This
	// value is the length of time between data uploads. For instance, if you select 5
	// minutes, Amazon Lookout for Equipment will upload the real-time data to the
	// source bucket once every 5 minutes. This frequency also determines how often
	// Amazon Lookout for Equipment starts a scheduled inference on your data. In this
	// example, it starts once every 5 minutes.
	DataUploadFrequency DataUploadFrequency

	// The Amazon Resource Name (ARN) of the inference scheduler.
	InferenceSchedulerArn *string

	// The name of the inference scheduler.
	InferenceSchedulerName *string

	// Indicates whether the latest execution for the inference scheduler was
	// Anomalous (anomalous events found) or Normal (no anomalous events found).
	LatestInferenceResult LatestInferenceResult

	// The Amazon Resource Name (ARN) of the machine learning model used by the
	// inference scheduler.
	ModelArn *string

	// The name of the machine learning model used for the inference scheduler.
	ModelName *string

	// Indicates the status of the inference scheduler.
	Status InferenceSchedulerStatus

	noSmithyDocumentSerde
}

// Gives statistics about how many files have been ingested, and which files have
// not been ingested, for a particular ingestion job.
type IngestedFilesSummary struct {

	// Indicates the number of files that were successfully ingested.
	//
	// This member is required.
	IngestedNumberOfFiles *int32

	// Indicates the total number of files that were submitted for ingestion.
	//
	// This member is required.
	TotalNumberOfFiles *int32

	// Indicates the number of files that were discarded. A file could be discarded
	// because its format is invalid (for example, a jpg or pdf) or not readable.
	DiscardedFiles []S3Object

	noSmithyDocumentSerde
}

// Specifies configuration information for the input data for the data ingestion
// job, including input data S3 location.
type IngestionInputConfiguration struct {

	// The location information for the S3 bucket used for input data for the data
	// ingestion.
	//
	// This member is required.
	S3InputConfiguration *IngestionS3InputConfiguration

	noSmithyDocumentSerde
}

// Specifies S3 configuration information for the input data for the data
// ingestion job.
type IngestionS3InputConfiguration struct {

	// The name of the S3 bucket used for the input data for the data ingestion.
	//
	// This member is required.
	Bucket *string

	// The pattern for matching the Amazon S3 files that will be used for ingestion.
	// If the schema was created previously without any KeyPattern, then the default
	// KeyPattern {prefix}/{component_name}/* is used to download files from Amazon S3
	// according to the schema. This field is required when ingestion is being done for
	// the first time. Valid Values: {prefix}/{component_name}_* |
	// {prefix}/{component_name}/* | {prefix}/{component_name}[DELIMITER]* (Allowed
	// delimiters : space, dot, underscore, hyphen)
	KeyPattern *string

	// The prefix for the S3 location being used for the input data for the data
	// ingestion.
	Prefix *string

	noSmithyDocumentSerde
}

// Entity that comprises aggregated information on sensors having insufficient
// data.
type InsufficientSensorData struct {

	// Parameter that describes the total number of sensors that have data completely
	// missing for it.
	//
	// This member is required.
	MissingCompleteSensorData *MissingCompleteSensorData

	// Parameter that describes the total number of sensors that have a short date
	// range of less than 90 days of data overall.
	//
	// This member is required.
	SensorsWithShortDateRange *SensorsWithShortDateRange

	noSmithyDocumentSerde
}

// Entity that comprises aggregated information on sensors having insufficient
// data.
type InvalidSensorData struct {

	// Indicates the number of sensors that have at least some invalid values.
	//
	// This member is required.
	AffectedSensorCount *int32

	// Indicates the total number of invalid values across all the sensors.
	//
	// This member is required.
	TotalNumberOfInvalidValues *int32

	noSmithyDocumentSerde
}

// Contains information about the label group.
type LabelGroupSummary struct {

	// The time at which the label group was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the label group.
	LabelGroupArn *string

	// The name of the label group.
	LabelGroupName *string

	// The time at which the label group was updated.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// Contains the configuration information for the S3 location being used to hold
// label data.
type LabelsInputConfiguration struct {

	// The name of the label group to be used for label data.
	LabelGroupName *string

	// Contains location information for the S3 location being used for label data.
	S3InputConfiguration *LabelsS3InputConfiguration

	noSmithyDocumentSerde
}

// The location information (prefix and bucket name) for the s3 location being
// used for label data.
type LabelsS3InputConfiguration struct {

	// The name of the S3 bucket holding the label data.
	//
	// This member is required.
	Bucket *string

	// The prefix for the S3 bucket used for the label data.
	Prefix *string

	noSmithyDocumentSerde
}

// Information about the label.
type LabelSummary struct {

	// The time at which the label was created.
	CreatedAt *time.Time

	// The timestamp indicating the end of the label.
	EndTime *time.Time

	// Indicates that a label pertains to a particular piece of equipment.
	Equipment *string

	// Indicates the type of anomaly associated with the label. Data in this field
	// will be retained for service usage. Follow best practices for the security of
	// your data.
	FaultCode *string

	// The Amazon Resource Name (ARN) of the label group.
	LabelGroupArn *string

	// The name of the label group.
	LabelGroupName *string

	// The ID of the label.
	LabelId *string

	// Indicates whether a labeled event represents an anomaly.
	Rating LabelRating

	// The timestamp indicating the start of the label.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Entity that comprises information on large gaps between consecutive timestamps
// in data.
type LargeTimestampGaps struct {

	// Indicates whether there is a potential data issue related to large gaps in
	// timestamps.
	//
	// This member is required.
	Status StatisticalIssueStatus

	// Indicates the size of the largest timestamp gap, in days.
	MaxTimestampGapInDays *int32

	// Indicates the number of large timestamp gaps, if there are any.
	NumberOfLargeTimestampGaps *int32

	noSmithyDocumentSerde
}

// Entity that comprises information on sensors that have sensor data completely
// missing.
type MissingCompleteSensorData struct {

	// Indicates the number of sensors that have data missing completely.
	//
	// This member is required.
	AffectedSensorCount *int32

	noSmithyDocumentSerde
}

// Entity that comprises aggregated information on sensors having missing data.
type MissingSensorData struct {

	// Indicates the number of sensors that have atleast some data missing.
	//
	// This member is required.
	AffectedSensorCount *int32

	// Indicates the total number of missing values across all the sensors.
	//
	// This member is required.
	TotalNumberOfMissingValues *int32

	noSmithyDocumentSerde
}

// Provides information about the specified machine learning model, including
// dataset and model names and ARNs, as well as status.
type ModelSummary struct {

	// The model version that the inference scheduler uses to run an inference
	// execution.
	ActiveModelVersion *int64

	// The Amazon Resource Name (ARN) of the model version that is set as active. The
	// active model version is the model version that the inference scheduler uses to
	// run an inference execution.
	ActiveModelVersionArn *string

	// The time at which the specific model was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the dataset used to create the model.
	DatasetArn *string

	// The name of the dataset being used for the machine learning model.
	DatasetName *string

	// Indicates the most recent model version that was generated by retraining.
	LatestScheduledRetrainingModelVersion *int64

	// Indicates the start time of the most recent scheduled retraining run.
	LatestScheduledRetrainingStartTime *time.Time

	// Indicates the status of the most recent scheduled retraining run.
	LatestScheduledRetrainingStatus ModelVersionStatus

	// The Amazon Resource Name (ARN) of the machine learning model.
	ModelArn *string

	// The name of the machine learning model.
	ModelName *string

	// Indicates the date that the next scheduled retraining run will start on.
	// Lookout for Equipment truncates the time you provide to the nearest UTC day (https://docs.aws.amazon.com/https:/docs.aws.amazon.com/cli/latest/userguide/cli-usage-parameters-types.html#parameter-type-timestamp)
	// .
	NextScheduledRetrainingStartDate *time.Time

	// Indicates the status of the retraining scheduler.
	RetrainingSchedulerStatus RetrainingSchedulerStatus

	// Indicates the status of the machine learning model.
	Status ModelStatus

	noSmithyDocumentSerde
}

// Contains information about the specific model version.
type ModelVersionSummary struct {

	// The time when this model version was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the model that this model version is a
	// version of.
	ModelArn *string

	// The name of the model that this model version is a version of.
	ModelName *string

	// The version of the model.
	ModelVersion *int64

	// The Amazon Resource Name (ARN) of the model version.
	ModelVersionArn *string

	// Indicates how this model version was generated.
	SourceType ModelVersionSourceType

	// The current status of the model version.
	Status ModelVersionStatus

	noSmithyDocumentSerde
}

// Entity that comprises information on monotonic values in the data.
type MonotonicValues struct {

	// Indicates whether there is a potential data issue related to having monotonic
	// values.
	//
	// This member is required.
	Status StatisticalIssueStatus

	// Indicates the monotonicity of values. Can be INCREASING, DECREASING, or STATIC.
	Monotonicity Monotonicity

	noSmithyDocumentSerde
}

// Entity that comprises information on operating modes in data.
type MultipleOperatingModes struct {

	// Indicates whether there is a potential data issue related to having multiple
	// operating modes.
	//
	// This member is required.
	Status StatisticalIssueStatus

	noSmithyDocumentSerde
}

// Provides information about the specified retraining scheduler, including model
// name, status, start date, frequency, and lookback window.
type RetrainingSchedulerSummary struct {

	// The number of past days of data used for retraining.
	LookbackWindow *string

	// The ARN of the model that the retraining scheduler is attached to.
	ModelArn *string

	// The name of the model that the retraining scheduler is attached to.
	ModelName *string

	// The frequency at which the model retraining is set. This follows the ISO 8601 (https://en.wikipedia.org/wiki/ISO_8601#Durations)
	// guidelines.
	RetrainingFrequency *string

	// The start date for the retraining scheduler. Lookout for Equipment truncates
	// the time you provide to the nearest UTC day.
	RetrainingStartDate *time.Time

	// The status of the retraining scheduler.
	Status RetrainingSchedulerStatus

	noSmithyDocumentSerde
}

// Contains information about an S3 bucket.
type S3Object struct {

	// The name of the specific S3 bucket.
	//
	// This member is required.
	Bucket *string

	// The Amazon Web Services Key Management Service (KMS key) key being used to
	// encrypt the S3 object. Without this key, data in the bucket is not accessible.
	//
	// This member is required.
	Key *string

	noSmithyDocumentSerde
}

// Summary of ingestion statistics like whether data exists, number of missing
// values, number of invalid values and so on related to the particular sensor.
type SensorStatisticsSummary struct {

	// Parameter that describes potential risk about whether data associated with the
	// sensor is categorical.
	CategoricalValues *CategoricalValues

	// Name of the component to which the particular sensor belongs for which the
	// statistics belong to.
	ComponentName *string

	// Indicates the time reference to indicate the end of valid data associated with
	// the sensor that the statistics belong to.
	DataEndTime *time.Time

	// Parameter that indicates whether data exists for the sensor that the statistics
	// belong to.
	DataExists bool

	// Indicates the time reference to indicate the beginning of valid data associated
	// with the sensor that the statistics belong to.
	DataStartTime *time.Time

	// Parameter that describes the total number of duplicate timestamp records
	// associated with the sensor that the statistics belong to.
	DuplicateTimestamps *CountPercent

	// Parameter that describes the total number of invalid date entries associated
	// with the sensor that the statistics belong to.
	InvalidDateEntries *CountPercent

	// Parameter that describes the total number of, and percentage of, values that
	// are invalid for the sensor that the statistics belong to.
	InvalidValues *CountPercent

	// Parameter that describes potential risk about whether data associated with the
	// sensor contains one or more large gaps between consecutive timestamps.
	LargeTimestampGaps *LargeTimestampGaps

	// Parameter that describes the total number of, and percentage of, values that
	// are missing for the sensor that the statistics belong to.
	MissingValues *CountPercent

	// Parameter that describes potential risk about whether data associated with the
	// sensor is mostly monotonic.
	MonotonicValues *MonotonicValues

	// Parameter that describes potential risk about whether data associated with the
	// sensor has more than one operating mode.
	MultipleOperatingModes *MultipleOperatingModes

	// Name of the sensor that the statistics belong to.
	SensorName *string

	noSmithyDocumentSerde
}

// Entity that comprises information on sensors that have shorter date range.
type SensorsWithShortDateRange struct {

	// Indicates the number of sensors that have less than 90 days of data.
	//
	// This member is required.
	AffectedSensorCount *int32

	noSmithyDocumentSerde
}

// A tag is a key-value pair that can be added to a resource as metadata.
type Tag struct {

	// The key for the specified tag.
	//
	// This member is required.
	Key *string

	// The value for the specified tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Entity that comprises information abount unsupported timestamps in the dataset.
type UnsupportedTimestamps struct {

	// Indicates the total number of unsupported timestamps across the ingested data.
	//
	// This member is required.
	TotalNumberOfUnsupportedTimestamps *int32

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
