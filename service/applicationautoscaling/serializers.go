// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationautoscaling

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsAwsjson11_serializeOpDeleteScalingPolicy struct {
}

func (*awsAwsjson11_serializeOpDeleteScalingPolicy) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDeleteScalingPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteScalingPolicyInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.DeleteScalingPolicy")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDeleteScalingPolicyInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDeleteScheduledAction struct {
}

func (*awsAwsjson11_serializeOpDeleteScheduledAction) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDeleteScheduledAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteScheduledActionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.DeleteScheduledAction")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDeleteScheduledActionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDeregisterScalableTarget struct {
}

func (*awsAwsjson11_serializeOpDeregisterScalableTarget) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDeregisterScalableTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeregisterScalableTargetInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.DeregisterScalableTarget")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDeregisterScalableTargetInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeScalableTargets struct {
}

func (*awsAwsjson11_serializeOpDescribeScalableTargets) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeScalableTargets) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeScalableTargetsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.DescribeScalableTargets")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeScalableTargetsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeScalingActivities struct {
}

func (*awsAwsjson11_serializeOpDescribeScalingActivities) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeScalingActivities) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeScalingActivitiesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.DescribeScalingActivities")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeScalingActivitiesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeScalingPolicies struct {
}

func (*awsAwsjson11_serializeOpDescribeScalingPolicies) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeScalingPolicies) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeScalingPoliciesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.DescribeScalingPolicies")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeScalingPoliciesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeScheduledActions struct {
}

func (*awsAwsjson11_serializeOpDescribeScheduledActions) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeScheduledActions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeScheduledActionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.DescribeScheduledActions")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDescribeScheduledActionsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpPutScalingPolicy struct {
}

func (*awsAwsjson11_serializeOpPutScalingPolicy) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpPutScalingPolicy) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutScalingPolicyInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.PutScalingPolicy")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentPutScalingPolicyInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpPutScheduledAction struct {
}

func (*awsAwsjson11_serializeOpPutScheduledAction) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpPutScheduledAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutScheduledActionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.PutScheduledAction")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentPutScheduledActionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpRegisterScalableTarget struct {
}

func (*awsAwsjson11_serializeOpRegisterScalableTarget) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpRegisterScalableTarget) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*RegisterScalableTargetInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AnyScaleFrontendService.RegisterScalableTarget")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentRegisterScalableTargetInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentCustomizedMetricSpecification(v *types.CustomizedMetricSpecification, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Dimensions != nil {
		ok := object.Key("Dimensions")
		if err := awsAwsjson11_serializeDocumentMetricDimensions(v.Dimensions, ok); err != nil {
			return err
		}
	}

	if v.MetricName != nil {
		ok := object.Key("MetricName")
		ok.String(*v.MetricName)
	}

	if v.Namespace != nil {
		ok := object.Key("Namespace")
		ok.String(*v.Namespace)
	}

	if len(v.Statistic) > 0 {
		ok := object.Key("Statistic")
		ok.String(string(v.Statistic))
	}

	if v.Unit != nil {
		ok := object.Key("Unit")
		ok.String(*v.Unit)
	}

	return nil
}

func awsAwsjson11_serializeDocumentMetricDimension(v *types.MetricDimension, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("Name")
		ok.String(*v.Name)
	}

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentMetricDimensions(v []types.MetricDimension, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentMetricDimension(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentPredefinedMetricSpecification(v *types.PredefinedMetricSpecification, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.PredefinedMetricType) > 0 {
		ok := object.Key("PredefinedMetricType")
		ok.String(string(v.PredefinedMetricType))
	}

	if v.ResourceLabel != nil {
		ok := object.Key("ResourceLabel")
		ok.String(*v.ResourceLabel)
	}

	return nil
}

func awsAwsjson11_serializeDocumentResourceIdsMaxLen1600(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentScalableTargetAction(v *types.ScalableTargetAction, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxCapacity != nil {
		ok := object.Key("MaxCapacity")
		ok.Integer(*v.MaxCapacity)
	}

	if v.MinCapacity != nil {
		ok := object.Key("MinCapacity")
		ok.Integer(*v.MinCapacity)
	}

	return nil
}

func awsAwsjson11_serializeDocumentStepAdjustment(v *types.StepAdjustment, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MetricIntervalLowerBound != nil {
		ok := object.Key("MetricIntervalLowerBound")
		ok.Double(*v.MetricIntervalLowerBound)
	}

	if v.MetricIntervalUpperBound != nil {
		ok := object.Key("MetricIntervalUpperBound")
		ok.Double(*v.MetricIntervalUpperBound)
	}

	if v.ScalingAdjustment != nil {
		ok := object.Key("ScalingAdjustment")
		ok.Integer(*v.ScalingAdjustment)
	}

	return nil
}

func awsAwsjson11_serializeDocumentStepAdjustments(v []types.StepAdjustment, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentStepAdjustment(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentStepScalingPolicyConfiguration(v *types.StepScalingPolicyConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.AdjustmentType) > 0 {
		ok := object.Key("AdjustmentType")
		ok.String(string(v.AdjustmentType))
	}

	if v.Cooldown != nil {
		ok := object.Key("Cooldown")
		ok.Integer(*v.Cooldown)
	}

	if len(v.MetricAggregationType) > 0 {
		ok := object.Key("MetricAggregationType")
		ok.String(string(v.MetricAggregationType))
	}

	if v.MinAdjustmentMagnitude != nil {
		ok := object.Key("MinAdjustmentMagnitude")
		ok.Integer(*v.MinAdjustmentMagnitude)
	}

	if v.StepAdjustments != nil {
		ok := object.Key("StepAdjustments")
		if err := awsAwsjson11_serializeDocumentStepAdjustments(v.StepAdjustments, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentSuspendedState(v *types.SuspendedState, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DynamicScalingInSuspended != nil {
		ok := object.Key("DynamicScalingInSuspended")
		ok.Boolean(*v.DynamicScalingInSuspended)
	}

	if v.DynamicScalingOutSuspended != nil {
		ok := object.Key("DynamicScalingOutSuspended")
		ok.Boolean(*v.DynamicScalingOutSuspended)
	}

	if v.ScheduledScalingSuspended != nil {
		ok := object.Key("ScheduledScalingSuspended")
		ok.Boolean(*v.ScheduledScalingSuspended)
	}

	return nil
}

func awsAwsjson11_serializeDocumentTargetTrackingScalingPolicyConfiguration(v *types.TargetTrackingScalingPolicyConfiguration, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CustomizedMetricSpecification != nil {
		ok := object.Key("CustomizedMetricSpecification")
		if err := awsAwsjson11_serializeDocumentCustomizedMetricSpecification(v.CustomizedMetricSpecification, ok); err != nil {
			return err
		}
	}

	if v.DisableScaleIn != nil {
		ok := object.Key("DisableScaleIn")
		ok.Boolean(*v.DisableScaleIn)
	}

	if v.PredefinedMetricSpecification != nil {
		ok := object.Key("PredefinedMetricSpecification")
		if err := awsAwsjson11_serializeDocumentPredefinedMetricSpecification(v.PredefinedMetricSpecification, ok); err != nil {
			return err
		}
	}

	if v.ScaleInCooldown != nil {
		ok := object.Key("ScaleInCooldown")
		ok.Integer(*v.ScaleInCooldown)
	}

	if v.ScaleOutCooldown != nil {
		ok := object.Key("ScaleOutCooldown")
		ok.Integer(*v.ScaleOutCooldown)
	}

	if v.TargetValue != nil {
		ok := object.Key("TargetValue")
		ok.Double(*v.TargetValue)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDeleteScalingPolicyInput(v *DeleteScalingPolicyInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.PolicyName != nil {
		ok := object.Key("PolicyName")
		ok.String(*v.PolicyName)
	}

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDeleteScheduledActionInput(v *DeleteScheduledActionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if v.ScheduledActionName != nil {
		ok := object.Key("ScheduledActionName")
		ok.String(*v.ScheduledActionName)
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDeregisterScalableTargetInput(v *DeregisterScalableTargetInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeScalableTargetsInput(v *DescribeScalableTargetsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.ResourceIds != nil {
		ok := object.Key("ResourceIds")
		if err := awsAwsjson11_serializeDocumentResourceIdsMaxLen1600(v.ResourceIds, ok); err != nil {
			return err
		}
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeScalingActivitiesInput(v *DescribeScalingActivitiesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeScalingPoliciesInput(v *DescribeScalingPoliciesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.PolicyNames != nil {
		ok := object.Key("PolicyNames")
		if err := awsAwsjson11_serializeDocumentResourceIdsMaxLen1600(v.PolicyNames, ok); err != nil {
			return err
		}
	}

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDescribeScheduledActionsInput(v *DescribeScheduledActionsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if v.ScheduledActionNames != nil {
		ok := object.Key("ScheduledActionNames")
		if err := awsAwsjson11_serializeDocumentResourceIdsMaxLen1600(v.ScheduledActionNames, ok); err != nil {
			return err
		}
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentPutScalingPolicyInput(v *PutScalingPolicyInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.PolicyName != nil {
		ok := object.Key("PolicyName")
		ok.String(*v.PolicyName)
	}

	if len(v.PolicyType) > 0 {
		ok := object.Key("PolicyType")
		ok.String(string(v.PolicyType))
	}

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	if v.StepScalingPolicyConfiguration != nil {
		ok := object.Key("StepScalingPolicyConfiguration")
		if err := awsAwsjson11_serializeDocumentStepScalingPolicyConfiguration(v.StepScalingPolicyConfiguration, ok); err != nil {
			return err
		}
	}

	if v.TargetTrackingScalingPolicyConfiguration != nil {
		ok := object.Key("TargetTrackingScalingPolicyConfiguration")
		if err := awsAwsjson11_serializeDocumentTargetTrackingScalingPolicyConfiguration(v.TargetTrackingScalingPolicyConfiguration, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentPutScheduledActionInput(v *PutScheduledActionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EndTime != nil {
		ok := object.Key("EndTime")
		ok.Double(smithytime.FormatEpochSeconds(*v.EndTime))
	}

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if v.ScalableTargetAction != nil {
		ok := object.Key("ScalableTargetAction")
		if err := awsAwsjson11_serializeDocumentScalableTargetAction(v.ScalableTargetAction, ok); err != nil {
			return err
		}
	}

	if v.Schedule != nil {
		ok := object.Key("Schedule")
		ok.String(*v.Schedule)
	}

	if v.ScheduledActionName != nil {
		ok := object.Key("ScheduledActionName")
		ok.String(*v.ScheduledActionName)
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	if v.StartTime != nil {
		ok := object.Key("StartTime")
		ok.Double(smithytime.FormatEpochSeconds(*v.StartTime))
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentRegisterScalableTargetInput(v *RegisterScalableTargetInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxCapacity != nil {
		ok := object.Key("MaxCapacity")
		ok.Integer(*v.MaxCapacity)
	}

	if v.MinCapacity != nil {
		ok := object.Key("MinCapacity")
		ok.Integer(*v.MinCapacity)
	}

	if v.ResourceId != nil {
		ok := object.Key("ResourceId")
		ok.String(*v.ResourceId)
	}

	if v.RoleARN != nil {
		ok := object.Key("RoleARN")
		ok.String(*v.RoleARN)
	}

	if len(v.ScalableDimension) > 0 {
		ok := object.Key("ScalableDimension")
		ok.String(string(v.ScalableDimension))
	}

	if len(v.ServiceNamespace) > 0 {
		ok := object.Key("ServiceNamespace")
		ok.String(string(v.ServiceNamespace))
	}

	if v.SuspendedState != nil {
		ok := object.Key("SuspendedState")
		if err := awsAwsjson11_serializeDocumentSuspendedState(v.SuspendedState, ok); err != nil {
			return err
		}
	}

	return nil
}
