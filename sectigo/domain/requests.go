package domain

import (
	"github.com/hm-edu/sectigo-client/sectigo/misc"
)

// ListItem represents a single item returned by the DomainService.List method.
type ListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CreateRequest represents the information required for a creating new domain.
type CreateRequest struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Active      bool                `json:"active"`
	Delegations []DelegationRequest `json:"delegations"`
}

// DelegationRequest represents the information about a delegation request.
type DelegationRequest struct {
	OrgID     int             `json:"orgId"`
	CertTypes []misc.CertType `json:"certTypes"`
}

// Delegation represents the information about a delegation status.
type Delegation struct {
	OrgID     int              `json:"orgId"`
	CertTypes []misc.CertType  `json:"certTypes"`
	Status    DelegationStatus `json:"status"`
}

// Details represents the information about a single domain.
type Details struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	DelegationStatus DelegationStatus `json:"delegationStatus"`
	State            State            `json:"state"`
	ValidationStatus ValidationStatus `json:"validationStatus"`
	DcvExpiration    misc.JSONDate    `json:"dcvExpiration"`
	Delegations      []Delegation     `json:"delegations"`
}
