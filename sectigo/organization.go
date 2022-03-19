package sectigo

// OrganizationService provides some methods handling sectigo organization actions.
type OrganizationService struct {
	Client *Client
}

// Department represents some brief information about a department of the organization.
type Department struct {
	ID         int    `json:"id"`
	ParentName string `json:"parentName"`
	Name       string `json:"name"`
}

// OrganizationItem represents some brief information about the organization.
type OrganizationItem struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Departments []Department `json:"departments"`
}

// OrganizationDetails represents detailed information about a single organization.
type OrganizationDetails struct {
	ID                        int          `json:"id"`
	Name                      string       `json:"name"`
	CertTypes                 []string     `json:"certTypes"`
	Departments               []Department `json:"departments"`
	Address1                  string       `json:"address1"`
	Address2                  string       `json:"address2"`
	Address3                  string       `json:"address3"`
	City                      string       `json:"city"`
	StateOrProvince           string       `json:"stateOrProvince"`
	PostalCode                string       `json:"postalCode"`
	Country                   string       `json:"country"`
	ValidationStatus          string       `json:"validationStatus"`
	SecondaryValidationStatus string       `json:"secondaryValidationStatus"`
	ClientCertificate         struct {
		AllowKeyRecoveryByMasterAdmins     bool `json:"allowKeyRecoveryByMasterAdmins"`
		AllowKeyRecoveryByOrgAdmins        bool `json:"allowKeyRecoveryByOrgAdmins"`
		AllowKeyRecoveryByDepartmentAdmins bool `json:"allowKeyRecoveryByDepartmentAdmins"`
	} `json:"clientCertificate"`
}

// CreateOrganizationRequest represents the information required for creating a new organization or department.
type CreateOrganizationRequest struct {
	ParentOrgName     string `json:"parentOrgName"`
	Name              string `json:"name"`
	Address1          string `json:"address1"`
	Address2          string `json:"address2"`
	Address3          string `json:"address3"`
	City              string `json:"city"`
	StateProvince     string `json:"stateProvince"`
	PostalCode        string `json:"postalCode"`
	Country           string `json:"country"`
	ClientCertificate struct {
		AllowKeyRecoveryByMasterAdmins     bool `json:"allowKeyRecoveryByMasterAdmins"`
		AllowKeyRecoveryByOrgAdmins        bool `json:"allowKeyRecoveryByOrgAdmins"`
		AllowKeyRecoveryByDepartmentAdmins bool `json:"allowKeyRecoveryByDepartmentAdmins"`
	} `json:"clientCertificate"`
}
