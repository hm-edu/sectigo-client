package sectigo

type OrganizationService struct {
	Client *Client
}

type Department struct {
	Id         int    `json:"id"`
	ParentName string `json:"parentName"`
	Name       string `json:"name"`
}

type OrganizationItem struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Departments []Department `json:"departments"`
}

type OrganizationDetails struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	CertTypes   []string `json:"certTypes"`
	Departments []struct {
		Id         int    `json:"id"`
		ParentName string `json:"parentName"`
		Name       string `json:"name"`
	} `json:"departments"`
	Address1                  string `json:"address1"`
	Address2                  string `json:"address2"`
	Address3                  string `json:"address3"`
	City                      string `json:"city"`
	StateOrProvince           string `json:"stateOrProvince"`
	PostalCode                string `json:"postalCode"`
	Country                   string `json:"country"`
	ValidationStatus          string `json:"validationStatus"`
	SecondaryValidationStatus string `json:"secondaryValidationStatus"`
	ClientCertificate         struct {
		AllowKeyRecoveryByMasterAdmins     bool `json:"allowKeyRecoveryByMasterAdmins"`
		AllowKeyRecoveryByOrgAdmins        bool `json:"allowKeyRecoveryByOrgAdmins"`
		AllowKeyRecoveryByDepartmentAdmins bool `json:"allowKeyRecoveryByDepartmentAdmins"`
	} `json:"clientCertificate"`
}

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
