package sectigo

import (
	"context"
	"fmt"

	"github.com/hm-edu/sectigo-client/sectigo/domain"
	"github.com/hm-edu/sectigo-client/sectigo/misc"
)

type DomainService struct {
	Client *Client
}

type ListDomainItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateDomainRequest struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Active      bool                `json:"active"`
	Delegations []DelegationRequest `json:"delegations"`
}

type DelegationRequest struct {
	OrgID     int             `json:"orgId"`
	CertTypes []misc.CertType `json:"certTypes"`
}

type Delegation struct {
	OrgID     int                     `json:"orgId"`
	CertTypes []misc.CertType         `json:"certTypes"`
	Status    domain.DelegationStatus `json:"status"`
}

type DomainInfos struct {
	Id               int                     `json:"id"`
	Name             string                  `json:"name"`
	DelegationStatus domain.DelegationStatus `json:"delegationStatus"`
	State            domain.State            `json:"state"`
	ValidationStatus domain.ValidationStatus `json:"validationStatus"`
	DcvExpiration    JsonDate                `json:"dcvExpiration"`
	Delegations      []Delegation            `json:"delegations"`
}

// List enumerates all existing domains
func (c *DomainService) List() (*[]ListDomainItem, error) {
	data, _, err := Get[[]ListDomainItem](context.Background(), c.Client, "/domain/v1")
	return data, err
}

// Infos gets the details of a single domain
func (c *DomainService) Infos(id int) (*DomainInfos, error) {
	data, _, err := Get[DomainInfos](context.Background(), c.Client, fmt.Sprintf("/domain/v1/%v", id))
	return data, err
}

// CreateDomain creates a new domain
func (c *DomainService) CreateDomain(request CreateDomainRequest) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, "/domain/v1", request)
	return err
}
