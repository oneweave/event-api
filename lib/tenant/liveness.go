package tenant

import (
	"github.com/oneweave/event-api/lib"
)

const (
	TenantManagementLivenessRequestedType = "tenant.management.liveness.requested"
	TenantManagementLivenessPassedType    = "tenant.management.liveness.passed"
	TenantManagementLivenessFailedType    = "tenant.management.liveness.failed"
)

// LivenessRequestedData is the payload for liveness request events.
type LivenessRequestedData struct{}

// NewLivenessRequestedData returns a new LivenessRequestedData.
func NewLivenessRequestedData() LivenessRequestedData {
	return LivenessRequestedData{}
}

// LivenessPassedData is the payload for successful liveness check events.
type LivenessPassedData struct {
	Status string `json:"status" bson:"status" validate:"required,eq=healthy"`
}

// NewLivenessPassedData returns a new LivenessPassedData.
func NewLivenessPassedData() LivenessPassedData {
	return LivenessPassedData{
		Status: "healthy",
	}
}

// LivenessFailedData is the payload for failed liveness check events.
type LivenessFailedData struct {
	Status string `json:"status" bson:"status" validate:"required,eq=unhealthy"`
	Error  string `json:"error" bson:"error" validate:"required"`
}

// NewLivenessFailedData returns a new LivenessFailedData.
func NewLivenessFailedData(errStr string) LivenessFailedData {
	return LivenessFailedData{
		Status: "unhealthy",
		Error:  errStr,
	}
}

type TenantManagementLivenessRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                `json:"type" bson:"type" validate:"required,eq=tenant.management.liveness.requested"`
	Data         LivenessRequestedData `json:"data" bson:"data"`
}

func NewTenantManagementLivenessRequestedCloudEvent() TenantManagementLivenessRequestedCloudEvent {
	return TenantManagementLivenessRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     TenantManagementLivenessRequestedType,
		Data:     NewLivenessRequestedData(),
	}
}

type TenantManagementLivenessPassedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=tenant.management.liveness.passed"`
	Data         LivenessPassedData `json:"data" bson:"data" validate:"required"`
}

func NewTenantManagementLivenessPassedCloudEvent() TenantManagementLivenessPassedCloudEvent {
	return TenantManagementLivenessPassedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     TenantManagementLivenessPassedType,
		Data:     NewLivenessPassedData(),
	}
}

type TenantManagementLivenessFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=tenant.management.liveness.failed"`
	Data         LivenessFailedData `json:"data" bson:"data" validate:"required"`
}

func NewTenantManagementLivenessFailedCloudEvent(errStr string) TenantManagementLivenessFailedCloudEvent {
	return TenantManagementLivenessFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     TenantManagementLivenessFailedType,
		Data:     NewLivenessFailedData(errStr),
	}
}
