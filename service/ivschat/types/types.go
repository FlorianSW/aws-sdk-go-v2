// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Configuration information for optional message review.
type MessageReviewHandler struct {

	// Specifies the fallback behavior (whether the message is allowed or denied) if
	// the handler does not return a valid response, encounters an error, or times out.
	// (For the timeout period, see  Service Quotas
	// (https://docs.aws.amazon.com/ivs/latest/userguide/service-quotas.html).) If
	// allowed, the message is delivered with returned content to all users connected
	// to the room. If denied, the message is not delivered to any user. Default:
	// ALLOW.
	FallbackResult FallbackResult

	// Identifier of the message review handler. Currently this must be an ARN of a
	// lambda function.
	Uri *string

	noSmithyDocumentSerde
}

// Summary information about a room.
type RoomSummary struct {

	// Room ARN.
	Arn *string

	// Time when the room was created. This is an ISO 8601 timestamp; note that this is
	// returned as a string.
	CreateTime *time.Time

	// Room ID, generated by the system. This is a relative identifier, the part of the
	// ARN that uniquely identifies the room.
	Id *string

	// Configuration information for optional review of messages.
	MessageReviewHandler *MessageReviewHandler

	// Room name. The value does not need to be unique.
	Name *string

	// Tags attached to the resource. See Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS Chat has no constraints beyond what is documented
	// there.
	Tags map[string]string

	// Time of the room’s last update. This is an ISO 8601 timestamp; note that this is
	// returned as a string.
	UpdateTime *time.Time

	noSmithyDocumentSerde
}

// This object is used in the ValidationException error.
type ValidationExceptionField struct {

	// Explanation of the reason for the validation error.
	//
	// This member is required.
	Message *string

	// Name of the field which failed validation.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
