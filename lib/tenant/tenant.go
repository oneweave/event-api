package tenant

const (
	// Prefix is the event-id prefix for tenant events.
	Prefix = "tnt"

	// TenantRegisteredV1Type is the CloudEvent type for a newly registered tenant.
	TenantRegisteredV1Type = "tenant.registered.v1"

	// TenantUpdatedV1Type is the CloudEvent type for an updated tenant.
	TenantUpdatedV1Type = "tenant.updated.v1"
)

// TenantBaseData is the minimal payload shared by all tenant lifecycle events.
// It carries the tenant's unique ID and its human-readable name.
type TenantBaseData struct {
	TenantID string `json:"tenantId" bson:"tenant_id" validate:"required,eventid"`
	Status   string `bson:"status" json:"status" validate:"required,oneof=active retired"`
	Name     string `json:"name"     bson:"name"      validate:"required"`
}

// NewTenantBaseData returns a zero-value-safe TenantBaseData ready to be populated.
func NewTenantBaseData() TenantBaseData {
	return TenantBaseData{
		Status: "active",
	}
}
