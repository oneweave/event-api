package tenant

// TenantUser represents a single user belonging to a tenant.
// It follows the simple schema: ID, first name, second name, and email address.
type TenantUser struct {
	UserID     string `json:"userId"     bson:"user_id"     validate:"required,eventid"`
	FirstName  string `json:"firstName"  bson:"first_name"  validate:"required"`
	SecondName string `json:"secondName" bson:"second_name" validate:"required"`
	Email      string `json:"email"      bson:"email"       validate:"required,email"`
}

// NewTenantUser returns a zero-value-safe TenantUser ready to be populated.
func NewTenantUser() TenantUser {
	return TenantUser{}
}
