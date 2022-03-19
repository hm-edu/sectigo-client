package sectigo

import (
	"context"
	"fmt"
	"github.com/hm-edu/sectigo-client/sectigo/domain"
)

// DomainService provides some methods handling sectigo domain actions.
type DomainService struct {
	Client *Client
}

// List enumerates all existing domains.
func (c *DomainService) List() (*[]domain.ListItem, error) {
	data, _, err := Get[[]domain.ListItem](context.Background(), c.Client, "/domain/v1")
	return data, err
}

// Infos gets the details of a single domain.
func (c *DomainService) Infos(id int) (*domain.Details, error) {
	data, _, err := Get[domain.Details](context.Background(), c.Client, fmt.Sprintf("/domain/v1/%v", id))
	return data, err
}

// CreateDomain creates a new domain.
func (c *DomainService) CreateDomain(request domain.CreateRequest) error {
	_, err := PostWithoutJSONResponse(context.Background(), c.Client, "/domain/v1", request)
	return err
}
