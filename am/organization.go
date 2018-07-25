package am

import "context"

const (
	// RNOrganizationSystem system only access (create/delete)
	RNOrganizationSystem = "lrn:service:organization:feature:system"
	// RNOrganizationManage organization specific management
	RNOrganizationManage = "lrn:service:organization:feature:manage"
)

// Organization represents an organization that has subscribed to our service
type Organization struct {
	OrgID           int    `json:"org_id"`
	OrgCID          string `json:"org_customer_id"`
	OrgName         string `json:"org_name"`
	OwnerEmail      string `json:"owner_email"`
	UserPoolID      string `json:"user_pool_id"`
	IdentityPoolID  string `json:"identity_pool_id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Phone           string `json:"phone"`
	Country         string `json:"country"`
	StatePrefecture string `json:"state_prefecture"`
	Street          string `json:"street"`
	Address1        string `json:"address1"`
	Address2        string `json:"address2"`
	City            string `json:"city"`
	PostalCode      string `json:"postal_code"`
	CreationTime    int64  `json:"creation_time"`
	StatusID        int    `json:"status_id"`
	Deleted         bool   `json:"deleted"`
	SubscriptionID  int    `json:"subscription_id"`
}

// OrgFilter for filtering organization list results
type OrgFilter struct {
	Start             int   `json:"start"`
	Limit             int   `json:"limit"`
	WithSubcription   bool  `json:"with_subscription"`
	SubValue          bool  `json:"sub_value"`
	SinceCreationTime int64 `json:"since_creation_time"`
	WithStatus        bool  `json:"with_status"`
	StatusValue       bool  `json:"status_value"`
	WithDeleted       bool  `json:"with_deleted"`
	DeletedValue      bool  `json:"deleted_value"`
}

// OrganizationService manages access to organizations
type OrganizationService interface {
	Get(ctx context.Context, userContext UserContext, orgName string) (oid int, org *Organization, err error)
	GetByCID(ctx context.Context, userContext UserContext, orgCID string) (oid int, org *Organization, err error)
	GetByID(ctx context.Context, userContext UserContext, orgID int) (oid int, org *Organization, err error)
	List(ctx context.Context, userContext UserContext, filter *OrgFilter) (orgs []*Organization, err error)
	Create(ctx context.Context, userContext UserContext, org *Organization) (oid int, uid int, ocid string, ucid string, err error)
	Update(ctx context.Context, userContext UserContext, org *Organization) (oid int, err error)
	Delete(ctx context.Context, userContext UserContext, orgID int) (oid int, err error)
}
