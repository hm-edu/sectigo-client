package person

// CreateRequest represents the information required for creating a new person.
type CreateRequest struct {
	FirstName       string   `json:"firstName"`
	MiddleName      string   `json:"middleName"`
	LastName        string   `json:"lastName"`
	Email           string   `json:"email"`
	OrganizationID  int      `json:"organizationId"`
	ValidationType  string   `json:"validationType"`
	Phone           string   `json:"phone"`
	SecondaryEmails []string `json:"secondaryEmails"`
	CommonName      string   `json:"commonName"`
	Eppn            string   `json:"eppn"`
	Upn             string   `json:"upn"`
}

// ListParams provides the possible filters that can be passed to the PersonService.List method.
type ListParams struct {
	Position       int    `url:"position,omitempty"`
	Size           int    `url:"size,omitempty"`
	Name           string `url:"name,omitempty"`
	Email          string `url:"email,omitempty"`
	CommonName     string `url:"commonName,omitempty"`
	Phone          string `url:"phone,omitempty"`
	SecondaryEmail string `url:"secondaryEmail,omitempty"`
	OrganizationID int    `url:"organizationId,omitempty"`
}

// ListItem represents a single item returned by the PersonService.List method.
type ListItem struct {
	ID              int      `json:"id"`
	FirstName       string   `json:"firstName"`
	MiddleName      string   `json:"middleName"`
	LastName        string   `json:"lastName"`
	Email           string   `json:"email"`
	OrganizationID  int      `json:"organizationId"`
	ValidationType  string   `json:"validationType"`
	Phone           string   `json:"phone"`
	SecondaryEmails []string `json:"secondaryEmails"`
	CommonName      string   `json:"commonName"`
	Eppn            string   `json:"eppn"`
	Upn             string   `json:"upn"`
}
