package sectigo

import (
	"context"
	"fmt"

	"github.com/hm-edu/sectigo-client/sectigo/domain"
	"github.com/hm-edu/sectigo-client/sectigo/misc"
)

// DomainService provides some methods handling sectigo domain actions.
type DomainService struct {
	Client *Client
}

// ListDomainItem represents a single item returned by the DomainService.List method.
type ListDomainItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CreateDomainRequest represents the information required for a creating new domain.
type CreateDomainRequest struct {
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
	OrgID     int                     `json:"orgId"`
	CertTypes []misc.CertType         `json:"certTypes"`
	Status    domain.DelegationStatus `json:"status"`
}

// DomainInfos represents the information about a single domain.
type DomainInfos struct {
	ID               int                     `json:"id"`
	Name             string                  `json:"name"`
	DelegationStatus domain.DelegationStatus `json:"delegationStatus"`
	State            domain.State            `json:"state"`
	ValidationStatus domain.ValidationStatus `json:"validationStatus"`
	DcvExpiration    JSONDate                `json:"dcvExpiration"`
	Delegations      []Delegation            `json:"delegations"`
}

// List enumerates all existing domains.
func (c *DomainService) List() (*[]ListDomainItem, error) {
	data, _, err := Get[[]ListDomainItem](context.Background(), c.Client, "/domain/v1")
	return data, err
}

// Infos gets the details of a single domain.
func (c *DomainService) Infos(id int) (*DomainInfos, error) {
	data, _, err := Get[DomainInfos](context.Background(), c.Client, fmt.Sprintf("/domain/v1/%v", id))
	return data, err
}

// CreateDomain creates a new domain.
func (c *DomainService) CreateDomain(request CreateDomainRequest) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, "/domain/v1", request)
	return err
}
