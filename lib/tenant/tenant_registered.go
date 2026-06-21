package tenant

import (
	"github.com/oneweave/event-api/lib"
)

// TenantRegisteredData is the payload for the tenant.registered.v1 CloudEvent.
// It extends TenantBaseData with the full user lists captured at registration time.
type TenantRegisteredData struct {
	TenantBaseData `json:",inline" yaml:",inline"`
}

// NewTenantRegisteredData returns a zero-value-safe TenantRegisteredData.
func NewTenantRegisteredData() TenantRegisteredData {
	return TenantRegisteredData{
		TenantBaseData: NewTenantBaseData(),
	}
}

// TenantRegisteredCloudEvent is the full CloudEvent envelope for tenant.registered.v1.
type TenantRegisteredCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string               `json:"type" bson:"type" validate:"required,eq=tenant.registered.v1"`
	Data         TenantRegisteredData `json:"data" bson:"data" validate:"required"`
}

// NewTenantRegisteredCloudEvent returns a fully-initialised TenantRegisteredCloudEvent
// with envelope defaults set.
func NewTenantRegisteredCloudEvent() TenantRegisteredCloudEvent {
	return TenantRegisteredCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     TenantRegisteredV1Type,
		Data:     NewTenantRegisteredData(),
	}
}
