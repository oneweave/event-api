package tenant

const (
	// Prefix is the event-id prefix for tenant events.
	Prefix = "tnt"

	// TenantRegisteredV1Type is the CloudEvent type for a newly registered tenant.
	TenantRegisteredV1Type = "tenant.registered.v1"

	// TenantUpdatedV1Type is the CloudEvent type for an updated tenant.
	TenantUpdatedV1Type = "tenant.updated.v1"
)

// TenantStatus represents the lifecycle state of a tenant.
type TenantStatus struct {
	Value string `bson:"value" json:"value" validate:"required,oneof=active retired"`
}

var (
	// TenantStatusActive indicates the tenant is live and operational.
	TenantStatusActive = TenantStatus{Value: "active"}
	// TenantStatusRetired indicates the tenant has been retired.
	TenantStatusRetired = TenantStatus{Value: "retired"}
)

// String returns the string representation of TenantStatus.
func (s TenantStatus) String() string {
	return s.Value
}

// TenantBaseData is the minimal payload shared by all tenant lifecycle events.
// It carries the tenant's unique ID and its human-readable name.
type TenantBaseData struct {
	TenantID string       `json:"tenantId" bson:"tenant_id" validate:"required,eventid"`
	Status   TenantStatus `json:"status"   bson:"status"    validate:"required"`
	Name     string       `json:"name"     bson:"name"      validate:"required"`
}

// NewTenantBaseData returns a zero-value-safe TenantBaseData ready to be populated.
func NewTenantBaseData() TenantBaseData {
	return TenantBaseData{
		Status: TenantStatusActive,
	}
}
