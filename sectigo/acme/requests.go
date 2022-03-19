package acme

// CreateACMERequest represnts the information that are required for creating a new ACME account.
type CreateACMERequest struct {
	Name           string `json:"name"`
	AcmeServer     string `json:"acmeServer"`
	OrganizationID int    `json:"organizationId"`
}

// Domain is an internal wrapper around a single domain.
type Domain struct {
	Name string `json:"name"`
}

// AddOrRemoveDomainsRequest contains the domains shal be added using ACMEService.AddDomains or ACMEService.DeleteDomains.
type AddOrRemoveDomainsRequest struct {
	Domains []Domain `json:"domains"`
}

// AddDomainsResponse contains the domains that where not added using ACMEService.AddDomains.
type AddDomainsResponse struct {
	NotAddedDomains []string `json:"notAddedDomains"`
}

// RemoveDomainsResponse contains the domains that where not removed using ACMEService.DeleteDomains.
type RemoveDomainsResponse struct {
	NotRemovedDomains []string `json:"notRemovedDomains"`
}

// ListACMEItem represents a single item returned by ACMEService.List.
type ListACMEItem struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	Status             string   `json:"status"`
	MacKey             string   `json:"macKey"`
	MacID              string   `json:"macId"`
	AcmeServer         string   `json:"acmeServer"`
	OrganizationID     int      `json:"organizationId"`
	CertValidationType string   `json:"certValidationType"`
	AccountID          string   `json:"accountId"`
	OvOrderNumber      int      `json:"ovOrderNumber"`
	Contacts           string   `json:"contacts"`
	Domains            []Domain `json:"domains"`
}

// ListACMEServerItem represents a single ACME server item returned by ACMEService.ListServers.
type ListACMEServerItem struct {
	Active             bool   `json:"active"`
	URL                string `json:"url"`
	CaID               int    `json:"caId"`
	Name               string `json:"name"`
	SingleProductID    int    `json:"singleProductId"`
	MultiProductID     int    `json:"multiProductId"`
	WcProductID        int    `json:"wcProductId"`
	CertValidationType string `json:"certValidationType"`
}

// ListRequest represents different parameters for the ACEMService.List method.
type ListRequest struct {
	// OrganizationID is required for the query.
	OrganizationID int    `url:"organizationId,omitempty"`
	Name           string `url:"name,omitempty"`
}
