package tenant

import (
	"github.com/oneweave/event-api/lib"
)

// TenantUpdatedCloudEvent is the full CloudEvent envelope for tenant.updated.v1.
// The payload is the minimal TenantBaseData (tenantId + name) since no user
// information needs to be carried in the update event.
type TenantUpdatedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string         `json:"type" bson:"type" validate:"required,eq=tenant.updated.v1"`
	Data         TenantBaseData `json:"data" bson:"data" validate:"required"`
}

// NewTenantUpdatedCloudEvent returns a fully-initialised TenantUpdatedCloudEvent
// with envelope defaults set.
func NewTenantUpdatedCloudEvent() TenantUpdatedCloudEvent {
	return TenantUpdatedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     TenantUpdatedV1Type,
		Data:     NewTenantBaseData(),
	}
}
